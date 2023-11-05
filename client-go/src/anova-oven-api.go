package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v3"
)

const (
	AUTH_URL              = "https://securetoken.googleapis.com/v1/token"
	WS_URL                = "devices.anovaculinary.io"
	CREDENTIALS_FILE_NAME = "creds_anova.yml"
)

func decodeStateMessage(jsonResult map[string]interface{}) {
	ovenId := jsonResult["payload"].(map[string]interface{})["cookerId"].(string)
	fullState := jsonResult["payload"].(map[string]interface{})["state"].(map[string]interface{})
	currentMode := fullState["state"].(map[string]interface{})["mode"].(string)
	currentTemperature := fullState["nodes"].(map[string]interface{})["temperatureBulbs"].(map[string]interface{})["dry"].(map[string]interface{})["current"].(map[string]interface{})["celsius"].(float64)
	timeLeftOnTimer := int(fullState["nodes"].(map[string]interface{})["timer"].(map[string]interface{})["current"].(float64))
	fmt.Printf("Oven with ID |%s|: mode %s, current temperature (dry) %.2f, timer %02d:%02d\n", ovenId, currentMode, currentTemperature, timeLeftOnTimer/60, timeLeftOnTimer%60)
}

func connectWebsocket(accessToken string) (c *websocket.Conn, err error) {
	fmt.Println("Preparing WS connection")
	var values url.Values = url.Values{}
	u := url.URL{Scheme: "wss", Host: WS_URL, Path: "/"}
	values.Add("token", accessToken)
	values.Add("supportedAccessories", "APO")
	values.Add("platform", "android")
	u.RawQuery = values.Encode()

	fmt.Println("Connecting to websocket")
	c, _, err = websocket.DefaultDialer.Dial(u.String(),
		http.Header{
			"Sec-WebSocket-Protocol": {"ANOVA_V2"},
		})

	if err != nil {
		return nil, err
	}

	return c, nil
}

func receiveMessages(c *websocket.Conn, showFullJSON bool) {
	jsonResult := make(map[string]interface{})
	var prettyJson bytes.Buffer
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}
		if showFullJSON {
			prettyJson.Reset()
			json.Indent(&prettyJson, message, "", "  ")
			fmt.Printf("recv: %s\n", prettyJson.String())
		}
		json.Unmarshal(message, &jsonResult)
		if jsonResult["command"].(string) == "EVENT_APO_STATE" {
			decodeStateMessage(jsonResult)
		}
	}
}

func readCredentials() (credentials map[string]string, err error) {
	yfile, err := os.ReadFile(CREDENTIALS_FILE_NAME)
	if err != nil {
		return nil, err
	}

	credentials = make(map[string]string)
	err = yaml.Unmarshal(yfile, credentials)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func getAccessToken(refreshToken string, apiKey string) (string, error) {
	var (
		authURL     *url.URL
		err         error
		ok          bool
		accessToken string
	)

	authURL, err = url.Parse(AUTH_URL)
	if err != nil {
		return "", err
	}

	values := url.Values{}
	values.Add("key", apiKey)
	authURL.RawQuery = values.Encode()

	response, err := http.PostForm(authURL.String(), url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	})
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var jsonResult map[string]string = make(map[string]string)
	if err := json.Unmarshal(body, &jsonResult); err != nil {
		return "", err
	}
	accessToken, ok = jsonResult["access_token"]
	if !ok {
		return "", errors.New("unable to locate access token in response")
	}

	return accessToken, nil
}

// Send a message to start/stop a cook.
// The message is based on a Go template stored in the file <templateName>
func sendOvenMessage(c *websocket.Conn, templateName string, cookerID string) error {
	var fillData struct {
		CookerID string
		UUIDs    []string
	}
	fillData.CookerID = cookerID
	fillData.UUIDs = make([]string, 10)
	for i := 0; i < 10; i++ {
		fillData.UUIDs[i] = uuid.New().String()
	}

	fmt.Println("Preparing cookStart message")
	tpl, err := template.ParseFiles(templateName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	buf := new(bytes.Buffer)
	fmt.Println("Executing template")
	err = tpl.Execute(buf, fillData)
	if err != nil {
		return err
	}
	// os.Stdout.Write(buf.Bytes())

	fmt.Println("Sending message to WS")
	err = c.WriteMessage(websocket.TextMessage, buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var (
		credentials  map[string]string
		refreshToken string
		accessToken  string
		apiKey       string
		err          error
		c            *websocket.Conn
	)

	var (
		showFullJSON *bool // Command-line flags
		startCook    *string
		stopCook     *bool
	)

	// Process flags
	showFullJSON = flag.Bool("showjson", false, "show the full JSON in received messages")
	startCook = flag.String("cook", "none", "name of a template file that will be used to start cooking")
	stopCook = flag.Bool("stop", false, "stop the current cooking process. <cook> will be ignored if <stop> is set!")
	flag.Parse()

	fmt.Println("Reading credentials from file")
	credentials, err = readCredentials()
	if err != nil {
		panic(err)
	}
	refreshToken = credentials["refresh_token"]
	apiKey = credentials["api_key"]

	fmt.Println("Creating an access token")
	accessToken, err = getAccessToken(refreshToken, apiKey)
	if err != nil {
		panic(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, err = connectWebsocket(accessToken)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	done := make(chan struct{})
	defer close(done)

	go receiveMessages(c, *showFullJSON)

	if *stopCook {
		err = sendOvenMessage(c, "templates/cook_stop.json.tpl", credentials["oven_id"])
		if err != nil {
			panic(err)
		}
	}

	if (!*stopCook) && (*startCook != "none") {
		err = sendOvenMessage(c, *startCook, credentials["oven_id"])
		if err != nil {
			panic(err)
		}
	}

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			fmt.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("error on write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
