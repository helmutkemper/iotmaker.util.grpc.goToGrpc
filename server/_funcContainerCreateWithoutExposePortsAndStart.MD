# Container Create Without Expose Ports And Start

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
      "ID": "e786cc606ec33ccb7ce9416224479e6b85f52eeedad31deeb868126c1fd5ee02"
    }
  ]
}
```