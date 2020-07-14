# Container Statistics One Shot By Name

```url
http://localhost:8081/containerStatisticsOneShotByName
```

```json
{
    "Name": "/new_mondo_delete_before_test"
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
      "read": "2020-07-14T17:51:11.986138637Z",
      "preread": "2020-07-14T17:51:10.983001017Z",
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
          "total_usage": 45059976685,
          "percpu_usage": [
            24581131597,
            20478845088
          ],
          "usage_in_kernelmode": 14690000000,
          "usage_in_usermode": 29230000000
        },
        "system_cpu_usage": 29488280000000,
        "online_cpus": 2,
        "throttling_data": {
          "periods": 0,
          "throttled_periods": 0,
          "throttled_time": 0
        }
      },
      "precpu_stats": {
        "cpu_usage": {
          "total_usage": 45056801766,
          "percpu_usage": [
            24580553048,
            20476248718
          ],
          "usage_in_kernelmode": 14690000000,
          "usage_in_usermode": 29220000000
        },
        "system_cpu_usage": 29486290000000,
        "online_cpus": 2,
        "throttling_data": {
          "periods": 0,
          "throttled_periods": 0,
          "throttled_time": 0
        }
      },
      "memory_stats": {
        "usage": 72839168,
        "max_usage": 75022336,
        "stats": {
          "active_anon": 70021120,
          "active_file": 20480,
          "cache": 675840,
          "dirty": 0,
          "hierarchical_memory_limit": 9223372036854771712,
          "hierarchical_memsw_limit": 9223372036854771712,
          "inactive_anon": 0,
          "inactive_file": 782336,
          "mapped_file": 0,
          "pgfault": 21087,
          "pgmajfault": 0,
          "pgpgin": 23034,
          "pgpgout": 5762,
          "rss": 69812224,
          "rss_huge": 0,
          "total_active_anon": 70021120,
          "total_active_file": 20480,
          "total_cache": 675840,
          "total_dirty": 0,
          "total_inactive_anon": 0,
          "total_inactive_file": 782336,
          "total_mapped_file": 0,
          "total_pgfault": 21087,
          "total_pgmajfault": 0,
          "total_pgpgin": 23034,
          "total_pgpgout": 5762,
          "total_rss": 69812224,
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