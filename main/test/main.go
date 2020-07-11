package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/joncalhoun/qson"
	"github.com/sirupsen/logrus"
	"github.com/zang-cloud/grpc-json/jsonpb"
)

func main() {
	var DefaultMarshaler = &jsonpb.Marshaler{EnumsAsInts: true, EmitDefaults: true, OrigName: true, Int64AsString: false, Uint64AsString: false}
	DefaultMarshaler.MarshalToString()
}
