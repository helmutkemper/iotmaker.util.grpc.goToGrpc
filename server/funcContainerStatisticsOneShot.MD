# container Statistics One Shot

```url
http://localhost:8081/containerStatisticsOneShot
```

```json
{
    "Id": "5150a300ecf00f69755c373696b1a7048b7312292e03cbeacbd74849df1cbe22"
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
      "read": "2020-07-14T14:35:59.850932022Z",
      "preread": "2020-07-14T14:35:58.847698948Z",
      "pids_stats": {
        "current": 32
      },
      "blkio_stats": {
        "io_service_bytes_recursive": [
          
        ],
        "io_serviced_recursive": [
          
        ],
        "io_queue_recursive": [
          
        ],
        "io_service_time_recursive": [
          
        ],
        "io_wait_time_recursive": [
          
        ],
        "io_merged_recursive": [
          
        ],
        "io_time_recursive": [
          
        ],
        "sectors_recursive": [
          
        ]
      },
      "num_procs": 0,
      "storage_stats": {
        
      },
      "cpu_stats": {
        "cpu_usage": {
          "total_usage": 12114670966,
          "percpu_usage": [
            7073661205,
            5041009761
          ],
          "usage_in_kernelmode": 3750000000,
          "usage_in_usermode": 7790000000
        },
        "system_cpu_usage": 16319390000000,
        "online_cpus": 2,
        "throttling_data": {
          "periods": 0,
          "throttled_periods": 0,
          "throttled_time": 0
        }
      },
      "precpu_stats": {
        "cpu_usage": {
          "total_usage": 12112073438,
          "percpu_usage": [
            7071636999,
            5040436439
          ],
          "usage_in_kernelmode": 3750000000,
          "usage_in_usermode": 7790000000
        },
        "system_cpu_usage": 16317390000000,
        "online_cpus": 2,
        "throttling_data": {
          "periods": 0,
          "throttled_periods": 0,
          "throttled_time": 0
        }
      },
      "memory_stats": {
        "usage": 73048064,
        "max_usage": 75022336,
        "stats": {
          "active_anon": 70578176,
          "active_file": 20480,
          "cache": 405504,
          "dirty": 0,
          "hierarchical_memory_limit": 9223372036854771712,
          "hierarchical_memsw_limit": 9223372036854771712,
          "inactive_anon": 0,
          "inactive_file": 442368,
          "mapped_file": 0,
          "pgfault": 20889,
          "pgmajfault": 0,
          "pgpgin": 20361,
          "pgpgout": 3000,
          "rss": 70512640,
          "rss_huge": 0,
          "total_active_anon": 70578176,
          "total_active_file": 20480,
          "total_cache": 405504,
          "total_dirty": 0,
          "total_inactive_anon": 0,
          "total_inactive_file": 442368,
          "total_mapped_file": 0,
          "total_pgfault": 20889,
          "total_pgmajfault": 0,
          "total_pgpgin": 20361,
          "total_pgpgout": 3000,
          "total_rss": 70512640,
          "total_rss_huge": 0,
          "total_unevictable": 0,
          "total_writeback": 0,
          "unevictable": 0,
          "writeback": 0
        },
        "limit": 2088132608
      }
    }
  ]
}
```