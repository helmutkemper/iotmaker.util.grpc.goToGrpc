# Container Create And Change Exposed Port

```url
http://localhost:8081/containerCreateAndChangeExposedPort
```

Post data:
```json
{
  "ImageName": "mongo:latest",
  "ContainerName": "new_mondo_delete_before_test",
  "Rest artPolicy": 3,
  "MountVolumes": [],
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
      "ID": "72edb7342a46aad127cd3c8e072f73855c67a2bc2becad4cac191ab00377354e"
    }
  ]
}
```