# Cook stage syntax

```
{
    "stepType": "stage",
    "id": "android-<uuid>",         # Or ios-<uuid>
    "title": "First stage",         # The name of the stage (as set in the app)
    "description": "",
    "type": "preheat",
    "userActionRequired": false,    # "false" if the stage starts automatically, "true" if it needs to be started manually
    "temperatureBulbs": {
        "dry": {                    # "dry" or "wet", depending on "mode"
            "setpoint": {           # Set temperature, in both Fahrenheit and Celsius. Unknown which one takes precedence if they differ!
                "fahrenheit": 410,
                "celsius": 210
            }
        },
        "mode": "dry"               # "sous-vide mode: on" == "mode: wet"; "sous-vide mode: off" == "mode: dry"
    },
    "heatingElements": {            # What heating elements are activated
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
    "fan": {                        # Fan speed
        "speed": 100
    },
    "vent": {                       # Unknown
        "open": false
    },
    "rackPosition": 3               # Tray position
}
```