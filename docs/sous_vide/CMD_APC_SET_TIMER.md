# CMD_APC_SET_TIMER

Set the cook time (for an idle device) or modify the existing timer (for a running device)

Syntax:
```
{
  "command": "CMD_APC_SET_TIMER",
  "requestId": "<random_22_digit_hex>",
  "payload": {
    "cookerId": "<your_cooker_id>",
    "type": "pro",
    "timer": <timer_in_seconds>,
    "requestId": "<random_22_digit_hex>"
  }
}
```

[Full command example](../examples/CMD_APC_SET_TIMER.json)
