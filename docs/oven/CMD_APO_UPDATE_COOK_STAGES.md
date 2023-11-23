# CMD_APO_UPDATE_COOK_STAGES

Update/modify the full list of cook stages.

Syntax:

```
{
  "command": "CMD_APO_UPDATE_COOK_STAGES",
  "payload": {
    "payload": {
      "stages": [
		<list of stages - see stage syntax>
      ]
    },
    "type": "CMD_APO_UPDATE_COOK_STAGES",
    "id": "<your_oven_id>"
  },
  "requestId": "<uuid>"
}

```

[Stage syntax](./stage.md)

[Example](../examples/CMD_APO_UPDATE_COOK_STAGES.json)