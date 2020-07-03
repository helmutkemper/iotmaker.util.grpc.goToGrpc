package server

import (
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportStatsToGRpc(stats types.Stats) (statistics *pb.Statistics) {
	var statsPidsStats = &pb.PidsStats{
		Current: stats.PidsStats.Current,
		Limit:   stats.PidsStats.Limit,
	}

	var statsBlkioStatsIoServiceBytesRecursive = make([]*pb.BlkioStatEntry, 0)
	for _, blkioStats := range stats.BlkioStats.IoServiceBytesRecursive {
		statsBlkioStatsIoServiceBytesRecursive = append(statsBlkioStatsIoServiceBytesRecursive, &pb.BlkioStatEntry{
			Major: blkioStats.Major,
			Minor: blkioStats.Minor,
			Op:    blkioStats.Op,
			Value: blkioStats.Value,
		})
	}

	var statsBlkioStatsIoServicedRecursive = make([]*pb.BlkioStatEntry, 0)
	for _, blkioStats := range stats.BlkioStats.IoServicedRecursive {
		statsBlkioStatsIoServicedRecursive = append(statsBlkioStatsIoServicedRecursive, &pb.BlkioStatEntry{
			Major: blkioStats.Major,
			Minor: blkioStats.Minor,
			Op:    blkioStats.Op,
			Value: blkioStats.Value,
		})
	}

	var statsBlkioStatsIoQueuedRecursive = make([]*pb.BlkioStatEntry, 0)
	for _, blkioStats := range stats.BlkioStats.IoQueuedRecursive {
		statsBlkioStatsIoQueuedRecursive = append(statsBlkioStatsIoQueuedRecursive, &pb.BlkioStatEntry{
			Major: blkioStats.Major,
			Minor: blkioStats.Minor,
			Op:    blkioStats.Op,
			Value: blkioStats.Value,
		})
	}

	var statsBlkioStatsIoServiceTimeRecursive = make([]*pb.BlkioStatEntry, 0)
	for _, blkioStats := range stats.BlkioStats.IoServiceTimeRecursive {
		statsBlkioStatsIoServiceTimeRecursive = append(statsBlkioStatsIoServiceTimeRecursive, &pb.BlkioStatEntry{
			Major: blkioStats.Major,
			Minor: blkioStats.Minor,
			Op:    blkioStats.Op,
			Value: blkioStats.Value,
		})
	}

	var statsBlkioStatsIoWaitTimeRecursive = make([]*pb.BlkioStatEntry, 0)
	for _, blkioStats := range stats.BlkioStats.IoWaitTimeRecursive {
		statsBlkioStatsIoWaitTimeRecursive = append(statsBlkioStatsIoWaitTimeRecursive, &pb.BlkioStatEntry{
			Major: blkioStats.Major,
			Minor: blkioStats.Minor,
			Op:    blkioStats.Op,
			Value: blkioStats.Value,
		})
	}

	var statsBlkioStatsIoMergedRecursive = make([]*pb.BlkioStatEntry, 0)
	for _, blkioStats := range stats.BlkioStats.IoMergedRecursive {
		statsBlkioStatsIoMergedRecursive = append(statsBlkioStatsIoMergedRecursive, &pb.BlkioStatEntry{
			Major: blkioStats.Major,
			Minor: blkioStats.Minor,
			Op:    blkioStats.Op,
			Value: blkioStats.Value,
		})
	}

	var statsBlkioStatsIoTimeRecursive = make([]*pb.BlkioStatEntry, 0)
	for _, blkioStats := range stats.BlkioStats.IoTimeRecursive {
		statsBlkioStatsIoTimeRecursive = append(statsBlkioStatsIoTimeRecursive, &pb.BlkioStatEntry{
			Major: blkioStats.Major,
			Minor: blkioStats.Minor,
			Op:    blkioStats.Op,
			Value: blkioStats.Value,
		})
	}

	var statsBlkioStatsSectorsRecursive = make([]*pb.BlkioStatEntry, 0)
	for _, blkioStats := range stats.BlkioStats.SectorsRecursive {
		statsBlkioStatsIoTimeRecursive = append(statsBlkioStatsSectorsRecursive, &pb.BlkioStatEntry{
			Major: blkioStats.Major,
			Minor: blkioStats.Minor,
			Op:    blkioStats.Op,
			Value: blkioStats.Value,
		})
	}

	var statsBlkioStats = &pb.BlkioStats{
		IoServiceBytesRecursive: statsBlkioStatsIoServiceBytesRecursive,
		IoServicedRecursive:     statsBlkioStatsIoServicedRecursive,
		IoQueuedRecursive:       statsBlkioStatsIoQueuedRecursive,
		IoServiceTimeRecursive:  statsBlkioStatsIoServiceTimeRecursive,
		IoWaitTimeRecursive:     statsBlkioStatsIoWaitTimeRecursive,
		IoMergedRecursive:       statsBlkioStatsIoMergedRecursive,
		IoTimeRecursive:         statsBlkioStatsIoTimeRecursive,
		SectorsRecursive:        statsBlkioStatsSectorsRecursive,
	}

	var statsStorageStats = &pb.StorageStats{
		ReadCountNormalized:  stats.StorageStats.ReadCountNormalized,
		ReadSizeBytes:        stats.StorageStats.ReadSizeBytes,
		WriteCountNormalized: stats.StorageStats.WriteCountNormalized,
		WriteSizeBytes:       stats.StorageStats.WriteSizeBytes,
	}

	var statsCPUStatsCPUUsage = &pb.CPUUsage{
		TotalUsage:        stats.CPUStats.CPUUsage.TotalUsage,
		PercpuUsage:       stats.CPUStats.CPUUsage.PercpuUsage,
		UsageInKernelmode: stats.CPUStats.CPUUsage.UsageInKernelmode,
		UsageInUsermode:   stats.CPUStats.CPUUsage.UsageInUsermode,
	}

	var statsCPUStatsThrottlingData = &pb.ThrottlingData{
		Periods:          stats.CPUStats.ThrottlingData.Periods,
		ThrottledPeriods: stats.CPUStats.ThrottlingData.ThrottledPeriods,
		ThrottledTime:    stats.CPUStats.ThrottlingData.ThrottledTime,
	}

	var statsCPUStats = &pb.CPUStats{
		CPUUsage: statsCPUStatsCPUUsage,

		SystemUsage: stats.CPUStats.SystemUsage,
		OnlineCPUs:  stats.CPUStats.OnlineCPUs,

		ThrottlingData: statsCPUStatsThrottlingData,
	}

	var statsPreCPUStatsCPUUsage = &pb.CPUUsage{
		TotalUsage:        stats.PreCPUStats.CPUUsage.TotalUsage,
		PercpuUsage:       stats.PreCPUStats.CPUUsage.PercpuUsage,
		UsageInKernelmode: stats.PreCPUStats.CPUUsage.UsageInKernelmode,
		UsageInUsermode:   stats.PreCPUStats.CPUUsage.UsageInUsermode,
	}

	var statsPreCPUStatsThrottlingData = &pb.ThrottlingData{
		Periods:          stats.PreCPUStats.ThrottlingData.Periods,
		ThrottledPeriods: stats.PreCPUStats.ThrottlingData.ThrottledPeriods,
		ThrottledTime:    stats.PreCPUStats.ThrottlingData.ThrottledTime,
	}

	var statsPreCPUStats = &pb.CPUStats{
		CPUUsage: statsPreCPUStatsCPUUsage,

		SystemUsage: stats.PreCPUStats.SystemUsage,
		OnlineCPUs:  stats.PreCPUStats.OnlineCPUs,

		ThrottlingData: statsPreCPUStatsThrottlingData,
	}

	var statsMemoryStats = &pb.MemoryStats{
		Usage:             stats.MemoryStats.Usage,
		MaxUsage:          stats.MemoryStats.MaxUsage,
		Stats:             stats.MemoryStats.Stats,
		Failcnt:           stats.MemoryStats.Failcnt,
		Limit:             stats.MemoryStats.Limit,
		Commit:            stats.MemoryStats.Commit,
		CommitPeak:        stats.MemoryStats.CommitPeak,
		PrivateWorkingSet: stats.MemoryStats.PrivateWorkingSet,
	}

	statistics = &pb.Statistics{
		Read:    stats.Read.Unix(),
		PreRead: stats.PreRead.Unix(),

		PidsStats:  statsPidsStats,
		BlkioStats: statsBlkioStats,

		NumProcs: stats.NumProcs,

		StorageStats: statsStorageStats,
		CPUStats:     statsCPUStats,
		PreCPUStats:  statsPreCPUStats,
		MemoryStats:  statsMemoryStats,
	}

	return
}
