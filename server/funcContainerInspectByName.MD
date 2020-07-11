# Container Inspect By Name

```url
http://localhost:8081/containerInspectByName?Name=delete_____a
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
        "Pid": 2140,
        "ExitCode": 0,
        "Error": "",
        "StartedAt": "2020-07-11T16:12:34.733074454Z",
        "FinishedAt": "2020-07-11T16:12:33.330706654Z"
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
        "SandboxID": "bf0bb168f2e7f2d2c46d6d3c03ebed757c6b3bbb59b98000e092b4f5ffa7d17a",
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
        "SandboxKey": "/var/run/docker/netns/bf0bb168f2e7",
        "SecondaryIPAddresses": null,
        "SecondaryIPv6Addresses": null,
        "EndpointID": "0707480e98908edbcc240f69283e3ce25cababfcf215b159a8eeba4030b7d658",
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
      }
    }
  ]
}
```