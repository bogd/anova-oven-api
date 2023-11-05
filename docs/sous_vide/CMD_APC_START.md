# CMD_APC_START

Starts a cooking process with the specified temperature and timer

Syntax:
```
{
  "command": "CMD_APC_START",
  "requestId": "<random_22_digit_hex>",
  "payload": {
    "cookerId": "<your_cooker_id>",         # From an `EVENT_APC_STATE` or from the app
    "type": "pro",
    "targetTemperature": <temp>,            
    "unit": "C",                            # C or F
    "timer": <timer_in_seconds>,
    "requestId": "<random_22_digit_hex>"
  }
}
```

[Full command example](../examples/CMD_APC_START.json).