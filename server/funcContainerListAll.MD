# Contanier List All

```url
http://localhost:8081/containersListAll
```

```json
{
  "Length": 6,
  "Limit": 0,
  "Skip": 0,
  "Success": true,
  "Error": [
    
  ],
  "Data": [
    {
      "Id": "cd6a4b5ca00a2e67e1af61def2c2b5ed97d75444f6f0a810dceda3fadf369765",
      "Names": [
        "/cont_server_delete_before_test"
      ],
      "Image": "image_server_delete_before_test:latest",
      "ImageID": "sha256:3ec59e3b6e293db8e3e5536454aa3a493a442d8533e7e8a184a73566a6b4f432",
      "Command": "/main",
      "Created": 1594411766,
      "Ports": [
        {
          "IP": "0.0.0.0",
          "PrivatePort": 8080,
          "PublicPort": 8080,
          "Type": "tcp"
        }
      ],
      "Labels": {
        
      },
      "State": "running",
      "Status": "Up 40 minutes",
      "HostConfig": {
        "NetworkMode": "default"
      },
      "NetworkSettings": {
        "Networks": {
          "network_delete_before_test_b": {
            "IPAMConfig": {
              "IPv4Address": "11.0.0.5"
            },
            "Links": null,
            "Aliases": null,
            "NetworkID": "86e8d30aa686ceddc1ec813674f355c49ad35dd848f05e29d8eb40333b1883ff",
            "EndpointID": "785dff1c511d12517cc760b7f84398879c4535ab01bbbdfa8b225b252ea57aaf",
            "Gateway": "11.0.0.1",
            "IPAddress": "11.0.0.5",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "MacAddress": "02:42:0b:00:00:05",
            "DriverOpts": null
          }
        }
      },
      "Mounts": [
        
      ]
    },
    {
      "Id": "63b43a9b85ef6761d3d9669c331e22bfe13907c3d40fd0c740688192c774ed58",
      "Names": [
        "/container_cc_delete_before_test"
      ],
      "Image": "mongo:3",
      "ImageID": "sha256:944c82bef68bc2ec6a58db644f36ac218304605605b3f65951c5f46e2da2d4d7",
      "Command": "docker-entrypoint.sh mongod",
      "Created": 1594411719,
      "Ports": [
        {
          "PrivatePort": 27017,
          "Type": "tcp"
        }
      ],
      "Labels": {
        
      },
      "State": "running",
      "Status": "Up 40 minutes",
      "HostConfig": {
        "NetworkMode": "default"
      },
      "NetworkSettings": {
        "Networks": {
          "network_delete_before_test_b": {
            "IPAMConfig": {
              "IPv4Address": "11.0.0.4"
            },
            "Links": null,
            "Aliases": null,
            "NetworkID": "86e8d30aa686ceddc1ec813674f355c49ad35dd848f05e29d8eb40333b1883ff",
            "EndpointID": "0c4790d83f0cf1fbd42195f8aa08a8f4a6020e3f80c4d6ebf792de2dda9e69d9",
            "Gateway": "11.0.0.1",
            "IPAddress": "11.0.0.4",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "MacAddress": "02:42:0b:00:00:04",
            "DriverOpts": null
          }
        }
      },
      "Mounts": [
        {
          "Type": "volume",
          "Name": "4e880f9ad2c334207bf68138c340a8945ad4e587e2a28d505b67789f97b011d7",
          "Source": "",
          "Destination": "/data/configdb",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "volume",
          "Name": "914f067c48dc2a0f710ae45bb7976cba24133d69f33f3f9cfdf4aa80f90a9301",
          "Source": "",
          "Destination": "/data/db",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "bind",
          "Source": "/host_mnt/c/Users/helmut.kemper.CS/GolandProjects/apagar/config.conf",
          "Destination": "/etc/mongo.conf",
          "Mode": "",
          "RW": true,
          "Propagation": "rprivate"
        }
      ]
    },
    {
      "Id": "e7d44c8f0495312589a0d9222cf12f821da5b01f65ee4e78f7ce47fc64f6486e",
      "Names": [
        "/container_bb_delete_before_test"
      ],
      "Image": "mongo:3",
      "ImageID": "sha256:944c82bef68bc2ec6a58db644f36ac218304605605b3f65951c5f46e2da2d4d7",
      "Command": "docker-entrypoint.sh mongod",
      "Created": 1594411715,
      "Ports": [
        {
          "PrivatePort": 27017,
          "Type": "tcp"
        }
      ],
      "Labels": {
        
      },
      "State": "running",
      "Status": "Up 40 minutes",
      "HostConfig": {
        "NetworkMode": "default"
      },
      "NetworkSettings": {
        "Networks": {
          "network_delete_before_test_b": {
            "IPAMConfig": {
              "IPv4Address": "11.0.0.3"
            },
            "Links": null,
            "Aliases": null,
            "NetworkID": "86e8d30aa686ceddc1ec813674f355c49ad35dd848f05e29d8eb40333b1883ff",
            "EndpointID": "8f1b1ebc7f81a30536f685f5eadca34353e230fb7d83032c344d2bca4b2ef406",
            "Gateway": "11.0.0.1",
            "IPAddress": "11.0.0.3",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "MacAddress": "02:42:0b:00:00:03",
            "DriverOpts": null
          }
        }
      },
      "Mounts": [
        {
          "Type": "volume",
          "Name": "cbbdbea801a51668145cb4baf7d25bf91fb8788397b5e7eb3262cfa0cd7ce50c",
          "Source": "",
          "Destination": "/data/configdb",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "volume",
          "Name": "ed90409aed1527f7b0f48fc6645dd47a87ec43ac8a62a61421e78da83be87015",
          "Source": "",
          "Destination": "/data/db",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "bind",
          "Source": "/host_mnt/c/Users/helmut.kemper.CS/GolandProjects/apagar/config.conf",
          "Destination": "/etc/mongo.conf",
          "Mode": "",
          "RW": true,
          "Propagation": "rprivate"
        }
      ]
    },
    {
      "Id": "cbc319cfc931ace3883b6f80a76a4fb3296e2d52f2d3349df2a64dae23eb3a0c",
      "Names": [
        "/container_aa_delete_before_test"
      ],
      "Image": "mongo:3",
      "ImageID": "sha256:944c82bef68bc2ec6a58db644f36ac218304605605b3f65951c5f46e2da2d4d7",
      "Command": "docker-entrypoint.sh mongod",
      "Created": 1594411711,
      "Ports": [
        {
          "PrivatePort": 27017,
          "Type": "tcp"
        }
      ],
      "Labels": {
        
      },
      "State": "running",
      "Status": "Up 40 minutes",
      "HostConfig": {
        "NetworkMode": "default"
      },
      "NetworkSettings": {
        "Networks": {
          "network_delete_before_test_b": {
            "IPAMConfig": {
              "IPv4Address": "11.0.0.2"
            },
            "Links": null,
            "Aliases": null,
            "NetworkID": "86e8d30aa686ceddc1ec813674f355c49ad35dd848f05e29d8eb40333b1883ff",
            "EndpointID": "7bbabcae76bccdc17fcb6bf9238cd3c4cbb367d4aca73ffe16468bffa9dc9116",
            "Gateway": "11.0.0.1",
            "IPAddress": "11.0.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "MacAddress": "02:42:0b:00:00:02",
            "DriverOpts": null
          }
        }
      },
      "Mounts": [
        {
          "Type": "volume",
          "Name": "0effbc14e66c009851e98edae65fa520d850143faf11b4594fda403aec9111df",
          "Source": "",
          "Destination": "/data/configdb",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "volume",
          "Name": "4534fab9513dcfd31db263ffb0db1372bd3ebcc72d65a86683b09a237fc10173",
          "Source": "",
          "Destination": "/data/db",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "bind",
          "Source": "/host_mnt/c/Users/helmut.kemper.CS/GolandProjects/apagar/config.conf",
          "Destination": "/etc/mongo.conf",
          "Mode": "",
          "RW": true,
          "Propagation": "rprivate"
        }
      ]
    },
    {
      "Id": "9741d024136301a9314055efdf3632c4b925036cf359f5a9e754a310ef5fb5e0",
      "Names": [
        "/delete_____a"
      ],
      "Image": "delete:latest",
      "ImageID": "sha256:b914df731dd6ddedc44d3dcd1a733a33411c3617894f80046eff94dba821ae7a",
      "Command": "./main",
      "Created": 1594409696,
      "Ports": [
        {
          "IP": "0.0.0.0",
          "PrivatePort": 3000,
          "PublicPort": 3000,
          "Type": "tcp"
        }
      ],
      "Labels": {
        
      },
      "State": "running",
      "Status": "Up 40 minutes",
      "HostConfig": {
        "NetworkMode": "default"
      },
      "NetworkSettings": {
        "Networks": {
          "bridge": {
            "IPAMConfig": null,
            "Links": null,
            "Aliases": null,
            "NetworkID": "4767ad67fdde556ed3dd88e5d73a9fafd0357d1b669ac2fb16ac9c2dcc65a435",
            "EndpointID": "0707480e98908edbcc240f69283e3ce25cababfcf215b159a8eeba4030b7d658",
            "Gateway": "172.17.0.1",
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "MacAddress": "02:42:ac:11:00:02",
            "DriverOpts": null
          }
        }
      },
      "Mounts": [
        {
          "Type": "bind",
          "Source": "/var/run/docker.sock",
          "Destination": "/var/run/docker.sock",
          "Mode": "",
          "RW": true,
          "Propagation": "rprivate"
        }
      ]
    },
    {
      "Id": "6162ec603638ebb1edb50cc5bfdf7292fb232c5e384e0841eb5d0080f339be68",
      "Names": [
        "/container_b_delete_before_test"
      ],
      "Image": "mongo:latest",
      "ImageID": "sha256:6d11486a97a77beaad31f63463a744dc3070ed4070bd15a695898a171f349441",
      "Command": "docker-entrypoint.sh mongod",
      "Created": 1594388580,
      "Ports": [
        {
          "IP": "0.0.0.0",
          "PrivatePort": 27017,
          "PublicPort": 27017,
          "Type": "tcp"
        }
      ],
      "Labels": {
        
      },
      "State": "running",
      "Status": "Up 40 minutes",
      "HostConfig": {
        "NetworkMode": "default"
      },
      "NetworkSettings": {
        "Networks": {
          "bridge": {
            "IPAMConfig": null,
            "Links": null,
            "Aliases": null,
            "NetworkID": "4767ad67fdde556ed3dd88e5d73a9fafd0357d1b669ac2fb16ac9c2dcc65a435",
            "EndpointID": "cedbb3b84a3bacc3c8243edeb1b8ea3c2bbe49f82a644151ed44357c166c8917",
            "Gateway": "172.17.0.1",
            "IPAddress": "172.17.0.3",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "MacAddress": "02:42:ac:11:00:03",
            "DriverOpts": null
          }
        }
      },
      "Mounts": [
        {
          "Type": "volume",
          "Name": "82ce0b093414575ec8b295c5143c61a16844db85b0379b69c28389a5afc438da",
          "Source": "",
          "Destination": "/data/configdb",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "volume",
          "Name": "1b9a4f14103a20b8bc902a070f67fdafc483a0bb60b3d9bc67c29f5aff0de1e6",
          "Source": "",
          "Destination": "/data/db",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "bind",
          "Source": "/host_mnt/c/Users/helmut.kemper.CS/GolandProjects/apagar/config.conf",
          "Destination": "/etc/mongo.conf",
          "Mode": "",
          "RW": true,
          "Propagation": "rprivate"
        }
      ]
    }
  ]
}
```