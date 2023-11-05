# CMD_APC_STOP

Stops the current cooking process. 

Syntax:

```
{
  "command": "CMD_APC_STOP",
  "requestId": ""<random_22_digit_hex>"",
  "payload": {
    "cookerId": "<your_cooker_id>",
    "type": "<cooker_type>",                    # "pro", "nano", etc
    "requestId": ""<random_22_digit_hex>""
  }
}
```

[Example](../examples/CMD_APC_STOP.json)