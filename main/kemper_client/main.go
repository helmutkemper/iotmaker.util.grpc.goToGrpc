package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
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
	mux.HandleFunc("/containersListAll", ListContainers)
	mux.HandleFunc("/networkListAll", NetworkList)
	//  mux.HandleFunc("/ImageList", ImageList)
	mux.HandleFunc("/containerInspectById", ContainerInspect)

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

	default:
		fmt.Print("type not found\n")
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

func ListContainers(w http.ResponseWriter, r *http.Request) {
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

/*
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

  ToJson(list, w, r)
}
*/

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
	_ = r

	var err error
	HeaderWrite(w, KHeaderTypeHtmlUtf8)
	err = TemplateParser(w, "page", nil)
	if err != nil {
		panic(err)
	}
}
