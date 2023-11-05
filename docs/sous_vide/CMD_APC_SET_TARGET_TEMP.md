# CMD_APC_SET_TARGET_TEMP

Change the set target temperature. Works for both idle devices, and devices currently running.

Syntax:
```
{
    "command": "CMD_APC_SET_TARGET_TEMP",
    "requestId": "<random_22_digit_hex>",
    "payload": {
      "cookerId": "<your_cooker_id>",
      "type": "pro",
      "targetTemperature": <temperature>,
      "unit": "C",
      "requestId": "<random_22_digit_hex>"
    }
  }
```

[Full command example](../examples/CMD_APC_SET_TARGET_TEMP.json)

