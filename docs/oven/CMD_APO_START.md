# CMD_APO_START


Start a cooking process with multiple stages.

> [!NOTE]
> It appears that the syntax is different for v2 ovens. I do not have the hardware to check, but you can find a [full command sample here](../examples/CMD_APO_START_v2.json). Credits go to @jasonjei, who [provided](#4) the sample.

Syntax:
```
{
  "command": "CMD_APO_START",
  "payload": {
    "payload": {
      "cookId": "android-<uuid>",  # Or ios-<uuid>
      "stages": [
        <stage_list>
      ]
    },
    "type": "CMD_APO_START",
    "id": "<your_oven_id>"
  },
  "requestId": "<uuid>"
}  
```

[Stage syntax](./stage.md)

[Full command example](../examples/CMD_APO_START.json). Notice that a stage with a timer is actually represented by two stages in the JSON (pretty much similar to the way you build stages in the app), to allow for the "start timer after preheat" and the "start timer manually" options.