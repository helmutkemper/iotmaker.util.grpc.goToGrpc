package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/server"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	pb "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc/main/protobuf"
	"google.golang.org/grpc"
)

type Output struct {
	Length  int64
	Limit   int64
	Skip    int64
	Success bool
	Error   []string
	Data    interface{}
}

const (
	KHttpServerPort    = 8081
	KGRpcServerAddress = "localhost:50051"
)

const (
	KHeaderTypeHtmlUtf8 = "text/html; charset=utf-8"
)

var (
	GRpcClient pb.DockerServerClient
)

type containerCreateAndChangeExposedPort struct {
	ImageName        string
	ContainerName    string
	RestartPolicy    iotmakerDocker.RestartPolicy
	MountVolumes     []mount.Mount
	ContainerNetwork string
	CurrentPort      []nat.Port
	ChangeToPort     []nat.Port
}

func main() {
	var err error

	conn, err := grpc.Dial(KGRpcServerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	GRpcClient = pb.NewDockerServerClient(conn)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.HandleFunc("/", serveTemplate)
	mux.HandleFunc("/networkListAll", NetworkList)
	mux.HandleFunc("/imageListAll", ImageList)
	mux.HandleFunc("/containerCreate", ContainerCreate)
	mux.HandleFunc("/containerStatisticsOneShot", ContainerStatisticsOneShot)
	mux.HandleFunc("/containerStatisticsOneShotByName", ContainerStatisticsOneShotByName)
	mux.HandleFunc("/containerCreateAndStart", ContainerCreateAndStart)
	mux.HandleFunc("/containerCreateAndChangeExposedPort", ContainerCreateAndChangeExposedPort)
	mux.HandleFunc("/containerCreateChangeExposedPortAndStart", ContainerCreateChangeExposedPortAndStart)
	mux.HandleFunc("/containerCreateAndExposePortsAutomatically", ContainerCreateAndExposePortsAutomatically)
	mux.HandleFunc("/containerCreateExposePortsAutomaticallyAndStart", ContainerCreateExposePortsAutomaticallyAndStart)
	mux.HandleFunc("/containerCreateWithoutExposePorts", ContainerCreateWithoutExposePorts)
	mux.HandleFunc("/containerCreateWithoutExposePortsAndStart", ContainerCreateWithoutExposePortsAndStart)
	mux.HandleFunc("/containersListAll", ContainersList)
	mux.HandleFunc("/containerStopAndRemove", ContainerStopAndRemove)
	mux.HandleFunc("/containerRemove", ContainerRemove)
	mux.HandleFunc("/containerStop", ContainerStop)
	mux.HandleFunc("/containerStart", ContainerStart)
	mux.HandleFunc("/containerInspectById", ContainerInspect)
	mux.HandleFunc("/containerFindIdByName", ContainerFindIdByName)
	mux.HandleFunc("/containerFindIdByNameContains", ContainerFindIdByNameContains)
	mux.HandleFunc("/containerInspectByName", ContainerInspectByName)
	mux.HandleFunc("/containerInspectByNameContains", ContainerInspectByNameContains)
	mux.HandleFunc("/imageBuildFromRemoteServer", ImageBuildFromRemoteServer)
	mux.HandleFunc("/imageBuildAndContainerStartFromRemoteServer", ImageBuildAndContainerStartFromRemoteServer)
	mux.HandleFunc("/imageBuildFromRemoteServerStatus", ImageBuildFromRemoteServerStatus)
	//mux.HandleFunc("/imageFindIdByName", ImageFindIdByName)
	//mux.HandleFunc("/imageListExposedPorts", ImageListExposedPorts)
	//mux.HandleFunc("/imageListExposedPortsByName", ImageListExposedPortsByName)
	//mux.HandleFunc("/imageListExposedVolumes", ImageListExposedVolumes)
	//mux.HandleFunc("/imageListExposedVolumesByName", ImageListExposedVolumesByName)
	//mux.HandleFunc("/imageMountNatPortList", ImageMountNatPortList)
	//mux.HandleFunc("/imageMountNatPortListChangeExposed", ImageMountNatPortListChangeExposed)
	//mux.HandleFunc("/imageMountNatPortListChangeExposedWithIpAddress", ImageMountNatPortListChangeExposedWithIpAddress)
	//mux.HandleFunc("/imagePull", ImagePull)
	//mux.HandleFunc("/imageRemove", ImageRemove)
	//mux.HandleFunc("/imageRemoveByName", ImageRemoveByName)

	server := fmt.Sprintf(":%v", KHttpServerPort)
	fmt.Printf("Listening on %v...", server)
	err = http.ListenAndServe(server, mux)
	if err != nil {
		panic(err)
	}
}

func ToJson(dataType interface{}, data interface{}, w http.ResponseWriter, r *http.Request) {
	var skipString, limitString string
	var skip, limit int64
	var toOut interface{}
	var err error
	var errorList = make([]string, 0)
	var length int64
	var output []byte
	var success = true

	skipString = r.URL.Query().Get("skip")
	if skipString == "" {
		skipString = "0"
	}

	limitString = r.URL.Query().Get("limit")
	if limitString == "" {
		limitString = "0"
	}

	skip, err = strconv.ParseInt(skipString, 10, 64)
	if err != nil {
		skip = 0
		errorList = append(errorList, fmt.Sprintf("query string skip error: %v", err.Error()))
	}
	if skip < 0 {
		skip = 0
		errorList = append(errorList, fmt.Sprint("query string skip error: skip must be a positive value"))
	}

	limit, err = strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		limit = 0
		errorList = append(errorList, fmt.Sprintf("query string limit error: %v", err.Error()))
	}
	if limit < 0 {
		limit = 0
		errorList = append(errorList, fmt.Sprint("query string limit error: limit must be a positive value"))
	}

	switch dataType.(type) {
	case pb.Empty:
		toOut = make([]int, 0)
		length = 1

		if skip > 0 {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		}

	case pb.ContainerFindIdByNameReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ContainerFindIdByNameReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ContainerFindIdByNameContainsReply:
		var list []server.NameAndId
		err = json.Unmarshal(data.(*pb.ContainerFindIdByNameContainsReply).Data, &list)
		if err != nil {
			length = 0
			limit = 0
			skip = 0
			success = false
			errorList = append(errorList, err.Error())
			toOut = make([]int, 0)
		} else {
			toOut = list

			length = int64(len(toOut.([]server.NameAndId)))

			if skip >= length {
				toOut = make([]int, 0)
				length = 0
				errorList = append(errorList, "skip overflow")
				success = false
			} else {
				toOut = toOut.([]server.NameAndId)[skip:]
				length = int64(len(toOut.([]server.NameAndId)))
			}

			if length > 0 {
				if limit > length && limit > 0 {
					toOut = toOut.([]server.NameAndId)[:length]
					length = int64(len(toOut.([]server.NameAndId)))
				} else if limit > 0 {
					toOut = toOut.([]server.NameAndId)[:limit]
					length = int64(len(toOut.([]server.NameAndId)))
				} else {
					toOut = toOut.([]server.NameAndId)
					length = int64(len(toOut.([]server.NameAndId)))
				}
			}
		}

	case pb.ContainerCreateAndChangeExposedPortReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ContainerCreateAndChangeExposedPortReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ContainerCreateReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ContainerCreateReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ContainerCreateAndExposePortsAutomaticallyReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ContainerCreateAndExposePortsAutomaticallyReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ContainerCreateWithoutExposePortsReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ContainerCreateWithoutExposePortsReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ContainerCreateWithoutExposePortsAndStartReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ContainerCreateWithoutExposePortsAndStartReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ContainerCreateExposePortsAutomaticallyAndStartReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ContainerCreateExposePortsAutomaticallyAndStartReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ContainerCreateAndStartReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ContainerCreateAndStartReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ContainerCreateChangeExposedPortAndStartReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ContainerCreateChangeExposedPortAndStartReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ImageBuildFromRemoteServerReply:
		var list = []map[string]string{
			{
				"ID": data.(*pb.ImageBuildFromRemoteServerReply).ID,
			},
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]map[string]string)[skip:]
			length = int64(len(toOut.([]map[string]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]map[string]string)[:length]
				length = int64(len(toOut.([]map[string]string)))
			} else if limit > 0 {
				toOut = toOut.([]map[string]string)[:limit]
				length = int64(len(toOut.([]map[string]string)))
			} else {
				toOut = toOut.([]map[string]string)
				length = int64(len(toOut.([]map[string]string)))
			}
		}

	case pb.ImageOrContainerBuildPullStatusRequest:
		var list server.BuildOrPullLog
		err = json.Unmarshal(data.(*pb.ImageOrContainerBuildPullStatusReply).Data, &list)
		if err != nil {
			length = 0
			limit = 0
			skip = 0
			success = false
			errorList = append(errorList, err.Error())
			toOut = make([]int, 0)
		} else {
			toOut = []server.BuildOrPullLog{
				list,
			}
			length = 1

			if skip >= length {
				toOut = make([]int, 0)
				length = 0
				errorList = append(errorList, "skip overflow")
				success = false
			} else {
				toOut = toOut.([]server.BuildOrPullLog)[skip:]
				length = int64(len(toOut.([]server.BuildOrPullLog)))
			}

			if length > 0 {
				if limit > length && limit > 0 {
					toOut = toOut.([]server.BuildOrPullLog)[:length]
					length = int64(len(toOut.([]server.BuildOrPullLog)))
				} else if limit > 0 {
					toOut = toOut.([]server.BuildOrPullLog)[:limit]
					length = int64(len(toOut.([]server.BuildOrPullLog)))
				} else {
					toOut = toOut.([]server.BuildOrPullLog)
					length = int64(len(toOut.([]server.BuildOrPullLog)))
				}
			}
		}

	case pb.ImageOrContainerBuildPullStatusReply:
		var list = []string{
			data.(*pb.ImageOrContainerBuildPullStatusReply).String(),
		}
		toOut = list
		length = 1

		if skip >= length {
			toOut = make([]int, 0)
			length = 0
			errorList = append(errorList, "skip overflow")
			success = false
		} else {
			toOut = toOut.([]string)[skip:]
			length = int64(len(toOut.([]string)))
		}

		if length > 0 {
			if limit > length && limit > 0 {
				toOut = toOut.([]string)[:length]
				length = int64(len(toOut.([]string)))
			} else if limit > 0 {
				toOut = toOut.([]string)[:limit]
				length = int64(len(toOut.([]string)))
			} else {
				toOut = toOut.([]string)
				length = int64(len(toOut.([]string)))
			}
		}

	case pb.ImageListReply:
		var list []types.ImageSummary
		err = json.Unmarshal(data.(*pb.ImageListReply).Data, &list)
		if err != nil {
			length = 0
			limit = 0
			skip = 0
			success = false
			errorList = append(errorList, err.Error())
			toOut = make([]int, 0)
		} else {
			toOut = list

			length = int64(len(toOut.([]types.ImageSummary)))

			if skip >= length {
				toOut = make([]int, 0)
				length = 0
				errorList = append(errorList, "skip overflow")
				success = false
			} else {
				toOut = toOut.([]types.ImageSummary)[skip:]
				length = int64(len(toOut.([]types.ImageSummary)))
			}

			if length > 0 {
				if limit > length && limit > 0 {
					toOut = toOut.([]types.ImageSummary)[:length]
					length = int64(len(toOut.([]types.ImageSummary)))
				} else if limit > 0 {
					toOut = toOut.([]types.ImageSummary)[:limit]
					length = int64(len(toOut.([]types.ImageSummary)))
				} else {
					toOut = toOut.([]types.ImageSummary)
					length = int64(len(toOut.([]types.ImageSummary)))
				}
			}
		}

	case pb.NetworkListReply:
		var list []types.NetworkResource
		err = json.Unmarshal(data.(*pb.NetworkListReply).Data, &list)
		if err != nil {
			length = 0
			limit = 0
			skip = 0
			success = false
			errorList = append(errorList, err.Error())
			toOut = make([]int, 0)
		} else {
			toOut = list

			length = int64(len(toOut.([]types.NetworkResource)))

			if skip >= length {
				toOut = make([]int, 0)
				length = 0
				errorList = append(errorList, "skip overflow")
				success = false
			} else {
				toOut = toOut.([]types.NetworkResource)[skip:]
				length = int64(len(toOut.([]types.NetworkResource)))
			}

			if length > 0 {
				if limit > length && limit > 0 {
					toOut = toOut.([]types.NetworkResource)[:length]
					length = int64(len(toOut.([]types.NetworkResource)))
				} else if limit > 0 {
					toOut = toOut.([]types.NetworkResource)[:limit]
					length = int64(len(toOut.([]types.NetworkResource)))
				} else {
					toOut = toOut.([]types.NetworkResource)
					length = int64(len(toOut.([]types.NetworkResource)))
				}
			}
		}

	case pb.ContainerListAllReply:
		var list []types.Container
		err = json.Unmarshal(data.(*pb.ContainerListAllReply).Data, &list)
		if err != nil {
			length = 0
			limit = 0
			skip = 0
			success = false
			errorList = append(errorList, err.Error())
			toOut = make([]int, 0)
		} else {
			toOut = list

			length = int64(len(toOut.([]types.Container)))

			if skip >= length {
				toOut = make([]int, 0)
				length = 0
				errorList = append(errorList, "skip overflow")
				success = false
			} else {
				toOut = toOut.([]types.Container)[skip:]
				length = int64(len(toOut.([]types.Container)))
			}

			if length > 0 {
				if limit > length && limit > 0 {
					toOut = toOut.([]types.Container)[:length]
					length = int64(len(toOut.([]types.Container)))
				} else if limit > 0 {
					toOut = toOut.([]types.Container)[:limit]
					length = int64(len(toOut.([]types.Container)))
				} else {
					toOut = toOut.([]types.Container)
					length = int64(len(toOut.([]types.Container)))
				}
			}
		}

	case pb.ContainerStatisticsOneShotReply:
		var list types.Stats
		err = json.Unmarshal(data.(*pb.ContainerStatisticsOneShotReply).Data, &list)
		if err != nil {
			length = 0
			limit = 0
			skip = 0
			success = false
			errorList = append(errorList, err.Error())
			toOut = make([]int, 0)
		} else {
			toOut = []types.Stats{
				list,
			}

			length = 1

			if skip >= length {
				toOut = make([]int, 0)
				length = 0
				errorList = append(errorList, "skip overflow")
				success = false
			} else {
				toOut = toOut.([]types.Stats)[skip:]
				length = int64(len(toOut.([]types.Stats)))
			}

			if length > 0 {
				if limit > length && limit > 0 {
					toOut = toOut.([]types.Stats)[:length]
					length = int64(len(toOut.([]types.Stats)))
				} else if limit > 0 {
					toOut = toOut.([]types.Stats)[:limit]
					length = int64(len(toOut.([]types.Stats)))
				} else {
					toOut = toOut.([]types.Stats)
					length = int64(len(toOut.([]types.Stats)))
				}
			}
		}

	case pb.ContainerInspectReply:
		toOut = make([]types.ContainerJSON, 0)
		var container types.ContainerJSON
		err = json.Unmarshal(data.([]byte), &container)
		if err != nil {
			length = 0
			limit = 0
			skip = 0
			success = false
			errorList = append(errorList, err.Error())
			toOut = make([]int, 0)
		} else {
			toOut = append(toOut.([]types.ContainerJSON), container)

			length = int64(len(toOut.([]types.ContainerJSON)))

			if skip >= length {
				toOut = make([]int, 0)
				length = 0
				errorList = append(errorList, "skip overflow")
				success = false
			} else {
				toOut = toOut.([]types.ContainerJSON)[skip:]
				length = int64(len(toOut.([]types.ContainerJSON)))
			}

			if length > 0 {
				if limit > length && limit > 0 {
					toOut = toOut.([]types.ContainerJSON)[:length]
					length = int64(len(toOut.([]types.ContainerJSON)))
				} else if limit > 0 {
					toOut = toOut.([]types.ContainerJSON)[:limit]
					length = int64(len(toOut.([]types.ContainerJSON)))
				} else {
					toOut = toOut.([]types.ContainerJSON)
					length = int64(len(toOut.([]types.ContainerJSON)))
				}
			}
		}

	case pb.ContainerInspectByNameReply:
		toOut = make([]types.ContainerJSON, 0)
		var container types.ContainerJSON
		err = json.Unmarshal(data.([]byte), &container)
		if err != nil {
			length = 0
			limit = 0
			skip = 0
			success = false
			errorList = append(errorList, err.Error())
			toOut = make([]int, 0)
		} else {
			toOut = append(toOut.([]types.ContainerJSON), container)

			length = int64(len(toOut.([]types.ContainerJSON)))

			if skip >= length {
				toOut = make([]int, 0)
				length = 0
				errorList = append(errorList, "skip overflow")
				success = false
			} else {
				toOut = toOut.([]types.ContainerJSON)[skip:]
				length = int64(len(toOut.([]types.ContainerJSON)))
			}

			if length > 0 {
				if limit > length && limit > 0 {
					toOut = toOut.([]types.ContainerJSON)[:length]
					length = int64(len(toOut.([]types.ContainerJSON)))
				} else if limit > 0 {
					toOut = toOut.([]types.ContainerJSON)[:limit]
					length = int64(len(toOut.([]types.ContainerJSON)))
				} else {
					toOut = toOut.([]types.ContainerJSON)
					length = int64(len(toOut.([]types.ContainerJSON)))
				}
			}
		}

	case pb.ContainerInspectByNameContainsReply:
		toOut = make([]types.ContainerJSON, 0)
		var container []types.ContainerJSON
		err = json.Unmarshal(data.([]byte), &container)
		if err != nil {
			length = 0
			limit = 0
			skip = 0
			success = false
			errorList = append(errorList, err.Error())
			toOut = make([]int, 0)
		} else {
			toOut = container

			length = int64(len(toOut.([]types.ContainerJSON)))

			if skip >= length {
				toOut = make([]int, 0)
				length = 0
				errorList = append(errorList, "skip overflow")
				success = false
			} else {
				toOut = toOut.([]types.ContainerJSON)[skip:]
				length = int64(len(toOut.([]types.ContainerJSON)))
			}

			if length > 0 {
				if limit > length && limit > 0 {
					toOut = toOut.([]types.ContainerJSON)[:length]
					length = int64(len(toOut.([]types.ContainerJSON)))
				} else if limit > 0 {
					toOut = toOut.([]types.ContainerJSON)[:limit]
					length = int64(len(toOut.([]types.ContainerJSON)))
				} else {
					toOut = toOut.([]types.ContainerJSON)
					length = int64(len(toOut.([]types.ContainerJSON)))
				}
			}
		}

	default:
		fmt.Print("tojson.error: type not found\n")
	}

	toJsonOut := Output{
		Length:  length,
		Limit:   limit,
		Skip:    skip,
		Success: success,
		Error:   errorList,
		Data:    toOut,
	}

	output, err = json.Marshal(&toJsonOut)
	if err != nil {
		toJsonOut.Length = 0
		toJsonOut.Skip = 0
		toJsonOut.Limit = 0
		toJsonOut.Success = false
		toJsonOut.Data = make([]int, 0)
		toJsonOut.Error = append(toJsonOut.Error, fmt.Sprintf("json marshal error: %v", err.Error()))

		output, err = json.Marshal(&toJsonOut)
		if err != nil {
			panic(err)
		}
	}

	_, err = w.Write(output)
	if err != nil {
		panic(err)
	}
}

func ImageBuildFromRemoteServer(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ImageBuildFromRemoteServerReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ImageBuildFromRemoteServer(
		ctx,
		&pb.ImageBuildFromRemoteServerRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ImageBuildFromRemoteServerReply{}, container, w, r)
}

func ImageBuildAndContainerStartFromRemoteServer(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ImageBuildFromRemoteServerReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ImageBuildAndContainerStartFromRemoteServer(
		ctx,
		&pb.ImageBuildFromRemoteServerRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ImageBuildFromRemoteServerReply{}, container, w, r)
}

func ImageBuildFromRemoteServerStatus(w http.ResponseWriter, r *http.Request) {
	var err error
	var status *pb.ImageOrContainerBuildPullStatusReply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	status, err = GRpcClient.ImageBuildFromRemoteServerStatus(
		ctx,
		&pb.ImageOrContainerBuildPullStatusRequest{
			ID: r.URL.Query().Get("ID"),
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ImageOrContainerBuildPullStatusRequest{}, status, w, r)
}

func ContainerStopAndRemove(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	_, err = GRpcClient.ContainerStopAndRemove(
		ctx,
		&pb.ContainerStopAndRemoveRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.Empty{}, nil, w, r)
}

func ContainerRemove(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	_, err = GRpcClient.ContainerRemove(
		ctx,
		&pb.ContainerRemoveRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.Empty{}, nil, w, r)
}

func ContainerStop(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	_, err = GRpcClient.ContainerStop(
		ctx,
		&pb.ContainerStopRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.Empty{}, nil, w, r)
}

func ContainerStart(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	_, err = GRpcClient.ContainerStart(
		ctx,
		&pb.ContainerStartRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.Empty{}, nil, w, r)
}

func ContainerCreateAndChangeExposedPort(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerCreateAndChangeExposedPortReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerCreateAndChangeExposedPort(
		ctx,
		&pb.ContainerCreateAndChangeExposedPortRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerCreateAndChangeExposedPortReply{}, container, w, r)
}

func ContainerCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerCreateReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerCreate(
		ctx,
		&pb.ContainerCreateRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerCreateReply{}, container, w, r)
}

func ContainerCreateAndStart(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerCreateAndStartReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerCreateAndStart(
		ctx,
		&pb.ContainerCreateAndStartRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerCreateAndStartReply{}, container, w, r)
}

func ContainerStatisticsOneShot(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerStatisticsOneShotReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12000)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerStatisticsOneShot(
		ctx,
		&pb.ContainerStatisticsOneShotRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerStatisticsOneShotReply{}, container, w, r)
}

func ContainerStatisticsOneShotByName(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerStatisticsOneShotReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12000)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerStatisticsOneShotByName(
		ctx,
		&pb.ContainerStatisticsOneShotByNameRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerStatisticsOneShotReply{}, container, w, r)
}

func ContainerCreateChangeExposedPortAndStart(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerCreateChangeExposedPortAndStartReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerCreateChangeExposedPortAndStart(
		ctx,
		&pb.ContainerCreateChangeExposedPortAndStartRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerCreateChangeExposedPortAndStartReply{}, container, w, r)
}

func ContainerCreateAndExposePortsAutomatically(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerCreateAndExposePortsAutomaticallyReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerCreateAndExposePortsAutomatically(
		ctx,
		&pb.ContainerCreateAndExposePortsAutomaticallyRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerCreateAndExposePortsAutomaticallyReply{}, container, w, r)
}

func ContainerCreateWithoutExposePorts(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerCreateWithoutExposePortsReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerCreateWithoutExposePorts(
		ctx,
		&pb.ContainerCreateWithoutExposePortsRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerCreateWithoutExposePortsReply{}, container, w, r)
}

func ContainerCreateWithoutExposePortsAndStart(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerCreateWithoutExposePortsAndStartReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerCreateWithoutExposePortsAndStart(
		ctx,
		&pb.ContainerCreateWithoutExposePortsAndStartRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerCreateWithoutExposePortsAndStartReply{}, container, w, r)
}

func ContainerCreateExposePortsAutomaticallyAndStart(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerCreateExposePortsAutomaticallyAndStartReply
	var body []byte

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("body.err: %v", err.Error())
		return
	}

	container, err = GRpcClient.ContainerCreateExposePortsAutomaticallyAndStart(
		ctx,
		&pb.ContainerCreateExposePortsAutomaticallyAndStartRequest{
			Data: body,
		},
	)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}
	ToJson(pb.ContainerCreateExposePortsAutomaticallyAndStartReply{}, container, w, r)
}

func ContainerFindIdByName(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerFindIdByNameReply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	container, err = GRpcClient.ContainerFindIdByName(ctx, &pb.ContainerFindIdByNameRequest{
		Name: r.URL.Query().Get("Name"),
	})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}

	ToJson(pb.ContainerFindIdByNameReply{}, container, w, r)
}

func ContainerInspectByName(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerInspectByNameReply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	container, err = GRpcClient.ContainerInspectByName(ctx, &pb.ContainerInspectByNameRequest{
		Name: r.URL.Query().Get("Name"),
	})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}

	ToJson(pb.ContainerInspectByNameReply{}, container.Data, w, r)
}

func ContainerFindIdByNameContains(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerFindIdByNameContainsReply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	container, err = GRpcClient.ContainerFindIdByNameContains(ctx, &pb.ContainerFindIdByNameContainsRequest{
		Name: r.URL.Query().Get("Name"),
	})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}

	ToJson(pb.ContainerFindIdByNameContainsReply{}, container, w, r)
}

func ContainerInspectByNameContains(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerInspectByNameContainsReply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	container, err = GRpcClient.ContainerInspectByNameContains(ctx, &pb.ContainerInspectByNameContainsRequest{
		Name: r.URL.Query().Get("Name"),
	})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}

	ToJson(pb.ContainerInspectByNameContainsReply{}, container.Data, w, r)
}

func ContainersList(w http.ResponseWriter, r *http.Request) {
	var err error
	var list *pb.ContainerListAllReply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	list, err = GRpcClient.ContainerListAll(ctx, &pb.Empty{})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}

	ToJson(pb.ContainerListAllReply{}, list, w, r)
}

func ContainerInspect(w http.ResponseWriter, r *http.Request) {
	var err error
	var container *pb.ContainerInspectReply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	container, err = GRpcClient.ContainerInspect(ctx, &pb.ContainerInspectRequest{
		ID: r.URL.Query().Get("ID"),
	})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}

	ToJson(pb.ContainerInspectReply{}, container.Data, w, r)
}

func ImageList(w http.ResponseWriter, r *http.Request) {
	var err error
	var list *pb.ImageListReply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	list, err = GRpcClient.ImageList(ctx, &pb.Empty{})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}

	ToJson(pb.ImageListReply{}, list, w, r)
}

func NetworkList(w http.ResponseWriter, r *http.Request) {
	var err error
	var list *pb.NetworkListReply

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	list, err = GRpcClient.NetworkList(ctx, &pb.Empty{})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return
	}

	ToJson(pb.NetworkListReply{}, list, w, r)
}

func HeaderWrite(w http.ResponseWriter, headerType string) {
	w.Header().Set("Content-Type", headerType)
	w.Header().Add("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Pragma", "no-cache")
}

func TemplateLoadDir(path string) (templateList []string, err error) {
	var dirDara []os.FileInfo
	dirDara, err = ioutil.ReadDir(path)
	if err != nil {
		return
	}

	templateList = make([]string, 0)
	for _, file := range dirDara {
		if file.IsDir() == false && strings.HasSuffix(file.Name(), ".tpl") {
			templateList = append(templateList, filepath.Join(path, file.Name()))
		}
	}

	return
}

func TemplateParser(w http.ResponseWriter, name string, data interface{}) (err error) {
	var templateDir = filepath.Join("static", "templates", "")
	var tmpl *template.Template
	var list []string
	var functions = map[string]interface{}{
		"toUp": strings.ToUpper,
	}

	tmpl = template.New("template")
	tmpl.Funcs(functions)
	list, err = TemplateLoadDir(templateDir)
	if err != nil {
		return
	}

	tmpl, err = tmpl.ParseFiles(list...)
	if err != nil {
		return
	}

	err = tmpl.ExecuteTemplate(w, name, data)
	return
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	return
	_ = r

	var err error
	HeaderWrite(w, KHeaderTypeHtmlUtf8)
	err = TemplateParser(w, "page", nil)
	if err != nil {
		panic(err)
	}
}
