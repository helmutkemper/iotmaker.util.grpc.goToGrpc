# Container Inspect

```url
http://localhost:8081/containerInspectById?ID=63b43a9b85ef6761d3d9669c331e22bfe13907c3d40fd0c740688192c774ed58
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
        "Pid": 2244,
        "ExitCode": 0,
        "Error": "",
        "StartedAt": "2020-07-11T16:12:34.941989354Z",
        "FinishedAt": "2020-07-11T16:12:33.325981854Z"
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
        "SandboxID": "7777f23a0f72592c5b1bbde7272ff219d5ada0bdfe636360f73a033193a2c1ab",
        "HairpinMode": false,
        "LinkLocalIPv6Address": "",
        "LinkLocalIPv6PrefixLen": 0,
        "Ports": {
          "27017/tcp": null
        },
        "SandboxKey": "/var/run/docker/netns/7777f23a0f72",
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
      }
    }
  ]
}
```