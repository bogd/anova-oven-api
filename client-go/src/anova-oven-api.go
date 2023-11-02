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
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v3"
)

const (
	AUTH_URL              = "https://securetoken.googleapis.com/v1/token"
	WS_URL                = "devices.anovaculinary.io"
	CREDENTIALS_FILE_NAME = "creds_anova.yml"
)

func decodeStateMessage(jsonResult map[string]interface{}) {
	fmt.Println("================ Oven mode is ===============")
	fmt.Println(jsonResult["payload"].(map[string]interface{})["state"].(map[string]interface{})["state"].(map[string]interface{})["mode"].(string))
	fmt.Println("=============================================")
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

// Send a message to start a cook.
// The message is based on a Go template stored in the file <templateName>
func sendStartMessage(templateName string) {
	fmt.Println("Sending cookStart message")
}

func main() {
	var (
		credentials  map[string]string
		refreshToken string
		accessToken  string
		apiKey       string
		err          error
	)

	var (
		showFullJSON *bool // Command-line flags
		startCook    *string
	)

	showFullJSON = flag.Bool("showjson", false, "show the full JSON in received messages")
	startCook = flag.String("cook", "none", "name of a template file that will be used to start cooking")
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

	fmt.Println("Preparing WS connection")
	var values url.Values = url.Values{}
	u := url.URL{Scheme: "wss", Host: WS_URL, Path: "/"}
	values.Add("token", accessToken)
	values.Add("supportedAccessories", "APO")
	values.Add("platform", "android")
	u.RawQuery = values.Encode()

	fmt.Println("Connecting to websocket")

	c, _, err := websocket.DefaultDialer.Dial(u.String(),
		http.Header{
			"Sec-WebSocket-Protocol": {"ANOVA_V2"},
		})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	done := make(chan struct{})
	defer close(done)

	go receiveMessages(c, *showFullJSON)

	if *startCook != "none" {
		sendStartMessage(*startCook)
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
				fmt.Println("write close:", err)
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
