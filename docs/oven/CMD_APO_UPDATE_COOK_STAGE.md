# CMD_APO_UPDATE_COOK_STAGE

Modify an existing cook stage

Syntax:

```
{
  "command": "CMD_APO_UPDATE_COOK_STAGE",
  "payload": {
    "payload": {
      "stepType": "stage",
      "id": "android-<uuid>",					# The ID of the existing stage you are modifying
      "title": "",
      "description": "",
      "type": "preheat",
      "userActionRequired": true,
      "temperatureBulbs": {
        "wet": {
          "setpoint": {
            "fahrenheit": 130,
            "celsius": 55
          }
        },
        "mode": "wet"
      },
      "heatingElements": {
        "rear": {
          "on": true
        },
        "top": {
          "on": false
        },
        "bottom": {
          "on": false
        }
      },
      "fan": {
        "speed": 100
      },
      "vent": {
        "open": false
      },
      "rackPosition": 3,
      "steamGenerators": {
        "relativeHumidity": {
          "setpoint": 100
        },
        "mode": "relative-humidity"
      }
    },
    "type": "CMD_APO_UPDATE_COOK_STAGE",
    "id": "<your_oven_id>"
  },
  "requestId": "<uuid>"
}
```

[Example](../examples/CMD_APO_UPDATE_COOK_STAGE.json)