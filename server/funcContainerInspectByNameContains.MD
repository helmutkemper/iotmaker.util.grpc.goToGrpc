# Container Inspect By Name Contains

```url
http://localhost:8081/containerInspectByNameContains?Name=delete
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
      "Created": "2020-07-10T20:09:26.198229388Z",
      "Path": "/main",
      "Args": [
        
      ],
      "State": {
        "Status": "running",
        "Running": true,
        "Paused": false,
        "Restarting": false,
        "OOMKilled": false,
        "Dead": false,
        "Pid": 2389,
        "ExitCode": 0,
        "Error": "",
        "StartedAt": "2020-07-11T20:54:15.807097609Z",
        "FinishedAt": "2020-07-11T20:54:14.195985509Z"
      },
      "Image": "sha256:3ec59e3b6e293db8e3e5536454aa3a493a442d8533e7e8a184a73566a6b4f432",
      "ResolvConfPath": "/var/lib/docker/containers/cd6a4b5ca00a2e67e1af61def2c2b5ed97d75444f6f0a810dceda3fadf369765/resolv.conf",
      "HostnamePath": "/var/lib/docker/containers/cd6a4b5ca00a2e67e1af61def2c2b5ed97d75444f6f0a810dceda3fadf369765/hostname",
      "HostsPath": "/var/lib/docker/containers/cd6a4b5ca00a2e67e1af61def2c2b5ed97d75444f6f0a810dceda3fadf369765/hosts",
      "LogPath": "/var/lib/docker/containers/cd6a4b5ca00a2e67e1af61def2c2b5ed97d75444f6f0a810dceda3fadf369765/cd6a4b5ca00a2e67e1af61def2c2b5ed97d75444f6f0a810dceda3fadf369765-json.log",
      "Name": "/cont_server_delete_before_test",
      "RestartCount": 0,
      "Driver": "overlay2",
      "Platform": "linux",
      "MountLabel": "",
      "ProcessLabel": "",
      "AppArmorProfile": "",
      "ExecIDs": null,
      "HostConfig": {
        "Binds": null,
        "ContainerIDFile": "",
        "LogConfig": {
          "Type": "json-file",
          "Config": {
            
          }
        },
        "NetworkMode": "default",
        "PortBindings": {
          "8080/tcp": [
            {
              "HostIp": "",
              "HostPort": "8080/tcp"
            }
          ]
        },
        "RestartPolicy": {
          "Name": "unless-stopped",
          "MaximumRetryCount": 0
        },
        "AutoRemove": false,
        "VolumeDriver": "",
        "VolumesFrom": null,
        "CapAdd": null,
        "CapDrop": null,
        "Capabilities": null,
        "Dns": [
          
        ],
        "DnsOptions": [
          
        ],
        "DnsSearch": [
          
        ],
        "ExtraHosts": null,
        "GroupAdd": null,
        "IpcMode": "private",
        "Cgroup": "",
        "Links": null,
        "OomScoreAdj": 0,
        "PidMode": "",
        "Privileged": false,
        "PublishAllPorts": false,
        "ReadonlyRootfs": false,
        "SecurityOpt": null,
        "UTSMode": "",
        "UsernsMode": "",
        "ShmSize": 67108864,
        "Runtime": "runc",
        "ConsoleSize": [
          0,
          0
        ],
        "Isolation": "",
        "CpuShares": 0,
        "Memory": 0,
        "NanoCpus": 0,
        "CgroupParent": "",
        "BlkioWeight": 0,
        "BlkioWeightDevice": null,
        "BlkioDeviceReadBps": null,
        "BlkioDeviceWriteBps": null,
        "BlkioDeviceReadIOps": null,
        "BlkioDeviceWriteIOps": null,
        "CpuPeriod": 0,
        "CpuQuota": 0,
        "CpuRealtimePeriod": 0,
        "CpuRealtimeRuntime": 0,
        "CpusetCpus": "",
        "CpusetMems": "",
        "Devices": null,
        "DeviceCgroupRules": null,
        "DeviceRequests": null,
        "KernelMemory": 0,
        "KernelMemoryTCP": 0,
        "MemoryReservation": 0,
        "MemorySwap": 0,
        "MemorySwappiness": null,
        "OomKillDisable": false,
        "PidsLimit": null,
        "Ulimits": null,
        "CpuCount": 0,
        "CpuPercent": 0,
        "IOMaximumIOps": 0,
        "IOMaximumBandwidth": 0,
        "MaskedPaths": [
          "/proc/asound",
          "/proc/acpi",
          "/proc/kcore",
          "/proc/keys",
          "/proc/latency_stats",
          "/proc/timer_list",
          "/proc/timer_stats",
          "/proc/sched_debug",
          "/proc/scsi",
          "/sys/firmware"
        ],
        "ReadonlyPaths": [
          "/proc/bus",
          "/proc/fs",
          "/proc/irq",
          "/proc/sys",
          "/proc/sysrq-trigger"
        ]
      },
      "GraphDriver": {
        "Data": {
          "LowerDir": "/var/lib/docker/overlay2/9595da6b9aca7cd010aa96e21ee3cc6169b14065cd2f5476c3b2e2df9e78aee1-init/diff:/var/lib/docker/overlay2/739391116eef2713b8bd05a9daa826b7a4872cd04346b6d8e8bc913a2fa6dcbc/diff",
          "MergedDir": "/var/lib/docker/overlay2/9595da6b9aca7cd010aa96e21ee3cc6169b14065cd2f5476c3b2e2df9e78aee1/merged",
          "UpperDir": "/var/lib/docker/overlay2/9595da6b9aca7cd010aa96e21ee3cc6169b14065cd2f5476c3b2e2df9e78aee1/diff",
          "WorkDir": "/var/lib/docker/overlay2/9595da6b9aca7cd010aa96e21ee3cc6169b14065cd2f5476c3b2e2df9e78aee1/work"
        },
        "Name": "overlay2"
      },
      "Mounts": [
        
      ],
      "Config": {
        "Hostname": "cd6a4b5ca00a",
        "Domainname": "",
        "User": "",
        "AttachStdin": false,
        "AttachStdout": false,
        "AttachStderr": false,
        "ExposedPorts": {
          "8080/tcp": {
            
          }
        },
        "Tty": false,
        "OpenStdin": false,
        "StdinOnce": false,
        "Env": [
          "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
        ],
        "Cmd": [
          "/main"
        ],
        "Image": "image_server_delete_before_test:latest",
        "Volumes": null,
        "WorkingDir": "",
        "Entrypoint": null,
        "OnBuild": null,
        "Labels": {
          
        }
      },
      "NetworkSettings": {
        "Bridge": "",
        "SandboxID": "5f0f62bd1c6dac0dfd584a3619d1544b75ffbe1aa05166143de4fee7d4dc95aa",
        "HairpinMode": false,
        "LinkLocalIPv6Address": "",
        "LinkLocalIPv6PrefixLen": 0,
        "Ports": {
          "8080/tcp": [
            {
              "HostIp": "0.0.0.0",
              "HostPort": "8080"
            }
          ]
        },
        "SandboxKey": "/var/run/docker/netns/5f0f62bd1c6d",
        "SecondaryIPAddresses": null,
        "SecondaryIPv6Addresses": null,
        "EndpointID": "",
        "Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "IPAddress": "",
        "IPPrefixLen": 0,
        "IPv6Gateway": "",
        "MacAddress": "",
        "Networks": {
          "network_delete_before_test_b": {
            "IPAMConfig": {
              "IPv4Address": "11.0.0.5"
            },
            "Links": null,
            "Aliases": [
              "cd6a4b5ca00a"
            ],
            "NetworkID": "86e8d30aa686ceddc1ec813674f355c49ad35dd848f05e29d8eb40333b1883ff",
            "EndpointID": "0da75a95071facfd15c61a75f5360af04bcb8e33ca7655775773c47dd13b7077",
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
      }
    },
    {
      "Id": "63b43a9b85ef6761d3d9669c331e22bfe13907c3d40fd0c740688192c774ed58",
      "Created": "2020-07-10T20:08:39.314530193Z",
      "Path": "docker-entrypoint.sh",
      "Args": [
        "mongod"
      ],
      "State": {
        "Status": "running",
        "Running": true,
        "Paused": false,
        "Restarting": false,
        "OOMKilled": false,
        "Dead": false,
        "Pid": 2078,
        "ExitCode": 0,
        "Error": "",
        "StartedAt": "2020-07-11T20:54:15.647988709Z",
        "FinishedAt": "2020-07-11T20:54:14.180271409Z"
      },
      "Image": "sha256:944c82bef68bc2ec6a58db644f36ac218304605605b3f65951c5f46e2da2d4d7",
      "ResolvConfPath": "/var/lib/docker/containers/63b43a9b85ef6761d3d9669c331e22bfe13907c3d40fd0c740688192c774ed58/resolv.conf",
      "HostnamePath": "/var/lib/docker/containers/63b43a9b85ef6761d3d9669c331e22bfe13907c3d40fd0c740688192c774ed58/hostname",
      "HostsPath": "/var/lib/docker/containers/63b43a9b85ef6761d3d9669c331e22bfe13907c3d40fd0c740688192c774ed58/hosts",
      "LogPath": "/var/lib/docker/containers/63b43a9b85ef6761d3d9669c331e22bfe13907c3d40fd0c740688192c774ed58/63b43a9b85ef6761d3d9669c331e22bfe13907c3d40fd0c740688192c774ed58-json.log",
      "Name": "/container_cc_delete_before_test",
      "RestartCount": 0,
      "Driver": "overlay2",
      "Platform": "linux",
      "MountLabel": "",
      "ProcessLabel": "",
      "AppArmorProfile": "",
      "ExecIDs": null,
      "HostConfig": {
        "Binds": null,
        "ContainerIDFile": "",
        "LogConfig": {
          "Type": "json-file",
          "Config": {
            
          }
        },
        "NetworkMode": "default",
        "PortBindings": null,
        "RestartPolicy": {
          "Name": "on-failure",
          "MaximumRetryCount": 0
        },
        "AutoRemove": false,
        "VolumeDriver": "",
        "VolumesFrom": null,
        "CapAdd": null,
        "CapDrop": null,
        "Capabilities": null,
        "Dns": [
          
        ],
        "DnsOptions": [
          
        ],
        "DnsSearch": [
          
        ],
        "ExtraHosts": null,
        "GroupAdd": null,
        "IpcMode": "private",
        "Cgroup": "",
        "Links": null,
        "OomScoreAdj": 0,
        "PidMode": "",
        "Privileged": false,
        "PublishAllPorts": false,
        "ReadonlyRootfs": false,
        "SecurityOpt": null,
        "UTSMode": "",
        "UsernsMode": "",
        "ShmSize": 67108864,
        "Runtime": "runc",
        "ConsoleSize": [
          0,
          0
        ],
        "Isolation": "",
        "CpuShares": 0,
        "Memory": 0,
        "NanoCpus": 0,
        "CgroupParent": "",
        "BlkioWeight": 0,
        "BlkioWeightDevice": null,
        "BlkioDeviceReadBps": null,
        "BlkioDeviceWriteBps": null,
        "BlkioDeviceReadIOps": null,
        "BlkioDeviceWriteIOps": null,
        "CpuPeriod": 0,
        "CpuQuota": 0,
        "CpuRealtimePeriod": 0,
        "CpuRealtimeRuntime": 0,
        "CpusetCpus": "",
        "CpusetMems": "",
        "Devices": null,
        "DeviceCgroupRules": null,
        "DeviceRequests": null,
        "KernelMemory": 0,
        "KernelMemoryTCP": 0,
        "MemoryReservation": 0,
        "MemorySwap": 0,
        "MemorySwappiness": null,
        "OomKillDisable": false,
        "PidsLimit": null,
        "Ulimits": null,
        "CpuCount": 0,
        "CpuPercent": 0,
        "IOMaximumIOps": 0,
        "IOMaximumBandwidth": 0,
        "Mounts": [
          {
            "Type": "bind",
            "Source": "/host_mnt/c/Users/helmut.kemper.CS/GolandProjects/apagar/config.conf",
            "Target": "/etc/mongo.conf"
          }
        ],
        "MaskedPaths": [
          "/proc/asound",
          "/proc/acpi",
          "/proc/kcore",
          "/proc/keys",
          "/proc/latency_stats",
          "/proc/timer_list",
          "/proc/timer_stats",
          "/proc/sched_debug",
          "/proc/scsi",
          "/sys/firmware"
        ],
        "ReadonlyPaths": [
          "/proc/bus",
          "/proc/fs",
          "/proc/irq",
          "/proc/sys",
          "/proc/sysrq-trigger"
        ]
      },
      "GraphDriver": {
        "Data": {
          "LowerDir": "/var/lib/docker/overlay2/e95891b3eccec4c304b56409bbd9ead0e99f0b3a6c9c8214f0857457eab75c3b-init/diff:/var/lib/docker/overlay2/37a90fbefeebbfee077e4ca7cfbd2085dfacbea3e44dc8f0db44947c66bf090d/diff:/var/lib/docker/overlay2/412cf94e7f60c552e124d6b141c6a4e96363b4de8a167d89d2a863e7fa4b2252/diff:/var/lib/docker/overlay2/719fec79dedd6add3d5a68bad91c218d497bb039a657539dafb48e6112a60f64/diff:/var/lib/docker/overlay2/2ea298a3d6e01ca4915e5725168d6de8cc36be9d76d61235990776dbf5c297c7/diff:/var/lib/docker/overlay2/072403bd7c974df22e203710d8d32f40c3ee6d272b0b5289c552aeae6d6d7b45/diff:/var/lib/docker/overlay2/a3fa03a0b26d7332fceb569d9b1bbd9cf4f00c810e9c4750a57bcc4ffeccf062/diff:/var/lib/docker/overlay2/60190b45c61567d76d3601922ee05e658db287a82e9d1c3045675d18cb16cba8/diff:/var/lib/docker/overlay2/bac87aadfa1f68b8c228b41ae270ad2ab4bb98521a60a7ccaf8b8c1f0000bb58/diff:/var/lib/docker/overlay2/f97033a63fe4ba4535dd8b19d0b85a9390214fbacc03d91b7b276a77a2ac086f/diff:/var/lib/docker/overlay2/b2720a9b34cc3f0ff201a5935e8bd10a752e5a06f6ccae7e915c8a7663674417/diff:/var/lib/docker/overlay2/7152b4e53bfe00bcaf249338206a873c6a64e6bad385bbc0fd4c6359a56420e9/diff:/var/lib/docker/overlay2/f4d3ccc17cb609724dc83aae85699b2614ac5811140064b6b80efc40f0f79ef6/diff:/var/lib/docker/overlay2/a84680b62c1be0b9c2c2f6fadad18e1cd4355edf4091b6bd1555b212fb7d9af3/diff",
          "MergedDir": "/var/lib/docker/overlay2/e95891b3eccec4c304b56409bbd9ead0e99f0b3a6c9c8214f0857457eab75c3b/merged",
          "UpperDir": "/var/lib/docker/overlay2/e95891b3eccec4c304b56409bbd9ead0e99f0b3a6c9c8214f0857457eab75c3b/diff",
          "WorkDir": "/var/lib/docker/overlay2/e95891b3eccec4c304b56409bbd9ead0e99f0b3a6c9c8214f0857457eab75c3b/work"
        },
        "Name": "overlay2"
      },
      "Mounts": [
        {
          "Type": "volume",
          "Name": "4e880f9ad2c334207bf68138c340a8945ad4e587e2a28d505b67789f97b011d7",
          "Source": "/var/lib/docker/volumes/4e880f9ad2c334207bf68138c340a8945ad4e587e2a28d505b67789f97b011d7/_data",
          "Destination": "/data/configdb",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "volume",
          "Name": "914f067c48dc2a0f710ae45bb7976cba24133d69f33f3f9cfdf4aa80f90a9301",
          "Source": "/var/lib/docker/volumes/914f067c48dc2a0f710ae45bb7976cba24133d69f33f3f9cfdf4aa80f90a9301/_data",
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
      ],
      "Config": {
        "Hostname": "63b43a9b85ef",
        "Domainname": "",
        "User": "",
        "AttachStdin": false,
        "AttachStdout": false,
        "AttachStderr": false,
        "ExposedPorts": {
          "27017/tcp": {
            
          }
        },
        "Tty": false,
        "OpenStdin": false,
        "StdinOnce": false,
        "Env": [
          "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
          "GOSU_VERSION=1.12",
          "JSYAML_VERSION=3.13.1",
          "GPG_KEYS=2930ADAE8CAF5059EE73BB4B58712A2291FA4AD5",
          "MONGO_PACKAGE=mongodb-org",
          "MONGO_REPO=repo.mongodb.org",
          "MONGO_MAJOR=3.6",
          "MONGO_VERSION=3.6.18"
        ],
        "Cmd": [
          "mongod"
        ],
        "Image": "mongo:3",
        "Volumes": {
          "/data/configdb": {
            
          },
          "/data/db": {
            
          }
        },
        "WorkingDir": "",
        "Entrypoint": [
          "docker-entrypoint.sh"
        ],
        "OnBuild": null,
        "Labels": {
          
        }
      },
      "NetworkSettings": {
        "Bridge": "",
        "SandboxID": "684476531e436dc9fcdcb6b565be38850dd556ed2b1c7d93ebb9ef13da0e21cf",
        "HairpinMode": false,
        "LinkLocalIPv6Address": "",
        "LinkLocalIPv6PrefixLen": 0,
        "Ports": {
          "27017/tcp": null
        },
        "SandboxKey": "/var/run/docker/netns/684476531e43",
        "SecondaryIPAddresses": null,
        "SecondaryIPv6Addresses": null,
        "EndpointID": "",
        "Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "IPAddress": "",
        "IPPrefixLen": 0,
        "IPv6Gateway": "",
        "MacAddress": "",
        "Networks": {
          "network_delete_before_test_b": {
            "IPAMConfig": {
              "IPv4Address": "11.0.0.4"
            },
            "Links": null,
            "Aliases": [
              "63b43a9b85ef"
            ],
            "NetworkID": "86e8d30aa686ceddc1ec813674f355c49ad35dd848f05e29d8eb40333b1883ff",
            "EndpointID": "0fb36ff419c3c0253bb4ad7a405671d083571d77221a6c42326197a198257cc0",
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
      }
    },
    {
      "Id": "e7d44c8f0495312589a0d9222cf12f821da5b01f65ee4e78f7ce47fc64f6486e",
      "Created": "2020-07-10T20:08:35.229855076Z",
      "Path": "docker-entrypoint.sh",
      "Args": [
        "mongod"
      ],
      "State": {
        "Status": "running",
        "Running": true,
        "Paused": false,
        "Restarting": false,
        "OOMKilled": false,
        "Dead": false,
        "Pid": 2314,
        "ExitCode": 0,
        "Error": "",
        "StartedAt": "2020-07-11T20:54:15.788371109Z",
        "FinishedAt": "2020-07-11T20:54:14.195287309Z"
      },
      "Image": "sha256:944c82bef68bc2ec6a58db644f36ac218304605605b3f65951c5f46e2da2d4d7",
      "ResolvConfPath": "/var/lib/docker/containers/e7d44c8f0495312589a0d9222cf12f821da5b01f65ee4e78f7ce47fc64f6486e/resolv.conf",
      "HostnamePath": "/var/lib/docker/containers/e7d44c8f0495312589a0d9222cf12f821da5b01f65ee4e78f7ce47fc64f6486e/hostname",
      "HostsPath": "/var/lib/docker/containers/e7d44c8f0495312589a0d9222cf12f821da5b01f65ee4e78f7ce47fc64f6486e/hosts",
      "LogPath": "/var/lib/docker/containers/e7d44c8f0495312589a0d9222cf12f821da5b01f65ee4e78f7ce47fc64f6486e/e7d44c8f0495312589a0d9222cf12f821da5b01f65ee4e78f7ce47fc64f6486e-json.log",
      "Name": "/container_bb_delete_before_test",
      "RestartCount": 0,
      "Driver": "overlay2",
      "Platform": "linux",
      "MountLabel": "",
      "ProcessLabel": "",
      "AppArmorProfile": "",
      "ExecIDs": null,
      "HostConfig": {
        "Binds": null,
        "ContainerIDFile": "",
        "LogConfig": {
          "Type": "json-file",
          "Config": {
            
          }
        },
        "NetworkMode": "default",
        "PortBindings": null,
        "RestartPolicy": {
          "Name": "on-failure",
          "MaximumRetryCount": 0
        },
        "AutoRemove": false,
        "VolumeDriver": "",
        "VolumesFrom": null,
        "CapAdd": null,
        "CapDrop": null,
        "Capabilities": null,
        "Dns": [
          
        ],
        "DnsOptions": [
          
        ],
        "DnsSearch": [
          
        ],
        "ExtraHosts": null,
        "GroupAdd": null,
        "IpcMode": "private",
        "Cgroup": "",
        "Links": null,
        "OomScoreAdj": 0,
        "PidMode": "",
        "Privileged": false,
        "PublishAllPorts": false,
        "ReadonlyRootfs": false,
        "SecurityOpt": null,
        "UTSMode": "",
        "UsernsMode": "",
        "ShmSize": 67108864,
        "Runtime": "runc",
        "ConsoleSize": [
          0,
          0
        ],
        "Isolation": "",
        "CpuShares": 0,
        "Memory": 0,
        "NanoCpus": 0,
        "CgroupParent": "",
        "BlkioWeight": 0,
        "BlkioWeightDevice": null,
        "BlkioDeviceReadBps": null,
        "BlkioDeviceWriteBps": null,
        "BlkioDeviceReadIOps": null,
        "BlkioDeviceWriteIOps": null,
        "CpuPeriod": 0,
        "CpuQuota": 0,
        "CpuRealtimePeriod": 0,
        "CpuRealtimeRuntime": 0,
        "CpusetCpus": "",
        "CpusetMems": "",
        "Devices": null,
        "DeviceCgroupRules": null,
        "DeviceRequests": null,
        "KernelMemory": 0,
        "KernelMemoryTCP": 0,
        "MemoryReservation": 0,
        "MemorySwap": 0,
        "MemorySwappiness": null,
        "OomKillDisable": false,
        "PidsLimit": null,
        "Ulimits": null,
        "CpuCount": 0,
        "CpuPercent": 0,
        "IOMaximumIOps": 0,
        "IOMaximumBandwidth": 0,
        "Mounts": [
          {
            "Type": "bind",
            "Source": "/host_mnt/c/Users/helmut.kemper.CS/GolandProjects/apagar/config.conf",
            "Target": "/etc/mongo.conf"
          }
        ],
        "MaskedPaths": [
          "/proc/asound",
          "/proc/acpi",
          "/proc/kcore",
          "/proc/keys",
          "/proc/latency_stats",
          "/proc/timer_list",
          "/proc/timer_stats",
          "/proc/sched_debug",
          "/proc/scsi",
          "/sys/firmware"
        ],
        "ReadonlyPaths": [
          "/proc/bus",
          "/proc/fs",
          "/proc/irq",
          "/proc/sys",
          "/proc/sysrq-trigger"
        ]
      },
      "GraphDriver": {
        "Data": {
          "LowerDir": "/var/lib/docker/overlay2/802af370ac727c49c7e206533b659798200e423c3286d9e6c7f4bec938bf6660-init/diff:/var/lib/docker/overlay2/37a90fbefeebbfee077e4ca7cfbd2085dfacbea3e44dc8f0db44947c66bf090d/diff:/var/lib/docker/overlay2/412cf94e7f60c552e124d6b141c6a4e96363b4de8a167d89d2a863e7fa4b2252/diff:/var/lib/docker/overlay2/719fec79dedd6add3d5a68bad91c218d497bb039a657539dafb48e6112a60f64/diff:/var/lib/docker/overlay2/2ea298a3d6e01ca4915e5725168d6de8cc36be9d76d61235990776dbf5c297c7/diff:/var/lib/docker/overlay2/072403bd7c974df22e203710d8d32f40c3ee6d272b0b5289c552aeae6d6d7b45/diff:/var/lib/docker/overlay2/a3fa03a0b26d7332fceb569d9b1bbd9cf4f00c810e9c4750a57bcc4ffeccf062/diff:/var/lib/docker/overlay2/60190b45c61567d76d3601922ee05e658db287a82e9d1c3045675d18cb16cba8/diff:/var/lib/docker/overlay2/bac87aadfa1f68b8c228b41ae270ad2ab4bb98521a60a7ccaf8b8c1f0000bb58/diff:/var/lib/docker/overlay2/f97033a63fe4ba4535dd8b19d0b85a9390214fbacc03d91b7b276a77a2ac086f/diff:/var/lib/docker/overlay2/b2720a9b34cc3f0ff201a5935e8bd10a752e5a06f6ccae7e915c8a7663674417/diff:/var/lib/docker/overlay2/7152b4e53bfe00bcaf249338206a873c6a64e6bad385bbc0fd4c6359a56420e9/diff:/var/lib/docker/overlay2/f4d3ccc17cb609724dc83aae85699b2614ac5811140064b6b80efc40f0f79ef6/diff:/var/lib/docker/overlay2/a84680b62c1be0b9c2c2f6fadad18e1cd4355edf4091b6bd1555b212fb7d9af3/diff",
          "MergedDir": "/var/lib/docker/overlay2/802af370ac727c49c7e206533b659798200e423c3286d9e6c7f4bec938bf6660/merged",
          "UpperDir": "/var/lib/docker/overlay2/802af370ac727c49c7e206533b659798200e423c3286d9e6c7f4bec938bf6660/diff",
          "WorkDir": "/var/lib/docker/overlay2/802af370ac727c49c7e206533b659798200e423c3286d9e6c7f4bec938bf6660/work"
        },
        "Name": "overlay2"
      },
      "Mounts": [
        {
          "Type": "bind",
          "Source": "/host_mnt/c/Users/helmut.kemper.CS/GolandProjects/apagar/config.conf",
          "Destination": "/etc/mongo.conf",
          "Mode": "",
          "RW": true,
          "Propagation": "rprivate"
        },
        {
          "Type": "volume",
          "Name": "cbbdbea801a51668145cb4baf7d25bf91fb8788397b5e7eb3262cfa0cd7ce50c",
          "Source": "/var/lib/docker/volumes/cbbdbea801a51668145cb4baf7d25bf91fb8788397b5e7eb3262cfa0cd7ce50c/_data",
          "Destination": "/data/configdb",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "volume",
          "Name": "ed90409aed1527f7b0f48fc6645dd47a87ec43ac8a62a61421e78da83be87015",
          "Source": "/var/lib/docker/volumes/ed90409aed1527f7b0f48fc6645dd47a87ec43ac8a62a61421e78da83be87015/_data",
          "Destination": "/data/db",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        }
      ],
      "Config": {
        "Hostname": "e7d44c8f0495",
        "Domainname": "",
        "User": "",
        "AttachStdin": false,
        "AttachStdout": false,
        "AttachStderr": false,
        "ExposedPorts": {
          "27017/tcp": {
            
          }
        },
        "Tty": false,
        "OpenStdin": false,
        "StdinOnce": false,
        "Env": [
          "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
          "GOSU_VERSION=1.12",
          "JSYAML_VERSION=3.13.1",
          "GPG_KEYS=2930ADAE8CAF5059EE73BB4B58712A2291FA4AD5",
          "MONGO_PACKAGE=mongodb-org",
          "MONGO_REPO=repo.mongodb.org",
          "MONGO_MAJOR=3.6",
          "MONGO_VERSION=3.6.18"
        ],
        "Cmd": [
          "mongod"
        ],
        "Image": "mongo:3",
        "Volumes": {
          "/data/configdb": {
            
          },
          "/data/db": {
            
          }
        },
        "WorkingDir": "",
        "Entrypoint": [
          "docker-entrypoint.sh"
        ],
        "OnBuild": null,
        "Labels": {
          
        }
      },
      "NetworkSettings": {
        "Bridge": "",
        "SandboxID": "6261fee5f68eff4b02f06a92bb4ccaac3e71a2d3fc174d1929cbd050975c6e9d",
        "HairpinMode": false,
        "LinkLocalIPv6Address": "",
        "LinkLocalIPv6PrefixLen": 0,
        "Ports": {
          "27017/tcp": null
        },
        "SandboxKey": "/var/run/docker/netns/6261fee5f68e",
        "SecondaryIPAddresses": null,
        "SecondaryIPv6Addresses": null,
        "EndpointID": "",
        "Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "IPAddress": "",
        "IPPrefixLen": 0,
        "IPv6Gateway": "",
        "MacAddress": "",
        "Networks": {
          "network_delete_before_test_b": {
            "IPAMConfig": {
              "IPv4Address": "11.0.0.3"
            },
            "Links": null,
            "Aliases": [
              "e7d44c8f0495"
            ],
            "NetworkID": "86e8d30aa686ceddc1ec813674f355c49ad35dd848f05e29d8eb40333b1883ff",
            "EndpointID": "5ed8198d6da9fef761e2f14c9cab2aaed8c81b4bbf9f74b675983f0ff6e4b1f6",
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
      }
    },
    {
      "Id": "cbc319cfc931ace3883b6f80a76a4fb3296e2d52f2d3349df2a64dae23eb3a0c",
      "Created": "2020-07-10T20:08:31.533510064Z",
      "Path": "docker-entrypoint.sh",
      "Args": [
        "mongod"
      ],
      "State": {
        "Status": "running",
        "Running": true,
        "Paused": false,
        "Restarting": false,
        "OOMKilled": false,
        "Dead": false,
        "Pid": 2229,
        "ExitCode": 0,
        "Error": "",
        "StartedAt": "2020-07-11T20:54:15.711697709Z",
        "FinishedAt": "2020-07-11T20:54:14.196102009Z"
      },
      "Image": "sha256:944c82bef68bc2ec6a58db644f36ac218304605605b3f65951c5f46e2da2d4d7",
      "ResolvConfPath": "/var/lib/docker/containers/cbc319cfc931ace3883b6f80a76a4fb3296e2d52f2d3349df2a64dae23eb3a0c/resolv.conf",
      "HostnamePath": "/var/lib/docker/containers/cbc319cfc931ace3883b6f80a76a4fb3296e2d52f2d3349df2a64dae23eb3a0c/hostname",
      "HostsPath": "/var/lib/docker/containers/cbc319cfc931ace3883b6f80a76a4fb3296e2d52f2d3349df2a64dae23eb3a0c/hosts",
      "LogPath": "/var/lib/docker/containers/cbc319cfc931ace3883b6f80a76a4fb3296e2d52f2d3349df2a64dae23eb3a0c/cbc319cfc931ace3883b6f80a76a4fb3296e2d52f2d3349df2a64dae23eb3a0c-json.log",
      "Name": "/container_aa_delete_before_test",
      "RestartCount": 0,
      "Driver": "overlay2",
      "Platform": "linux",
      "MountLabel": "",
      "ProcessLabel": "",
      "AppArmorProfile": "",
      "ExecIDs": null,
      "HostConfig": {
        "Binds": null,
        "ContainerIDFile": "",
        "LogConfig": {
          "Type": "json-file",
          "Config": {
            
          }
        },
        "NetworkMode": "default",
        "PortBindings": null,
        "RestartPolicy": {
          "Name": "on-failure",
          "MaximumRetryCount": 0
        },
        "AutoRemove": false,
        "VolumeDriver": "",
        "VolumesFrom": null,
        "CapAdd": null,
        "CapDrop": null,
        "Capabilities": null,
        "Dns": [
          
        ],
        "DnsOptions": [
          
        ],
        "DnsSearch": [
          
        ],
        "ExtraHosts": null,
        "GroupAdd": null,
        "IpcMode": "private",
        "Cgroup": "",
        "Links": null,
        "OomScoreAdj": 0,
        "PidMode": "",
        "Privileged": false,
        "PublishAllPorts": false,
        "ReadonlyRootfs": false,
        "SecurityOpt": null,
        "UTSMode": "",
        "UsernsMode": "",
        "ShmSize": 67108864,
        "Runtime": "runc",
        "ConsoleSize": [
          0,
          0
        ],
        "Isolation": "",
        "CpuShares": 0,
        "Memory": 0,
        "NanoCpus": 0,
        "CgroupParent": "",
        "BlkioWeight": 0,
        "BlkioWeightDevice": null,
        "BlkioDeviceReadBps": null,
        "BlkioDeviceWriteBps": null,
        "BlkioDeviceReadIOps": null,
        "BlkioDeviceWriteIOps": null,
        "CpuPeriod": 0,
        "CpuQuota": 0,
        "CpuRealtimePeriod": 0,
        "CpuRealtimeRuntime": 0,
        "CpusetCpus": "",
        "CpusetMems": "",
        "Devices": null,
        "DeviceCgroupRules": null,
        "DeviceRequests": null,
        "KernelMemory": 0,
        "KernelMemoryTCP": 0,
        "MemoryReservation": 0,
        "MemorySwap": 0,
        "MemorySwappiness": null,
        "OomKillDisable": false,
        "PidsLimit": null,
        "Ulimits": null,
        "CpuCount": 0,
        "CpuPercent": 0,
        "IOMaximumIOps": 0,
        "IOMaximumBandwidth": 0,
        "Mounts": [
          {
            "Type": "bind",
            "Source": "/host_mnt/c/Users/helmut.kemper.CS/GolandProjects/apagar/config.conf",
            "Target": "/etc/mongo.conf"
          }
        ],
        "MaskedPaths": [
          "/proc/asound",
          "/proc/acpi",
          "/proc/kcore",
          "/proc/keys",
          "/proc/latency_stats",
          "/proc/timer_list",
          "/proc/timer_stats",
          "/proc/sched_debug",
          "/proc/scsi",
          "/sys/firmware"
        ],
        "ReadonlyPaths": [
          "/proc/bus",
          "/proc/fs",
          "/proc/irq",
          "/proc/sys",
          "/proc/sysrq-trigger"
        ]
      },
      "GraphDriver": {
        "Data": {
          "LowerDir": "/var/lib/docker/overlay2/51c50db93cb81ee2aff9931fd22c0b548b2f57f1c878ff2325cb6d6e26c7a4e1-init/diff:/var/lib/docker/overlay2/37a90fbefeebbfee077e4ca7cfbd2085dfacbea3e44dc8f0db44947c66bf090d/diff:/var/lib/docker/overlay2/412cf94e7f60c552e124d6b141c6a4e96363b4de8a167d89d2a863e7fa4b2252/diff:/var/lib/docker/overlay2/719fec79dedd6add3d5a68bad91c218d497bb039a657539dafb48e6112a60f64/diff:/var/lib/docker/overlay2/2ea298a3d6e01ca4915e5725168d6de8cc36be9d76d61235990776dbf5c297c7/diff:/var/lib/docker/overlay2/072403bd7c974df22e203710d8d32f40c3ee6d272b0b5289c552aeae6d6d7b45/diff:/var/lib/docker/overlay2/a3fa03a0b26d7332fceb569d9b1bbd9cf4f00c810e9c4750a57bcc4ffeccf062/diff:/var/lib/docker/overlay2/60190b45c61567d76d3601922ee05e658db287a82e9d1c3045675d18cb16cba8/diff:/var/lib/docker/overlay2/bac87aadfa1f68b8c228b41ae270ad2ab4bb98521a60a7ccaf8b8c1f0000bb58/diff:/var/lib/docker/overlay2/f97033a63fe4ba4535dd8b19d0b85a9390214fbacc03d91b7b276a77a2ac086f/diff:/var/lib/docker/overlay2/b2720a9b34cc3f0ff201a5935e8bd10a752e5a06f6ccae7e915c8a7663674417/diff:/var/lib/docker/overlay2/7152b4e53bfe00bcaf249338206a873c6a64e6bad385bbc0fd4c6359a56420e9/diff:/var/lib/docker/overlay2/f4d3ccc17cb609724dc83aae85699b2614ac5811140064b6b80efc40f0f79ef6/diff:/var/lib/docker/overlay2/a84680b62c1be0b9c2c2f6fadad18e1cd4355edf4091b6bd1555b212fb7d9af3/diff",
          "MergedDir": "/var/lib/docker/overlay2/51c50db93cb81ee2aff9931fd22c0b548b2f57f1c878ff2325cb6d6e26c7a4e1/merged",
          "UpperDir": "/var/lib/docker/overlay2/51c50db93cb81ee2aff9931fd22c0b548b2f57f1c878ff2325cb6d6e26c7a4e1/diff",
          "WorkDir": "/var/lib/docker/overlay2/51c50db93cb81ee2aff9931fd22c0b548b2f57f1c878ff2325cb6d6e26c7a4e1/work"
        },
        "Name": "overlay2"
      },
      "Mounts": [
        {
          "Type": "volume",
          "Name": "0effbc14e66c009851e98edae65fa520d850143faf11b4594fda403aec9111df",
          "Source": "/var/lib/docker/volumes/0effbc14e66c009851e98edae65fa520d850143faf11b4594fda403aec9111df/_data",
          "Destination": "/data/configdb",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "volume",
          "Name": "4534fab9513dcfd31db263ffb0db1372bd3ebcc72d65a86683b09a237fc10173",
          "Source": "/var/lib/docker/volumes/4534fab9513dcfd31db263ffb0db1372bd3ebcc72d65a86683b09a237fc10173/_data",
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
      ],
      "Config": {
        "Hostname": "cbc319cfc931",
        "Domainname": "",
        "User": "",
        "AttachStdin": false,
        "AttachStdout": false,
        "AttachStderr": false,
        "ExposedPorts": {
          "27017/tcp": {
            
          }
        },
        "Tty": false,
        "OpenStdin": false,
        "StdinOnce": false,
        "Env": [
          "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
          "GOSU_VERSION=1.12",
          "JSYAML_VERSION=3.13.1",
          "GPG_KEYS=2930ADAE8CAF5059EE73BB4B58712A2291FA4AD5",
          "MONGO_PACKAGE=mongodb-org",
          "MONGO_REPO=repo.mongodb.org",
          "MONGO_MAJOR=3.6",
          "MONGO_VERSION=3.6.18"
        ],
        "Cmd": [
          "mongod"
        ],
        "Image": "mongo:3",
        "Volumes": {
          "/data/configdb": {
            
          },
          "/data/db": {
            
          }
        },
        "WorkingDir": "",
        "Entrypoint": [
          "docker-entrypoint.sh"
        ],
        "OnBuild": null,
        "Labels": {
          
        }
      },
      "NetworkSettings": {
        "Bridge": "",
        "SandboxID": "73738ac49cb3555f30a8192557cb327e376d6565f562af2d5e3f5890420c0c1d",
        "HairpinMode": false,
        "LinkLocalIPv6Address": "",
        "LinkLocalIPv6PrefixLen": 0,
        "Ports": {
          "27017/tcp": null
        },
        "SandboxKey": "/var/run/docker/netns/73738ac49cb3",
        "SecondaryIPAddresses": null,
        "SecondaryIPv6Addresses": null,
        "EndpointID": "",
        "Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "IPAddress": "",
        "IPPrefixLen": 0,
        "IPv6Gateway": "",
        "MacAddress": "",
        "Networks": {
          "network_delete_before_test_b": {
            "IPAMConfig": {
              "IPv4Address": "11.0.0.2"
            },
            "Links": null,
            "Aliases": [
              "cbc319cfc931"
            ],
            "NetworkID": "86e8d30aa686ceddc1ec813674f355c49ad35dd848f05e29d8eb40333b1883ff",
            "EndpointID": "528a88347c25c4186d46e8828749e90b5aed53b90df792c97b79eaf78374870b",
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
      }
    },
    {
      "Id": "9741d024136301a9314055efdf3632c4b925036cf359f5a9e754a310ef5fb5e0",
      "Created": "2020-07-10T19:34:56.884081528Z",
      "Path": "./main",
      "Args": [
        
      ],
      "State": {
        "Status": "running",
        "Running": true,
        "Paused": false,
        "Restarting": false,
        "OOMKilled": false,
        "Dead": false,
        "Pid": 2142,
        "ExitCode": 0,
        "Error": "",
        "StartedAt": "2020-07-11T20:54:15.542470309Z",
        "FinishedAt": "2020-07-11T20:54:14.182886709Z"
      },
      "Image": "sha256:b914df731dd6ddedc44d3dcd1a733a33411c3617894f80046eff94dba821ae7a",
      "ResolvConfPath": "/var/lib/docker/containers/9741d024136301a9314055efdf3632c4b925036cf359f5a9e754a310ef5fb5e0/resolv.conf",
      "HostnamePath": "/var/lib/docker/containers/9741d024136301a9314055efdf3632c4b925036cf359f5a9e754a310ef5fb5e0/hostname",
      "HostsPath": "/var/lib/docker/containers/9741d024136301a9314055efdf3632c4b925036cf359f5a9e754a310ef5fb5e0/hosts",
      "LogPath": "/var/lib/docker/containers/9741d024136301a9314055efdf3632c4b925036cf359f5a9e754a310ef5fb5e0/9741d024136301a9314055efdf3632c4b925036cf359f5a9e754a310ef5fb5e0-json.log",
      "Name": "/delete_____a",
      "RestartCount": 0,
      "Driver": "overlay2",
      "Platform": "linux",
      "MountLabel": "",
      "ProcessLabel": "",
      "AppArmorProfile": "",
      "ExecIDs": null,
      "HostConfig": {
        "Binds": null,
        "ContainerIDFile": "",
        "LogConfig": {
          "Type": "json-file",
          "Config": {
            
          }
        },
        "NetworkMode": "default",
        "PortBindings": {
          "3000/tcp": [
            {
              "HostIp": "",
              "HostPort": "3000/tcp"
            }
          ]
        },
        "RestartPolicy": {
          "Name": "unless-stopped",
          "MaximumRetryCount": 0
        },
        "AutoRemove": false,
        "VolumeDriver": "",
        "VolumesFrom": null,
        "CapAdd": null,
        "CapDrop": null,
        "Capabilities": null,
        "Dns": [
          
        ],
        "DnsOptions": [
          
        ],
        "DnsSearch": [
          
        ],
        "ExtraHosts": null,
        "GroupAdd": null,
        "IpcMode": "private",
        "Cgroup": "",
        "Links": null,
        "OomScoreAdj": 0,
        "PidMode": "",
        "Privileged": false,
        "PublishAllPorts": false,
        "ReadonlyRootfs": false,
        "SecurityOpt": null,
        "UTSMode": "",
        "UsernsMode": "",
        "ShmSize": 67108864,
        "Runtime": "runc",
        "ConsoleSize": [
          0,
          0
        ],
        "Isolation": "",
        "CpuShares": 0,
        "Memory": 0,
        "NanoCpus": 0,
        "CgroupParent": "",
        "BlkioWeight": 0,
        "BlkioWeightDevice": null,
        "BlkioDeviceReadBps": null,
        "BlkioDeviceWriteBps": null,
        "BlkioDeviceReadIOps": null,
        "BlkioDeviceWriteIOps": null,
        "CpuPeriod": 0,
        "CpuQuota": 0,
        "CpuRealtimePeriod": 0,
        "CpuRealtimeRuntime": 0,
        "CpusetCpus": "",
        "CpusetMems": "",
        "Devices": null,
        "DeviceCgroupRules": null,
        "DeviceRequests": null,
        "KernelMemory": 0,
        "KernelMemoryTCP": 0,
        "MemoryReservation": 0,
        "MemorySwap": 0,
        "MemorySwappiness": null,
        "OomKillDisable": false,
        "PidsLimit": null,
        "Ulimits": null,
        "CpuCount": 0,
        "CpuPercent": 0,
        "IOMaximumIOps": 0,
        "IOMaximumBandwidth": 0,
        "Mounts": [
          {
            "Type": "bind",
            "Source": "/var/run/docker.sock",
            "Target": "/var/run/docker.sock"
          }
        ],
        "MaskedPaths": [
          "/proc/asound",
          "/proc/acpi",
          "/proc/kcore",
          "/proc/keys",
          "/proc/latency_stats",
          "/proc/timer_list",
          "/proc/timer_stats",
          "/proc/sched_debug",
          "/proc/scsi",
          "/sys/firmware"
        ],
        "ReadonlyPaths": [
          "/proc/bus",
          "/proc/fs",
          "/proc/irq",
          "/proc/sys",
          "/proc/sysrq-trigger"
        ]
      },
      "GraphDriver": {
        "Data": {
          "LowerDir": "/var/lib/docker/overlay2/769bc42f5ecea897e39ca4e52f8d7bc7182b6b4fcb3869c1907115fd1e22c114-init/diff:/var/lib/docker/overlay2/42d05bd4bd434b1b7068e5091fa393b9178fa80ff00f696a949988139f78fc4e/diff",
          "MergedDir": "/var/lib/docker/overlay2/769bc42f5ecea897e39ca4e52f8d7bc7182b6b4fcb3869c1907115fd1e22c114/merged",
          "UpperDir": "/var/lib/docker/overlay2/769bc42f5ecea897e39ca4e52f8d7bc7182b6b4fcb3869c1907115fd1e22c114/diff",
          "WorkDir": "/var/lib/docker/overlay2/769bc42f5ecea897e39ca4e52f8d7bc7182b6b4fcb3869c1907115fd1e22c114/work"
        },
        "Name": "overlay2"
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
      ],
      "Config": {
        "Hostname": "9741d0241363",
        "Domainname": "",
        "User": "",
        "AttachStdin": false,
        "AttachStdout": false,
        "AttachStderr": false,
        "ExposedPorts": {
          "3000/tcp": {
            
          }
        },
        "Tty": false,
        "OpenStdin": false,
        "StdinOnce": false,
        "Env": [
          "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
        ],
        "Cmd": [
          "./main"
        ],
        "Image": "delete:latest",
        "Volumes": {
          "/var/run/docker.sock": {
            
          }
        },
        "WorkingDir": "",
        "Entrypoint": null,
        "OnBuild": null,
        "Labels": {
          
        }
      },
      "NetworkSettings": {
        "Bridge": "",
        "SandboxID": "b6bf4de384f963d7ed5e7b2acdad2f5eceec553780d3d4917edb775abbcf0684",
        "HairpinMode": false,
        "LinkLocalIPv6Address": "",
        "LinkLocalIPv6PrefixLen": 0,
        "Ports": {
          "3000/tcp": [
            {
              "HostIp": "0.0.0.0",
              "HostPort": "3000"
            }
          ]
        },
        "SandboxKey": "/var/run/docker/netns/b6bf4de384f9",
        "SecondaryIPAddresses": null,
        "SecondaryIPv6Addresses": null,
        "EndpointID": "e3504fa600a7bdb144350dc77c4d5e6b502c848c628bebab9bf29aa7a562218c",
        "Gateway": "172.17.0.1",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "IPAddress": "172.17.0.2",
        "IPPrefixLen": 16,
        "IPv6Gateway": "",
        "MacAddress": "02:42:ac:11:00:02",
        "Networks": {
          "bridge": {
            "IPAMConfig": null,
            "Links": null,
            "Aliases": null,
            "NetworkID": "d5c45404eed452ab6908b06f0f492cf05ef16a97cf105eb1d581452b59859470",
            "EndpointID": "e3504fa600a7bdb144350dc77c4d5e6b502c848c628bebab9bf29aa7a562218c",
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
      }
    },
    {
      "Id": "6162ec603638ebb1edb50cc5bfdf7292fb232c5e384e0841eb5d0080f339be68",
      "Created": "2020-07-10T13:43:00.390992916Z",
      "Path": "docker-entrypoint.sh",
      "Args": [
        "mongod"
      ],
      "State": {
        "Status": "running",
        "Running": true,
        "Paused": false,
        "Restarting": false,
        "OOMKilled": false,
        "Dead": false,
        "Pid": 2171,
        "ExitCode": 0,
        "Error": "",
        "StartedAt": "2020-07-11T20:54:15.636334409Z",
        "FinishedAt": "2020-07-11T20:54:14.180594309Z"
      },
      "Image": "sha256:6d11486a97a77beaad31f63463a744dc3070ed4070bd15a695898a171f349441",
      "ResolvConfPath": "/var/lib/docker/containers/6162ec603638ebb1edb50cc5bfdf7292fb232c5e384e0841eb5d0080f339be68/resolv.conf",
      "HostnamePath": "/var/lib/docker/containers/6162ec603638ebb1edb50cc5bfdf7292fb232c5e384e0841eb5d0080f339be68/hostname",
      "HostsPath": "/var/lib/docker/containers/6162ec603638ebb1edb50cc5bfdf7292fb232c5e384e0841eb5d0080f339be68/hosts",
      "LogPath": "/var/lib/docker/containers/6162ec603638ebb1edb50cc5bfdf7292fb232c5e384e0841eb5d0080f339be68/6162ec603638ebb1edb50cc5bfdf7292fb232c5e384e0841eb5d0080f339be68-json.log",
      "Name": "/container_b_delete_before_test",
      "RestartCount": 0,
      "Driver": "overlay2",
      "Platform": "linux",
      "MountLabel": "",
      "ProcessLabel": "",
      "AppArmorProfile": "",
      "ExecIDs": null,
      "HostConfig": {
        "Binds": null,
        "ContainerIDFile": "",
        "LogConfig": {
          "Type": "json-file",
          "Config": {
            
          }
        },
        "NetworkMode": "default",
        "PortBindings": {
          "27017/tcp": [
            {
              "HostIp": "",
              "HostPort": "27017/tcp"
            }
          ]
        },
        "RestartPolicy": {
          "Name": "unless-stopped",
          "MaximumRetryCount": 0
        },
        "AutoRemove": false,
        "VolumeDriver": "",
        "VolumesFrom": null,
        "CapAdd": null,
        "CapDrop": null,
        "Capabilities": null,
        "Dns": [
          
        ],
        "DnsOptions": [
          
        ],
        "DnsSearch": [
          
        ],
        "ExtraHosts": null,
        "GroupAdd": null,
        "IpcMode": "private",
        "Cgroup": "",
        "Links": null,
        "OomScoreAdj": 0,
        "PidMode": "",
        "Privileged": false,
        "PublishAllPorts": false,
        "ReadonlyRootfs": false,
        "SecurityOpt": null,
        "UTSMode": "",
        "UsernsMode": "",
        "ShmSize": 67108864,
        "Runtime": "runc",
        "ConsoleSize": [
          0,
          0
        ],
        "Isolation": "",
        "CpuShares": 0,
        "Memory": 0,
        "NanoCpus": 0,
        "CgroupParent": "",
        "BlkioWeight": 0,
        "BlkioWeightDevice": null,
        "BlkioDeviceReadBps": null,
        "BlkioDeviceWriteBps": null,
        "BlkioDeviceReadIOps": null,
        "BlkioDeviceWriteIOps": null,
        "CpuPeriod": 0,
        "CpuQuota": 0,
        "CpuRealtimePeriod": 0,
        "CpuRealtimeRuntime": 0,
        "CpusetCpus": "",
        "CpusetMems": "",
        "Devices": null,
        "DeviceCgroupRules": null,
        "DeviceRequests": null,
        "KernelMemory": 0,
        "KernelMemoryTCP": 0,
        "MemoryReservation": 0,
        "MemorySwap": 0,
        "MemorySwappiness": null,
        "OomKillDisable": false,
        "PidsLimit": null,
        "Ulimits": null,
        "CpuCount": 0,
        "CpuPercent": 0,
        "IOMaximumIOps": 0,
        "IOMaximumBandwidth": 0,
        "Mounts": [
          {
            "Type": "bind",
            "Source": "/host_mnt/c/Users/helmut.kemper.CS/GolandProjects/apagar/config.conf",
            "Target": "/etc/mongo.conf"
          }
        ],
        "MaskedPaths": [
          "/proc/asound",
          "/proc/acpi",
          "/proc/kcore",
          "/proc/keys",
          "/proc/latency_stats",
          "/proc/timer_list",
          "/proc/timer_stats",
          "/proc/sched_debug",
          "/proc/scsi",
          "/sys/firmware"
        ],
        "ReadonlyPaths": [
          "/proc/bus",
          "/proc/fs",
          "/proc/irq",
          "/proc/sys",
          "/proc/sysrq-trigger"
        ]
      },
      "GraphDriver": {
        "Data": {
          "LowerDir": "/var/lib/docker/overlay2/3bd2b8ecbfd3e20eb1a35c7152ed56b5771a35643942887b7f0d24d8fa8004df-init/diff:/var/lib/docker/overlay2/14a45fd997ee201b9c8aef3947ba235e7d12a3181fa51cb58f3811bdabb7fd06/diff:/var/lib/docker/overlay2/065ba9891164f4680de8339396f67ceff179093201ccac628aa06fa761edda9b/diff:/var/lib/docker/overlay2/18f5d464c0e9c215dadd53154c9272941da0e2a80c52ad7a75c4d59182e9cc08/diff:/var/lib/docker/overlay2/cf83fa051ede97f58d8172ebaecc782104b07404bd6318756f9af21997587865/diff:/var/lib/docker/overlay2/1bee6683a4a529f55d4cbfb73ecc5563724e13e9f2d241e69a92739ee808dada/diff:/var/lib/docker/overlay2/5fd3492329d922a511c047593f973651d23a13fbfe9274ca6093d26c0d2ff77f/diff:/var/lib/docker/overlay2/d250a5106f71f3d3a5135663e9e1e9ceebd56f69c43474cb7bf0538050f9722a/diff:/var/lib/docker/overlay2/92991743124e5440bca9b720762dbde28f734737afd97558873e8614664d51c1/diff:/var/lib/docker/overlay2/2fe7d8c7708c8a2ba2bcd5bd110622fff1be99862799e11c19b9b29dd7dec43f/diff:/var/lib/docker/overlay2/aa68a4ad48e8575744902cd56c2b235af14778d407a78a5852b370865e0d4d06/diff:/var/lib/docker/overlay2/5f2fe41a64318eaea4976f51bcd52f9ac4ea013d5980af85d8d8563934535d74/diff:/var/lib/docker/overlay2/a6c093de29c074a5872439162bdb67bf0927db1ed3e6bbf5a5872e7984b46cc4/diff:/var/lib/docker/overlay2/a133f40011ec644f15bd8f09b2b88f9551552f8ffb2fa744d4c8a3d9f86aed6a/diff",
          "MergedDir": "/var/lib/docker/overlay2/3bd2b8ecbfd3e20eb1a35c7152ed56b5771a35643942887b7f0d24d8fa8004df/merged",
          "UpperDir": "/var/lib/docker/overlay2/3bd2b8ecbfd3e20eb1a35c7152ed56b5771a35643942887b7f0d24d8fa8004df/diff",
          "WorkDir": "/var/lib/docker/overlay2/3bd2b8ecbfd3e20eb1a35c7152ed56b5771a35643942887b7f0d24d8fa8004df/work"
        },
        "Name": "overlay2"
      },
      "Mounts": [
        {
          "Type": "volume",
          "Name": "82ce0b093414575ec8b295c5143c61a16844db85b0379b69c28389a5afc438da",
          "Source": "/var/lib/docker/volumes/82ce0b093414575ec8b295c5143c61a16844db85b0379b69c28389a5afc438da/_data",
          "Destination": "/data/configdb",
          "Driver": "local",
          "Mode": "",
          "RW": true,
          "Propagation": ""
        },
        {
          "Type": "volume",
          "Name": "1b9a4f14103a20b8bc902a070f67fdafc483a0bb60b3d9bc67c29f5aff0de1e6",
          "Source": "/var/lib/docker/volumes/1b9a4f14103a20b8bc902a070f67fdafc483a0bb60b3d9bc67c29f5aff0de1e6/_data",
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
      ],
      "Config": {
        "Hostname": "6162ec603638",
        "Domainname": "",
        "User": "",
        "AttachStdin": false,
        "AttachStdout": false,
        "AttachStderr": false,
        "ExposedPorts": {
          "27017/tcp": {
            
          }
        },
        "Tty": false,
        "OpenStdin": false,
        "StdinOnce": false,
        "Env": [
          "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
          "GOSU_VERSION=1.12",
          "JSYAML_VERSION=3.13.1",
          "GPG_KEYS=E162F504A20CDF15827F718D4B7C549A058F8B6B",
          "MONGO_PACKAGE=mongodb-org",
          "MONGO_REPO=repo.mongodb.org",
          "MONGO_MAJOR=4.2",
          "MONGO_VERSION=4.2.8"
        ],
        "Cmd": [
          "mongod"
        ],
        "Image": "mongo:latest",
        "Volumes": {
          "/data/configdb": {
            
          },
          "/data/db": {
            
          }
        },
        "WorkingDir": "",
        "Entrypoint": [
          "docker-entrypoint.sh"
        ],
        "OnBuild": null,
        "Labels": {
          
        }
      },
      "NetworkSettings": {
        "Bridge": "",
        "SandboxID": "4d16b7ac8027acd802315e346bc19c9a769b55f6702220666698c9ca0bb10b24",
        "HairpinMode": false,
        "LinkLocalIPv6Address": "",
        "LinkLocalIPv6PrefixLen": 0,
        "Ports": {
          "27017/tcp": [
            {
              "HostIp": "0.0.0.0",
              "HostPort": "27017"
            }
          ]
        },
        "SandboxKey": "/var/run/docker/netns/4d16b7ac8027",
        "SecondaryIPAddresses": null,
        "SecondaryIPv6Addresses": null,
        "EndpointID": "b8f86ea0eb732f9b1cafc4318f8bf79390563057b46308bd9c2bd3fafae2dc5e",
        "Gateway": "172.17.0.1",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "IPAddress": "172.17.0.3",
        "IPPrefixLen": 16,
        "IPv6Gateway": "",
        "MacAddress": "02:42:ac:11:00:03",
        "Networks": {
          "bridge": {
            "IPAMConfig": null,
            "Links": null,
            "Aliases": null,
            "NetworkID": "d5c45404eed452ab6908b06f0f492cf05ef16a97cf105eb1d581452b59859470",
            "EndpointID": "b8f86ea0eb732f9b1cafc4318f8bf79390563057b46308bd9c2bd3fafae2dc5e",
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
      }
    }
  ]
}
```