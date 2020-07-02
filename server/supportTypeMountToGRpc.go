package server

import (
	"github.com/docker/docker/api/types"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportTypeMountToGRpc(mp []types.MountPoint) (mountPoint []*pb.MountPoint) {
	mountPoint = make([]*pb.MountPoint, 0)

	for _, point := range mp {
		mountPoint = append(mountPoint, &pb.MountPoint{
			Type:        string(point.Type),
			Name:        point.Name,
			Source:      point.Source,
			Destination: point.Destination,
			Driver:      point.Driver,
			Mode:        point.Mode,
			RW:          point.RW,
			Propagation: string(point.Propagation),
		})
	}

	return
}
