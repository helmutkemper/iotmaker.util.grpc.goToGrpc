package server

import (
	"github.com/docker/docker/api/types/mount"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportGRpcToArrayMount(
	pt []*pb.Mount,
) (
	mountList []mount.Mount,
) {

	mountList = make([]mount.Mount, 0)
	for _, m := range pt {
		mountList = append(mountList, mount.Mount{
			Type:        m.Type,
			Source:      m.Source,
			Target:      m.Target,
			ReadOnly:    m.ReadOnly,
			Consistency: mount.Consistency(m.Consistency),
		})
	}

	return
}
