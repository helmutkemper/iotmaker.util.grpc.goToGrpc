# Container Create Change Exposed Port And Start

```url
http://localhost:8081/containerCreateChangeExposedPortAndStart
```

Post data:
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
  "ContainerNetwork": "network",
  "CurrentPort": [
    "27017/tcp"
  ],
  "ChangeToPort": [
    "27016/tcp"
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
      "ID": "d9af9d3454c3ead435361cc9f16717d219a2caf745c5d90ebac44d68a131e608"
    }
  ]
}
```