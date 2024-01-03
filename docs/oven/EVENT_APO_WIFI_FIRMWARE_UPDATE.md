# EVENT_APO_WIFI_FIRMWARE_UPDATE

Looks like it sends when the oven has an update to apply. Weirdly the release notes aren't public.

```
{
   "command" : "EVENT_APO_WIFI_FIRMWARE_UPDATE",
   "payload" : {
      "cookerId" : "01234XXXXXXXXX",
      "ota" : {
         "available" : true,
         "description" : "2.1.8 - Prevent cooks when idle or NTC broken https://www.notion.so/anovaculinary/2-1-8-Prevent-cooks-when-idle-or-NTC-broken-5b6491b9e0d643478c5015f0a667a76a?pvs=4",
         "required" : false,
         "url" : "https://storage.googleapis.com/anova-app.appspot.com/oven-firmware/oven-controller-2.1.8.bin",
         "version" : "2.1.8"
      },
      "type" : "oven_v1",
      "version" : "2.1.7"
   }
}
```
