package server

import (
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
)

func SupportBuildStatusToGRpc(st iotmakerDocker.ContainerPullStatusSendToChannel) (status *pb.ImageOrContainerBuildPullStatusReply) {
	status = &pb.ImageOrContainerBuildPullStatusReply{
		Waiting: int64(st.Waiting),
		Downloading: &pb.ImageOrContainerBuildPullSubStatus{
			Count:   int64(st.Downloading.Count),
			Current: int64(st.Downloading.Current),
			Total:   int64(st.Downloading.Total),
			Percent: st.Downloading.Percent,
		},
		VerifyingChecksum: int64(st.VerifyingChecksum),
		DownloadComplete:  int64(st.DownloadComplete),
		Extracting: &pb.ImageOrContainerBuildPullSubStatus{
			Count:   int64(st.Extracting.Count),
			Current: int64(st.Extracting.Current),
			Total:   int64(st.Extracting.Total),
			Percent: st.Extracting.Percent,
		},
		PullComplete:               int64(st.PullComplete),
		ImageName:                  st.ImageName,
		ImageID:                    st.ImageID,
		Closed:                     st.Closed,
		Stream:                     st.Stream,
		SuccessfullyBuildContainer: st.SuccessfullyBuildContainer,
		SuccessfullyBuildImage:     st.SuccessfullyBuildImage,
	}

	return
}
