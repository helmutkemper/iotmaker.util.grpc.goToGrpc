# Network List All

```url
http://localhost:8081/networkListAll
```

```json
{
  "Length": 5,
  "Limit": 0,
  "Skip": 0,
  "Success": true,
  "Error": [
    
  ],
  "Data": [
    {
      "Name": "network_delete_before_test",
      "Id": "866b8c43a7b262712112708c3e0e340e3dba8824c69358d98cb40c7e3b419467",
      "Created": "2020-07-10T19:28:10.67675937Z",
      "Scope": "local",
      "Driver": "bridge",
      "EnableIPv6": false,
      "IPAM": {
        "Driver": "default",
        "Options": null,
        "Config": [
          {
            "Subnet": "10.0.0.0/16",
            "Gateway": "10.0.0.1"
          }
        ]
      },
      "Internal": false,
      "Attachable": true,
      "Ingress": false,
      "ConfigFrom": {
        "Network": ""
      },
      "ConfigOnly": false,
      "Containers": {
        
      },
      "Options": {
        
      },
      "Labels": {
        "name": "network_delete_before_test"
      }
    },
    {
      "Name": "network_delete_before_test_b",
      "Id": "86e8d30aa686ceddc1ec813674f355c49ad35dd848f05e29d8eb40333b1883ff",
      "Created": "2020-07-10T20:08:26.391730078Z",
      "Scope": "local",
      "Driver": "bridge",
      "EnableIPv6": false,
      "IPAM": {
        "Driver": "default",
        "Options": null,
        "Config": [
          {
            "Subnet": "11.0.0.0/16",
            "Gateway": "11.0.0.1"
          }
        ]
      },
      "Internal": false,
      "Attachable": true,
      "Ingress": false,
      "ConfigFrom": {
        "Network": ""
      },
      "ConfigOnly": false,
      "Containers": {
        
      },
      "Options": {
        
      },
      "Labels": {
        "name": "network_delete_before_test_b"
      }
    },
    {
      "Name": "none",
      "Id": "e15b5d3c7c0cb63f77ee075544b5c7b7a4afa30ba080dc7c8bdfc54f0a2390f2",
      "Created": "2020-06-23T20:36:55.386452816Z",
      "Scope": "local",
      "Driver": "null",
      "EnableIPv6": false,
      "IPAM": {
        "Driver": "default",
        "Options": null,
        "Config": [
          
        ]
      },
      "Internal": false,
      "Attachable": false,
      "Ingress": false,
      "ConfigFrom": {
        "Network": ""
      },
      "ConfigOnly": false,
      "Containers": {
        
      },
      "Options": {
        
      },
      "Labels": {
        
      }
    },
    {
      "Name": "bridge",
      "Id": "4767ad67fdde556ed3dd88e5d73a9fafd0357d1b669ac2fb16ac9c2dcc65a435",
      "Created": "2020-07-11T16:12:34.120664454Z",
      "Scope": "local",
      "Driver": "bridge",
      "EnableIPv6": false,
      "IPAM": {
        "Driver": "default",
        "Options": null,
        "Config": [
          {
            "Subnet": "172.17.0.0/16",
            "Gateway": "172.17.0.1"
          }
        ]
      },
      "Internal": false,
      "Attachable": false,
      "Ingress": false,
      "ConfigFrom": {
        "Network": ""
      },
      "ConfigOnly": false,
      "Containers": {
        
      },
      "Options": {
        "com.docker.network.bridge.default_bridge": "true",
        "com.docker.network.bridge.enable_icc": "true",
        "com.docker.network.bridge.enable_ip_masquerade": "true",
        "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
        "com.docker.network.bridge.name": "docker0",
        "com.docker.network.driver.mtu": "1500"
      },
      "Labels": {
        
      }
    },
    {
      "Name": "host",
      "Id": "4dae06c212e7cee674be37cb3fbf541897ba29b86bf41866fd28e08d397f1271",
      "Created": "2020-06-23T20:36:55.420086816Z",
      "Scope": "local",
      "Driver": "host",
      "EnableIPv6": false,
      "IPAM": {
        "Driver": "default",
        "Options": null,
        "Config": [
          
        ]
      },
      "Internal": false,
      "Attachable": false,
      "Ingress": false,
      "ConfigFrom": {
        "Network": ""
      },
      "ConfigOnly": false,
      "Containers": {
        
      },
      "Options": {
        
      },
      "Labels": {
        
      }
    }
  ]
}
```