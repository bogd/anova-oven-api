# Anova API Documentation

## Disclaimer

> [!WARNING]
> 
> This project is in no way affiliated with Anova. Anything you find here is just based on whatever public information I was able to gather. And if the past is any guide, things can (and will) change based on Anova's whims. 
> 
> Use this at your own risk.

## Getting the credentials

### Getting the refresh token

The actual application runs on Firebase, and authenticates via Google's Identity Platform (OAuth, I believe?). So you will need either an access token (which has a limited lifetime before it expires), or the refresh token (which you can exchange for a new access token whenever needed).

1. Using Chrome, connect to [https://oven.anovaculinary.com](https://oven.anovaculinary.com) . Authenticate using whatever mechanism you normally use. 
2. Press F12, and go to `Application > Storage > IndexedDB > firebaseLocalStorageDB > firebaseLocalStorage > value > stsTokenManager > refreshToken`
3. If you are able to copy/paste that token, great! In my case, Chrome did not allow that, so I had to install a third-party extension (IndexedDBEdit). This adds a new tab under F12, so you can easily go to `F12 > IndexedDBEdit > firebaseLocalStorageDB > firebaseLocalStorage` and get the `refreshToken` value from the JSON.
4. While you are here, also get the `apiKey` (it should be the same one for everyone)

### Getting the access token
You could get the access token directly from the IndexedDB, but (as mentioned above) it will expire in a short time. If you have the refresh token, you can exchange it for a new access token whenever you wish, as described in [Google's docs](https://cloud.google.com/identity-platform/docs/use-rest-api). I am using `curl` here as an example:

```
curl 'https://securetoken.googleapis.com/v1/token?key=<API_KEY>' \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    --data 'grant_type=refresh_token&refresh_token=<REFRESH_TOKEN>'
```

Replace `API_KEY` and `REFRESH_TOKEN` with the values you got in the previous step.

You will get a JSON as a result. Copy the value for `id_token` - this is the one you will need later.

### Connecting to the API

The API runs on `wss://devices.anovaculinary.io` . You will need to connect to it using a client that supports WebSockets - either postman or an online tool like [this one](https://www.piesocket.com/websocket-tester).

Before connecting, you will need to set the following:
1. Query parameters:
* token: the access token (`id_token`) you got in the previous step
* supportedAccessories:
- APO (for Anova Precision Oven)
- APC (for Anova Precision Cooker)
- a combination if you have both (for example, `APO,APC`)
* platform: ios (you can also use `android` as a value, if you prefer :) )
1. Headers:
* (for APO only) add a header called `Sec-WebSocket-Protocol`, with a value of `ANOVA_V2`

Example - setting query parameters in postman:

![Query Parameters](assets/parameters.png)

Example - setting headers in postman:

![Headers](assets/headers.png)

### Getting information

Once connected, you should start to receive `EVENT_APO_STATE` and/or `EVENT_APC_STATE` messages (depending on the devices you have registered, and the headers you set when connecting to the API). They are received periodically, about every 30 seconds.

You will probably also receive an `EVENT_APO_WIFI_LIST`/`EVENT_APC_WIFI_LIST` message.

![State messages](assets/connection.png)

The documentation for the messages is here:
* [EVENT_APO_STATE](./oven/EVENT_APO_STATE.md)
* [EVENT_APC_STATE](./oven/EVENT_APC_STATE.md)


Look for `cookerId` in the received messages - that is your device (oven/sous-vide) ID, and you will need that when sending commands to it!

### Sending commands

*Note*: There are many places where the requests use UUIDs. As far as I was able to tell, you could place pretty much anything there, and it will work. The recommended way is to use version 4 UUIDs - you can generate them for example [here](https://www.uuidgenerator.net/).

To tell a device what to do, you need to send a JSON-formatted message over the websocket connection. Make sure to place your own device ID (collected from the `STATE` messages - see above) under `id` (for oven) or `cookerId` (for sous-vide cooker)!

As mentioned above, whenever the command needs a UUID, you can write your own, or you can [generate](https://www.uuidgenerator.net/) one.

Oven commands:
* Start cook: [CMD_APO_START](./oven/CMD_APO_START.md)
* Stop cook: [CMD_APO_STOP](./oven/CMD_APO_STOP.md)
* Update cook stages (e.g. add a stage): [CMD_APO_UPDATE_COOK_STAGES](./oven/CMD_APO_UPDATE_COOK_STAGES.md)
* Update a specific cook stage: [CMD_APO_UPDATE_COOK_STAGE](./oven/CMD_APO_UPDATE_COOK_STAGE.md)
* Start another stage: [CMD_APO_START_STAGE](./oven/CMD_APO_START_STAGE.md)

Cooker commands:
* Start cook: [CMD_APC_START](./sous_vide/CMD_APC_START.md)
* Stop cook: [CMD_APC_STOP](./sous_vide/CMD_APC_STOP.md)
* Set target temperature: [CMD_APC_SET_TARGET_TEMP](./sous_vide/CMD_APC_SET_TARGET_TEMP.md)
* Set timer: [CMD_APC_SET_TIMER](./sous_vide/CMD_APC_SET_TIMER.md)

For every command with a `requestId`, you should receive a `RESPONSE` message of this form:
```
{
  "command": "RESPONSE",
  "requestId": <id>,
  "payload": {
    "status": "ok"
  }
} 
```

Also, the `EVENT_*_STATE` messages will continue to arrive periodically, as described above.