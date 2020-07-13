# Container Create

```url
http://localhost:8081/containerCreate
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
      "ID": "dbcbb1afd2deb8c703a0793d7f092aa89eba99ed3f0625344e1384620f2c9a93"
    }
  ]
}
```