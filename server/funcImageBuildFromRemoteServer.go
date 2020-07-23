package server

import (
	"context"
	"encoding/json"
	"errors"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/util"
	"sync"
	"time"
)

type PullStatusList struct {
	m   map[string]BuildOrPullLog
	mux sync.Mutex
}

func (el *PullStatusList) Verify(key string) (found bool) {
	el.mux.Lock()
	defer el.mux.Unlock()

	if len(el.m) == 0 {
		return
	}

	_, found = el.m[key]
	return
}

func (el *PullStatusList) Set(key string, value BuildOrPullLog) {
	el.mux.Lock()
	defer el.mux.Unlock()

	if len(el.m) == 0 {
		el.m = make(map[string]BuildOrPullLog)
	}

	el.m[key] = value
}

func (el *PullStatusList) Get(key string) (value BuildOrPullLog, found bool) {
	el.mux.Lock()
	defer el.mux.Unlock()

	value, found = el.m[key]
	return
}

func (el *PullStatusList) TickerDeleteOldData() {
	el.mux.Lock()
	defer el.mux.Unlock()

	for k := range el.m {
		start := el.m[k].Start
		if time.Since(start) >= 2*time.Second*60*60 {
			delete(el.m, k)
		}
	}
}

var pullStatusList PullStatusList
var pullStatusTicker = time.NewTicker(30 * time.Second * 60)

type JSonImageBuildFromRemoteServer struct {
	ImageName  string
	ImageTags  []string
	ServerPath string
}

type BuildOrPullLog struct {
	Status iotmakerDocker.ContainerPullStatusSendToChannel
	Log    string
	Start  time.Time
}

func (el *GRpcServer) ImageBuildFromRemoteServer(
	ctx context.Context,
	in *pb.ImageBuildFromRemoteServerRequest,
) (
	response *pb.ImageBuildFromRemoteServerReply,
	err error,
) {

	_ = ctx
	err = el.Init()
	if err != nil {
		return
	}

	var pullStatusChannel = make(chan iotmakerDocker.ContainerPullStatusSendToChannel, 1)

	var imageChannelID string

	for {
		imageChannelID = util.RandId30()
		if pullStatusList.Verify(imageChannelID) == false {
			break
		}
	}
	pullStatusList.Set(
		imageChannelID,
		BuildOrPullLog{
			Start: time.Now(),
		},
	)

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel, imageChannelID string) {

		for {
			select {
			case status := <-c:
				var tmp, _ = pullStatusList.Get(imageChannelID)
				tmp.Status = status
				tmp.Log += status.Stream

				pullStatusList.Set(imageChannelID, tmp)

				if status.Closed == true {
					return
				}
			}
		}

	}(pullStatusChannel, imageChannelID)

	var body = in.GetData()
	var inData JSonImageBuildFromRemoteServer
	err = json.Unmarshal(body, &inData)
	if err != nil {
		err = errors.New("json unmarshal error: " + err.Error())
		return
	}

	err = el.dockerSystem.ImageBuildFromRemoteServer(
		inData.ServerPath,
		inData.ImageName,
		inData.ImageTags,
		&pullStatusChannel,
	)
	if err != nil {
		return
	}

	response = &pb.ImageBuildFromRemoteServerReply{
		ID: imageChannelID,
	}

	return
}

func init() {
	go func() {
		for {
			select {
			case <-pullStatusTicker.C:
				pullStatusList.TickerDeleteOldData()
			}
		}
	}()
}
