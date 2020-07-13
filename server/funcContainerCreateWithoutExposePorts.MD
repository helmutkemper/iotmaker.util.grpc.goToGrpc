# Container Create Without Expose Ports

```url
http://localhost:8081/containerCreateWithoutExposePorts
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
  "NetworkName": ""
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
      "ID": "f481b56140214ab151984e7979716263136248b4a19af9046a6ce5f08980ef3b"
    }
  ]
}
```