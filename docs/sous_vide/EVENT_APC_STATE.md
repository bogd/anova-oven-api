# EVENT_APC_STATE

Received periodically (every 30 seconds when device is idle, every 2 seconds when the device is running). 

Syntax:
```
{
  "command": "EVENT_APC_STATE",
  "payload": {
    "cookerId": "<your_cooker_id>",
    "type": "pro",
    "state": {
      "audio-control": {
        "file-name": "",
        "volume": 24
      },
      "boot-id": "",
      "cap-touch": {
        "minus-button": 0,
        "play-button": 0,
        "plus-button": 0,
        "target-temperature-button": 0,
        "timer-button": 0,
        "water-temperature-button": 0
      },
      "heater-control": {
        "duty-cycle": 99.9
      },
      "job": {
        "cook-time-seconds": 6960,
        "id": "<random_22_digit_hex>",
        "mode": "COOK",
        "ota-url": "",
        "target-temperature": 69,
        "temperature-unit": "C"
      },
      "job-status": {
        "cook-time-remaining": 6960,
        "job-start-systick": 146442132,
        "provisioning-pairing-code": 0,
        "state": "PREHEATING",
        "state-change-systick": 146442132
      },
      "motor-control": {
        "duty-cycle": 74.925
      },
      "motor-info": {
        "rpm": 0
      },
      "network-info": {
        "bssid": "<bssid>",
        "connection-status": "connected-station",
        "is-provisioning": false,
        "mac-address": "<cooker-mac>",
        "mode": "station",
        "security-type": "WPA2",
        "ssid": "<your_ssid>"
      },
      "pin-info": {
        "device-safe": 1,
        "motor-stuck": 0,
        "water-leak": 0,
        "water-level-critical": 0,
        "water-level-low": 0
      },
      "system-info-2640": {
        "firmware-version": "3.3.01",
        "firmware-version-sha": "3.3.01",
        "largest-free-heap-size": 1208,
        "mcu-temperature": 23,
        "systick": 146450562,
        "total-free-heap-size": 1208,
        "total-heap-size": 16384
      },
      "system-info-3220": {
        "firmware-version": "3.3.01",
        "firmware-version-sha": "3.3.01",
        "fwUpgradeStatus": 0,
        "has-real-cert-catalog": true,
        "largest-free-heap-size": 20432,
        "systick": 146419510,
        "total-free-heap-size": 21376,
        "total-heap-size": 65536
      },
      "temperature-info": {
        "heater-temperature": 5.42,
        "triac-temperature": 11,
        "water-temperature": 4.10
      }
    }
  }
}
```