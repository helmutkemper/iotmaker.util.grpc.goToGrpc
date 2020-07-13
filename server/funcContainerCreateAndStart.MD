# Container Create And Start

```url
http://localhost:8081/containerCreateAndStart
```

```json
{
  "ImageName": "mongo:latest",
  "ContainerName": "new_mondo_delete_before_test",
  "Rest artPolicy": 3,
  "MountVolumes": [
    {
      "Type": "bind",
      "Source": "c:/data/db",
      "Target": "/data/db",
      "ReadOnly": false,
      "Consistency": "consistent"
    }
  ],
  "NetworkName": "",
  "PortList": [
    "27017/tcp"
  ]
}
```

```json
{
  "Length": 1,
  "Limit": 0,
  "Skip": 0,
  "Success": true,
  "Error": [
    
  ],
  "Data": [
    {
      "ID": "d25ccebf831b52eb6cf834dbcd8a970a73b6e33d07a0cf17add0261d3712d5f6"
    }
  ]
}
```