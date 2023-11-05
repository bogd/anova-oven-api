{
  "command": "CMD_APO_START",
  "payload": {
    "payload": {
      "cookId": "android-{{index .UUIDs 0}}",
      "stages": [
        {
          "stepType": "stage",
          "id": "android-{{index .UUIDs 1}}",
          "title": "Stage 1",
          "description": "",
          "type": "preheat",
          "userActionRequired": false,
          "temperatureBulbs": {
            "dry": {
              "setpoint": {
                "fahrenheit": 356,
                "celsius": 180
              }
            },
            "mode": "dry"
          },
          "heatingElements": {
            "bottom": {
              "on": false
            },
            "top": {
              "on": false
            },
            "rear": {
              "on": true
            }
          },
          "fan": {
            "speed": 100
          },
          "vent": {
            "open": false
          },
          "rackPosition": 1
        },
        {
          "stepType": "stage",
          "id": "android-{{index .UUIDs 2}}",
          "title": "",
          "description": "",
          "type": "cook",
          "userActionRequired": false,
          "temperatureBulbs": {
            "dry": {
              "setpoint": {
                "fahrenheit": 356,
                "celsius": 180
              }
            },
            "mode": "dry"
          },
          "heatingElements": {
            "bottom": {
              "on": false
            },
            "top": {
              "on": false
            },
            "rear": {
              "on": true
            }
          },
          "fan": {
            "speed": 100
          },
          "vent": {
            "open": false
          },
          "rackPosition": 1,
          "timerAdded": true,
          "probeAdded": false,
          "timer": {
            "initial": 1200
          }
        },
        {
          "stepType": "stage",
          "id": "android-{{index .UUIDs 3}}",
          "title": "Stage 2",
          "description": "",
          "type": "preheat",
          "userActionRequired": true,
          "temperatureBulbs": {
            "dry": {
              "setpoint": {
                "fahrenheit": 392,
                "celsius": 200
              }
            },
            "mode": "dry"
          },
          "heatingElements": {
            "bottom": {
              "on": false
            },
            "top": {
              "on": false
            },
            "rear": {
              "on": true
            }
          },
          "fan": {
            "speed": 100
          },
          "vent": {
            "open": false
          },
          "rackPosition": 1,
          "steamGenerators": {
            "steamPercentage": {
              "setpoint": 100
            },
            "mode": "steam-percentage"
          }
        },
        {
          "stepType": "stage",
          "id": "android-{{index .UUIDs 4}}",
          "title": "",
          "description": "",
          "type": "cook",
          "userActionRequired": false,
          "temperatureBulbs": {
            "dry": {
              "setpoint": {
                "fahrenheit": 392,
                "celsius": 200
              }
            },
            "mode": "dry"
          },
          "heatingElements": {
            "bottom": {
              "on": false
            },
            "top": {
              "on": false
            },
            "rear": {
              "on": true
            }
          },
          "fan": {
            "speed": 100
          },
          "vent": {
            "open": false
          },
          "rackPosition": 1,
          "timerAdded": true,
          "probeAdded": false,
          "steamGenerators": {
            "steamPercentage": {
              "setpoint": 100
            },
            "mode": "steam-percentage"
          },
          "timer": {
            "initial": 1200
          }
        },
        {
          "stepType": "stage",
          "id": "android-{{index .UUIDs 5}}",
          "title": "Stage 3",
          "description": "",
          "type": "preheat",
          "userActionRequired": false,
          "temperatureBulbs": {
            "dry": {
              "setpoint": {
                "fahrenheit": 410,
                "celsius": 210
              }
            },
            "mode": "dry"
          },
          "heatingElements": {
            "bottom": {
              "on": false
            },
            "top": {
              "on": false
            },
            "rear": {
              "on": true
            }
          },
          "fan": {
            "speed": 100
          },
          "vent": {
            "open": false
          },
          "rackPosition": 3
        },
        {
          "stepType": "stage",
          "id": "android-{{index .UUIDs 6}}",
          "title": "",
          "description": "",
          "type": "cook",
          "userActionRequired": false,
          "temperatureBulbs": {
            "dry": {
              "setpoint": {
                "fahrenheit": 410,
                "celsius": 210
              }
            },
            "mode": "dry"
          },
          "heatingElements": {
            "bottom": {
              "on": false
            },
            "top": {
              "on": false
            },
            "rear": {
              "on": true
            }
          },
          "fan": {
            "speed": 100
          },
          "vent": {
            "open": false
          },
          "rackPosition": 3,
          "timerAdded": true,
          "probeAdded": false,
          "timer": {
            "initial": 120
          }
        }
      ]
    },
    "type": "CMD_APO_START",
    "id": "{{.CookerID}}"
  },
  "requestId": "{{index .UUIDs 7}}"
}
