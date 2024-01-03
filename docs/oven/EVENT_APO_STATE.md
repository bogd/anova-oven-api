# EVENT_APO_STATE

Received periodically (every 30 seconds when device is idle, every 2 seconds when a cook is in progress). 

Syntax:
```
{
   "command" : "EVENT_APO_STATE",
   "payload" : {
      "cookerId" : "01234XXXXXXXXXX",
      "state" : {
         "nodes" : {
            "door" : {
               "closed" : true
            },
            "fan" : {
               "failed" : false,
               "speed" : 0
            },
            "heatingElements" : {
               "bottom" : {
                  "failed" : false,
                  "on" : true,
                  "watts" : 0
               },
               "rear" : {
                  "failed" : false,
                  "on" : false,
                  "watts" : 0
               },
               "top" : {
                  "failed" : false,
                  "on" : false,
                  "watts" : 0
               }
            },
            "lamp" : {
               "failed" : false,
               "on" : false,
               "preference" : "on"
            },
            "steamGenerators" : {
               "boiler" : {
                  "celsius" : 38.75,
                  "descaleRequired" : false,
                  "dosed" : false,
                  "failed" : false,
                  "overheated" : false,
                  "watts" : 0
               },
               "evaporator" : {
                  "celsius" : 38.75,
                  "failed" : false,
                  "overheated" : false,
                  "watts" : 0
               },
               "mode" : "idle",
               "relativeHumidity" : {
                  "current" : 100
               }
            },
            "temperatureBulbs" : {
               "dry" : {
                  "current" : {
                     "celsius" : 24.36,
                     "fahrenheit" : 75.85
                  },
                  "setpoint" : {
                     "celsius" : 58.33,
                     "fahrenheit" : 137
                  }
               },
               "dryBottom" : {
                  "current" : {
                     "celsius" : 23.73,
                     "fahrenheit" : 74.72
                  },
                  "overheated" : false
               },
               "dryTop" : {
                  "current" : {
                     "celsius" : 24.36,
                     "fahrenheit" : 75.85
                  },
                  "overheated" : false
               },
               "mode" : "dry",
               "wet" : {
                  "current" : {
                     "celsius" : 24.36,
                     "fahrenheit" : 75.85
                  },
                  "doseFailed" : false,
                  "dosed" : false
               }
            },
            "temperatureProbe" : {
               "connected" : false
            },
            "timer" : {
               "current" : 0,
               "initial" : 0,
               "mode" : "idle"
            },
            "userInterfaceCircuit" : {
               "communicationFailed" : false
            },
            "vent" : {
               "open" : true
            },
            "waterTank" : {
               "empty" : false
            }
         },
         "state" : {
            "mode" : "idle",
            "processedCommandIds" : [
               "59320e4b-114b-480e-8a59-3c4139111495",
               "c61bdf52-f79d-40ff-8f63-062effa784a4",
               "9450de29-55ee-429d-896d-3be22a2eed1f",
               "dbd56098-2e44-4b42-8533-73264bc8ee16",
               "d7e6155d-b308-44ea-906e-abd09db2d91e",
               "4eda829c-dc56-45e0-83b3-f462338ea952",
               "d0364ff1-e975-4548-8d97-638ea8f28e26",
               "5ee8134c-930e-4228-a9ac-5c96eb93c006",
               "8becfbe5-0429-4521-8741-aba2163edca7",
               "7a981aab-3de8-49a2-91d4-3036ff2d3a47"
            ],
            "temperatureUnit" : "F"
         },
         "systemInfo" : {
            "firmwareUpdatedTimestamp" : "2023-10-18T09:53:56Z",
            "firmwareVersion" : "2.1.7",
            "hardwareVersion" : "120V Universal",
            "lastConnectedTimestamp" : "2024-01-03T21:51:31Z",
            "lastDisconnectedTimestamp" : "2024-01-03T21:51:28Z",
            "online" : true,
            "powerHertz" : 60,
            "powerMains" : 120,
            "triacsFailed" : false,
            "uiFirmwareVersion" : "1.0.22",
            "uiHardwareVersion" : "UI_RENASAS"
         },
         "updatedTimestamp" : "2024-01-03T22:09:05Z",
         "version" : 1
      },
      "type" : "oven_v1"
   }
}
```
