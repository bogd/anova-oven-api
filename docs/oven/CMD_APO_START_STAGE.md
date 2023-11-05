# CMD_APO_START_STAGE

Start another stage of the current cook. You will need the stage ID (set when you ordered the cook, or gathered from an `EVENT_APO_STATE` message).

Syntax:
```
{
  "command": "CMD_APO_START_STAGE",
  "payload": {
    "payload": {
      "stageId": "android-<uuid>"                  # Needs to match an existing `stageId`
    },
    "type": "CMD_APO_START_STAGE",
    "id": "0123456789abcdef"
  },
  "requestId": "<uuid>"                             # Your oven ID
} 
```

[Example](../examples/CMD_APO_START_STAGE.json)