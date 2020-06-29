//protoc --go_opt=paths=source_relative --go_out=plugins=grpc:. *.proto
package main

import (
	"bytes"
	"errors"
	"fmt"
	goToGrpc "github.com/helmutkemper/iotmaker.util.grpc.goToGrpc"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func main() {

	/*
	  toFile := bytes.Buffer{}
	  a := []string{"bool", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16",
	    "uint32", "uint64", "uintptr", "float32", "float64", "interface{}", "string"}
	  for _, v := range a {
	    for _, v2 := range a {
	      toFile.WriteString( fmt.Sprintf("case map[%v]%v:\n", v, v2) )
	      toFile.WriteString( fmt.Sprintf("  keyType = \"%v\"\n", v) )
	      toFile.WriteString( fmt.Sprintf("  keyValue = \"%v\"\n", v2) )
	    }
	  }

	  ioutil.WriteFile("./out.txt", toFile.Bytes(), os.ModePerm)

	  os.Exit(0)
	*/

	var err error
	var content = []byte(`
syntax = "proto3";

option go_package = "github.com/helmutkemper/iotmaker_docker_communication_grpc";

package iotmakerDockerCommunicationGrpc;

`)

	err = ioutil.WriteFile("./out.proto", content, os.ModePerm)
	if err != nil {
		panic(err)
	}

	var a1 goToGrpc.ContainerJSON
	test(&a1)
	var a2 goToGrpc.Mount
	test(&a2)
	var a3 goToGrpc.HealthcheckResult
	test(&a3)
	var a4 goToGrpc.WeightDevice
	test(&a4)
	var a5 goToGrpc.ThrottleDevice
	test(&a5)
	var a6 goToGrpc.DeviceMapping
	test(&a6)
	var a7 goToGrpc.DeviceRequest
	test(&a7)
	var a8 goToGrpc.Ulimit
	test(&a8)
	var a9 goToGrpc.MountPoint
	test(&a9)
	var b1 goToGrpc.Address
	test(&b1)
	var b2 goToGrpc.EndpointSettings
	test(&b2)
	var b3 goToGrpc.PortBinding
	test(&b3)

}

func test(i interface{}) {
	var err error
	var file *os.File
	var buffer bytes.Buffer

	elementName := reflect.ValueOf(i).Elem().Type().Name()
	element := reflect.ValueOf(i).Elem()

	buffer.WriteString(fmt.Sprintf("message %v {\n", elementName))

	for i := 0; i < element.NumField(); i += 1 {
		field := element.Field(i)
		nameOfField := element.Type().Field(i).Name

		err = ToScalarValue(&buffer, field, nameOfField, i)
		if err != nil {
			panic(err)
		}
	}

	buffer.WriteString("}\n")
	buffer.WriteString("\n")

	file, err = os.OpenFile("./out.proto", os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
}

func ToScalarValue(
	buffer *bytes.Buffer,
	element reflect.Value,
	nameOfField string,
	i int,
) (
	err error,
) {

	nameOfField = strings.Replace(nameOfField, "main.", "", -1)
	var elementTypeText = element.Type().String()
	elementTypeText = strings.Replace(elementTypeText, "main.", "", -1)
	switch element.Type().Kind() {
	case reflect.Invalid:
		err = errors.New("ToScalarValue() function found an invalid value")
	case reflect.Bool:
		buffer.WriteString("  " + elementTypeText + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uintptr:
		buffer.WriteString("  " + removePtrFromString(elementTypeText) + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Complex64:
		err = errors.New("ToScalarValue() function has't Complex64 code")
	case reflect.Complex128:
		err = errors.New("ToScalarValue() function has't Complex128 code")
	case reflect.Array:
		//err = errors.New("ToScalarValue() function has't array code. >"+nameOfField)
	case reflect.Chan:
		break
	case reflect.Func:
		break
	case reflect.Interface:
		buffer.WriteString("  interface " + elementTypeText + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Map:
		mapConverter(buffer, element, nameOfField, i)
	case reflect.Struct:
		err = ToStructType(element)
		buffer.WriteString("  " + elementTypeText + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.UnsafePointer:
		break
	case reflect.String:
		buffer.WriteString("  string " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint:
		buffer.WriteString("  uint64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint8:
		buffer.WriteString("  uint32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint16:
		buffer.WriteString("  uint32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint32:
		buffer.WriteString("  uint32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Uint64:
		buffer.WriteString("  uint64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int:
		buffer.WriteString("  int64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int8:
		buffer.WriteString("  int8 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int16:
		buffer.WriteString("  int16 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int32:
		buffer.WriteString("  int32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Int64:
		buffer.WriteString("  int64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Float32:
		buffer.WriteString("  float " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Float64:
		buffer.WriteString("  double " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Slice:
		buffer.WriteString("  repeated " + removeArrFromString(elementTypeText) + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case reflect.Ptr:
		buffer.WriteString("  " + removePtrFromString(elementTypeText) + " " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	}

	return
}

func ToStructType(
	//buffer *bytes.Buffer,
	element reflect.Value,
) (
	err error,
) {

	var file *os.File
	var buffer bytes.Buffer

	//var elementTypeText = element.Type().String()
	var nameOfStruct = element.Type().Name()
	_ = nameOfStruct
	t := element
	_ = t

	buffer.WriteString(fmt.Sprintf("message %v {\n", nameOfStruct))

	for i := 0; i < element.NumField(); i += 1 {
		nameOfField := element.Type().Field(i).Name
		field := element.Field(i)
		err = ToScalarValue(&buffer, field, nameOfField, i)
		if err != nil {
			return
		}
	}

	buffer.WriteString("}\n")
	buffer.WriteString("\n")

	file, err = os.OpenFile("./out.proto", os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	return
}

func removeArrFromString(t string) (convertedType string) {
	return strings.Replace(t, "[]", "", 1)
}

func removePtrFromString(t string) (convertedType string) {
	return strings.Replace(t, "*", "", 1)
}

func mapConverter(buffer *bytes.Buffer, element reflect.Value, nameOfField string, i int) {
	var keyType, keyValue string

	switch element.Interface().(type) {
	case map[bool]bool:
		keyType = "bool"
		keyValue = "bool"
	case map[bool]int:
		keyType = "bool"
		keyValue = "int"
	case map[bool]int8:
		keyType = "bool"
		keyValue = "int8"
	case map[bool]int16:
		keyType = "bool"
		keyValue = "int16"
	case map[bool]int32:
		keyType = "bool"
		keyValue = "int32"
	case map[bool]int64:
		keyType = "bool"
		keyValue = "int64"
	case map[bool]uint:
		keyType = "bool"
		keyValue = "uint"
	case map[bool]uint8:
		keyType = "bool"
		keyValue = "uint8"
	case map[bool]uint16:
		keyType = "bool"
		keyValue = "uint16"
	case map[bool]uint32:
		keyType = "bool"
		keyValue = "uint32"
	case map[bool]uint64:
		keyType = "bool"
		keyValue = "uint64"
	case map[bool]uintptr:
		keyType = "bool"
		keyValue = "uintptr"
	case map[bool]float32:
		keyType = "bool"
		keyValue = "float32"
	case map[bool]float64:
		keyType = "bool"
		keyValue = "float64"
	case map[bool]interface{}:
		keyType = "bool"
		keyValue = "interface{}"
	case map[bool]string:
		keyType = "bool"
		keyValue = "string"
	case map[int]bool:
		keyType = "int"
		keyValue = "bool"
	case map[int]int:
		keyType = "int"
		keyValue = "int"
	case map[int]int8:
		keyType = "int"
		keyValue = "int8"
	case map[int]int16:
		keyType = "int"
		keyValue = "int16"
	case map[int]int32:
		keyType = "int"
		keyValue = "int32"
	case map[int]int64:
		keyType = "int"
		keyValue = "int64"
	case map[int]uint:
		keyType = "int"
		keyValue = "uint"
	case map[int]uint8:
		keyType = "int"
		keyValue = "uint8"
	case map[int]uint16:
		keyType = "int"
		keyValue = "uint16"
	case map[int]uint32:
		keyType = "int"
		keyValue = "uint32"
	case map[int]uint64:
		keyType = "int"
		keyValue = "uint64"
	case map[int]uintptr:
		keyType = "int"
		keyValue = "uintptr"
	case map[int]float32:
		keyType = "int"
		keyValue = "float32"
	case map[int]float64:
		keyType = "int"
		keyValue = "float64"
	case map[int]interface{}:
		keyType = "int"
		keyValue = "interface{}"
	case map[int]string:
		keyType = "int"
		keyValue = "string"
	case map[int8]bool:
		keyType = "int8"
		keyValue = "bool"
	case map[int8]int:
		keyType = "int8"
		keyValue = "int"
	case map[int8]int8:
		keyType = "int8"
		keyValue = "int8"
	case map[int8]int16:
		keyType = "int8"
		keyValue = "int16"
	case map[int8]int32:
		keyType = "int8"
		keyValue = "int32"
	case map[int8]int64:
		keyType = "int8"
		keyValue = "int64"
	case map[int8]uint:
		keyType = "int8"
		keyValue = "uint"
	case map[int8]uint8:
		keyType = "int8"
		keyValue = "uint8"
	case map[int8]uint16:
		keyType = "int8"
		keyValue = "uint16"
	case map[int8]uint32:
		keyType = "int8"
		keyValue = "uint32"
	case map[int8]uint64:
		keyType = "int8"
		keyValue = "uint64"
	case map[int8]uintptr:
		keyType = "int8"
		keyValue = "uintptr"
	case map[int8]float32:
		keyType = "int8"
		keyValue = "float32"
	case map[int8]float64:
		keyType = "int8"
		keyValue = "float64"
	case map[int8]interface{}:
		keyType = "int8"
		keyValue = "interface{}"
	case map[int8]string:
		keyType = "int8"
		keyValue = "string"
	case map[int16]bool:
		keyType = "int16"
		keyValue = "bool"
	case map[int16]int:
		keyType = "int16"
		keyValue = "int"
	case map[int16]int8:
		keyType = "int16"
		keyValue = "int8"
	case map[int16]int16:
		keyType = "int16"
		keyValue = "int16"
	case map[int16]int32:
		keyType = "int16"
		keyValue = "int32"
	case map[int16]int64:
		keyType = "int16"
		keyValue = "int64"
	case map[int16]uint:
		keyType = "int16"
		keyValue = "uint"
	case map[int16]uint8:
		keyType = "int16"
		keyValue = "uint8"
	case map[int16]uint16:
		keyType = "int16"
		keyValue = "uint16"
	case map[int16]uint32:
		keyType = "int16"
		keyValue = "uint32"
	case map[int16]uint64:
		keyType = "int16"
		keyValue = "uint64"
	case map[int16]uintptr:
		keyType = "int16"
		keyValue = "uintptr"
	case map[int16]float32:
		keyType = "int16"
		keyValue = "float32"
	case map[int16]float64:
		keyType = "int16"
		keyValue = "float64"
	case map[int16]interface{}:
		keyType = "int16"
		keyValue = "interface{}"
	case map[int16]string:
		keyType = "int16"
		keyValue = "string"
	case map[int32]bool:
		keyType = "int32"
		keyValue = "bool"
	case map[int32]int:
		keyType = "int32"
		keyValue = "int"
	case map[int32]int8:
		keyType = "int32"
		keyValue = "int8"
	case map[int32]int16:
		keyType = "int32"
		keyValue = "int16"
	case map[int32]int32:
		keyType = "int32"
		keyValue = "int32"
	case map[int32]int64:
		keyType = "int32"
		keyValue = "int64"
	case map[int32]uint:
		keyType = "int32"
		keyValue = "uint"
	case map[int32]uint8:
		keyType = "int32"
		keyValue = "uint8"
	case map[int32]uint16:
		keyType = "int32"
		keyValue = "uint16"
	case map[int32]uint32:
		keyType = "int32"
		keyValue = "uint32"
	case map[int32]uint64:
		keyType = "int32"
		keyValue = "uint64"
	case map[int32]uintptr:
		keyType = "int32"
		keyValue = "uintptr"
	case map[int32]float32:
		keyType = "int32"
		keyValue = "float32"
	case map[int32]float64:
		keyType = "int32"
		keyValue = "float64"
	case map[int32]interface{}:
		keyType = "int32"
		keyValue = "interface{}"
	case map[int32]string:
		keyType = "int32"
		keyValue = "string"
	case map[int64]bool:
		keyType = "int64"
		keyValue = "bool"
	case map[int64]int:
		keyType = "int64"
		keyValue = "int"
	case map[int64]int8:
		keyType = "int64"
		keyValue = "int8"
	case map[int64]int16:
		keyType = "int64"
		keyValue = "int16"
	case map[int64]int32:
		keyType = "int64"
		keyValue = "int32"
	case map[int64]int64:
		keyType = "int64"
		keyValue = "int64"
	case map[int64]uint:
		keyType = "int64"
		keyValue = "uint"
	case map[int64]uint8:
		keyType = "int64"
		keyValue = "uint8"
	case map[int64]uint16:
		keyType = "int64"
		keyValue = "uint16"
	case map[int64]uint32:
		keyType = "int64"
		keyValue = "uint32"
	case map[int64]uint64:
		keyType = "int64"
		keyValue = "uint64"
	case map[int64]uintptr:
		keyType = "int64"
		keyValue = "uintptr"
	case map[int64]float32:
		keyType = "int64"
		keyValue = "float32"
	case map[int64]float64:
		keyType = "int64"
		keyValue = "float64"
	case map[int64]interface{}:
		keyType = "int64"
		keyValue = "interface{}"
	case map[int64]string:
		keyType = "int64"
		keyValue = "string"
	case map[uint]bool:
		keyType = "uint"
		keyValue = "bool"
	case map[uint]int:
		keyType = "uint"
		keyValue = "int"
	case map[uint]int8:
		keyType = "uint"
		keyValue = "int8"
	case map[uint]int16:
		keyType = "uint"
		keyValue = "int16"
	case map[uint]int32:
		keyType = "uint"
		keyValue = "int32"
	case map[uint]int64:
		keyType = "uint"
		keyValue = "int64"
	case map[uint]uint:
		keyType = "uint"
		keyValue = "uint"
	case map[uint]uint8:
		keyType = "uint"
		keyValue = "uint8"
	case map[uint]uint16:
		keyType = "uint"
		keyValue = "uint16"
	case map[uint]uint32:
		keyType = "uint"
		keyValue = "uint32"
	case map[uint]uint64:
		keyType = "uint"
		keyValue = "uint64"
	case map[uint]uintptr:
		keyType = "uint"
		keyValue = "uintptr"
	case map[uint]float32:
		keyType = "uint"
		keyValue = "float32"
	case map[uint]float64:
		keyType = "uint"
		keyValue = "float64"
	case map[uint]interface{}:
		keyType = "uint"
		keyValue = "interface{}"
	case map[uint]string:
		keyType = "uint"
		keyValue = "string"
	case map[uint8]bool:
		keyType = "uint8"
		keyValue = "bool"
	case map[uint8]int:
		keyType = "uint8"
		keyValue = "int"
	case map[uint8]int8:
		keyType = "uint8"
		keyValue = "int8"
	case map[uint8]int16:
		keyType = "uint8"
		keyValue = "int16"
	case map[uint8]int32:
		keyType = "uint8"
		keyValue = "int32"
	case map[uint8]int64:
		keyType = "uint8"
		keyValue = "int64"
	case map[uint8]uint:
		keyType = "uint8"
		keyValue = "uint"
	case map[uint8]uint8:
		keyType = "uint8"
		keyValue = "uint8"
	case map[uint8]uint16:
		keyType = "uint8"
		keyValue = "uint16"
	case map[uint8]uint32:
		keyType = "uint8"
		keyValue = "uint32"
	case map[uint8]uint64:
		keyType = "uint8"
		keyValue = "uint64"
	case map[uint8]uintptr:
		keyType = "uint8"
		keyValue = "uintptr"
	case map[uint8]float32:
		keyType = "uint8"
		keyValue = "float32"
	case map[uint8]float64:
		keyType = "uint8"
		keyValue = "float64"
	case map[uint8]interface{}:
		keyType = "uint8"
		keyValue = "interface{}"
	case map[uint8]string:
		keyType = "uint8"
		keyValue = "string"
	case map[uint16]bool:
		keyType = "uint16"
		keyValue = "bool"
	case map[uint16]int:
		keyType = "uint16"
		keyValue = "int"
	case map[uint16]int8:
		keyType = "uint16"
		keyValue = "int8"
	case map[uint16]int16:
		keyType = "uint16"
		keyValue = "int16"
	case map[uint16]int32:
		keyType = "uint16"
		keyValue = "int32"
	case map[uint16]int64:
		keyType = "uint16"
		keyValue = "int64"
	case map[uint16]uint:
		keyType = "uint16"
		keyValue = "uint"
	case map[uint16]uint8:
		keyType = "uint16"
		keyValue = "uint8"
	case map[uint16]uint16:
		keyType = "uint16"
		keyValue = "uint16"
	case map[uint16]uint32:
		keyType = "uint16"
		keyValue = "uint32"
	case map[uint16]uint64:
		keyType = "uint16"
		keyValue = "uint64"
	case map[uint16]uintptr:
		keyType = "uint16"
		keyValue = "uintptr"
	case map[uint16]float32:
		keyType = "uint16"
		keyValue = "float32"
	case map[uint16]float64:
		keyType = "uint16"
		keyValue = "float64"
	case map[uint16]interface{}:
		keyType = "uint16"
		keyValue = "interface{}"
	case map[uint16]string:
		keyType = "uint16"
		keyValue = "string"
	case map[uint32]bool:
		keyType = "uint32"
		keyValue = "bool"
	case map[uint32]int:
		keyType = "uint32"
		keyValue = "int"
	case map[uint32]int8:
		keyType = "uint32"
		keyValue = "int8"
	case map[uint32]int16:
		keyType = "uint32"
		keyValue = "int16"
	case map[uint32]int32:
		keyType = "uint32"
		keyValue = "int32"
	case map[uint32]int64:
		keyType = "uint32"
		keyValue = "int64"
	case map[uint32]uint:
		keyType = "uint32"
		keyValue = "uint"
	case map[uint32]uint8:
		keyType = "uint32"
		keyValue = "uint8"
	case map[uint32]uint16:
		keyType = "uint32"
		keyValue = "uint16"
	case map[uint32]uint32:
		keyType = "uint32"
		keyValue = "uint32"
	case map[uint32]uint64:
		keyType = "uint32"
		keyValue = "uint64"
	case map[uint32]uintptr:
		keyType = "uint32"
		keyValue = "uintptr"
	case map[uint32]float32:
		keyType = "uint32"
		keyValue = "float32"
	case map[uint32]float64:
		keyType = "uint32"
		keyValue = "float64"
	case map[uint32]interface{}:
		keyType = "uint32"
		keyValue = "interface{}"
	case map[uint32]string:
		keyType = "uint32"
		keyValue = "string"
	case map[uint64]bool:
		keyType = "uint64"
		keyValue = "bool"
	case map[uint64]int:
		keyType = "uint64"
		keyValue = "int"
	case map[uint64]int8:
		keyType = "uint64"
		keyValue = "int8"
	case map[uint64]int16:
		keyType = "uint64"
		keyValue = "int16"
	case map[uint64]int32:
		keyType = "uint64"
		keyValue = "int32"
	case map[uint64]int64:
		keyType = "uint64"
		keyValue = "int64"
	case map[uint64]uint:
		keyType = "uint64"
		keyValue = "uint"
	case map[uint64]uint8:
		keyType = "uint64"
		keyValue = "uint8"
	case map[uint64]uint16:
		keyType = "uint64"
		keyValue = "uint16"
	case map[uint64]uint32:
		keyType = "uint64"
		keyValue = "uint32"
	case map[uint64]uint64:
		keyType = "uint64"
		keyValue = "uint64"
	case map[uint64]uintptr:
		keyType = "uint64"
		keyValue = "uintptr"
	case map[uint64]float32:
		keyType = "uint64"
		keyValue = "float32"
	case map[uint64]float64:
		keyType = "uint64"
		keyValue = "float64"
	case map[uint64]interface{}:
		keyType = "uint64"
		keyValue = "interface{}"
	case map[uint64]string:
		keyType = "uint64"
		keyValue = "string"
	case map[uintptr]bool:
		keyType = "uintptr"
		keyValue = "bool"
	case map[uintptr]int:
		keyType = "uintptr"
		keyValue = "int"
	case map[uintptr]int8:
		keyType = "uintptr"
		keyValue = "int8"
	case map[uintptr]int16:
		keyType = "uintptr"
		keyValue = "int16"
	case map[uintptr]int32:
		keyType = "uintptr"
		keyValue = "int32"
	case map[uintptr]int64:
		keyType = "uintptr"
		keyValue = "int64"
	case map[uintptr]uint:
		keyType = "uintptr"
		keyValue = "uint"
	case map[uintptr]uint8:
		keyType = "uintptr"
		keyValue = "uint8"
	case map[uintptr]uint16:
		keyType = "uintptr"
		keyValue = "uint16"
	case map[uintptr]uint32:
		keyType = "uintptr"
		keyValue = "uint32"
	case map[uintptr]uint64:
		keyType = "uintptr"
		keyValue = "uint64"
	case map[uintptr]uintptr:
		keyType = "uintptr"
		keyValue = "uintptr"
	case map[uintptr]float32:
		keyType = "uintptr"
		keyValue = "float32"
	case map[uintptr]float64:
		keyType = "uintptr"
		keyValue = "float64"
	case map[uintptr]interface{}:
		keyType = "uintptr"
		keyValue = "interface{}"
	case map[uintptr]string:
		keyType = "uintptr"
		keyValue = "string"
	case map[float32]bool:
		keyType = "float32"
		keyValue = "bool"
	case map[float32]int:
		keyType = "float32"
		keyValue = "int"
	case map[float32]int8:
		keyType = "float32"
		keyValue = "int8"
	case map[float32]int16:
		keyType = "float32"
		keyValue = "int16"
	case map[float32]int32:
		keyType = "float32"
		keyValue = "int32"
	case map[float32]int64:
		keyType = "float32"
		keyValue = "int64"
	case map[float32]uint:
		keyType = "float32"
		keyValue = "uint"
	case map[float32]uint8:
		keyType = "float32"
		keyValue = "uint8"
	case map[float32]uint16:
		keyType = "float32"
		keyValue = "uint16"
	case map[float32]uint32:
		keyType = "float32"
		keyValue = "uint32"
	case map[float32]uint64:
		keyType = "float32"
		keyValue = "uint64"
	case map[float32]uintptr:
		keyType = "float32"
		keyValue = "uintptr"
	case map[float32]float32:
		keyType = "float32"
		keyValue = "float32"
	case map[float32]float64:
		keyType = "float32"
		keyValue = "float64"
	case map[float32]interface{}:
		keyType = "float32"
		keyValue = "interface{}"
	case map[float32]string:
		keyType = "float32"
		keyValue = "string"
	case map[float64]bool:
		keyType = "float64"
		keyValue = "bool"
	case map[float64]int:
		keyType = "float64"
		keyValue = "int"
	case map[float64]int8:
		keyType = "float64"
		keyValue = "int8"
	case map[float64]int16:
		keyType = "float64"
		keyValue = "int16"
	case map[float64]int32:
		keyType = "float64"
		keyValue = "int32"
	case map[float64]int64:
		keyType = "float64"
		keyValue = "int64"
	case map[float64]uint:
		keyType = "float64"
		keyValue = "uint"
	case map[float64]uint8:
		keyType = "float64"
		keyValue = "uint8"
	case map[float64]uint16:
		keyType = "float64"
		keyValue = "uint16"
	case map[float64]uint32:
		keyType = "float64"
		keyValue = "uint32"
	case map[float64]uint64:
		keyType = "float64"
		keyValue = "uint64"
	case map[float64]uintptr:
		keyType = "float64"
		keyValue = "uintptr"
	case map[float64]float32:
		keyType = "float64"
		keyValue = "float32"
	case map[float64]float64:
		keyType = "float64"
		keyValue = "float64"
	case map[float64]interface{}:
		keyType = "float64"
		keyValue = "interface{}"
	case map[float64]string:
		keyType = "float64"
		keyValue = "string"
	case map[interface{}]bool:
		keyType = "interface{}"
		keyValue = "bool"
	case map[interface{}]int:
		keyType = "interface{}"
		keyValue = "int"
	case map[interface{}]int8:
		keyType = "interface{}"
		keyValue = "int8"
	case map[interface{}]int16:
		keyType = "interface{}"
		keyValue = "int16"
	case map[interface{}]int32:
		keyType = "interface{}"
		keyValue = "int32"
	case map[interface{}]int64:
		keyType = "interface{}"
		keyValue = "int64"
	case map[interface{}]uint:
		keyType = "interface{}"
		keyValue = "uint"
	case map[interface{}]uint8:
		keyType = "interface{}"
		keyValue = "uint8"
	case map[interface{}]uint16:
		keyType = "interface{}"
		keyValue = "uint16"
	case map[interface{}]uint32:
		keyType = "interface{}"
		keyValue = "uint32"
	case map[interface{}]uint64:
		keyType = "interface{}"
		keyValue = "uint64"
	case map[interface{}]uintptr:
		keyType = "interface{}"
		keyValue = "uintptr"
	case map[interface{}]float32:
		keyType = "interface{}"
		keyValue = "float32"
	case map[interface{}]float64:
		keyType = "interface{}"
		keyValue = "float64"
	case map[interface{}]interface{}:
		keyType = "interface{}"
		keyValue = "interface{}"
	case map[interface{}]string:
		keyType = "interface{}"
		keyValue = "string"
	case map[string]bool:
		keyType = "string"
		keyValue = "bool"
	case map[string]int:
		keyType = "string"
		keyValue = "int"
	case map[string]int8:
		keyType = "string"
		keyValue = "int8"
	case map[string]int16:
		keyType = "string"
		keyValue = "int16"
	case map[string]int32:
		keyType = "string"
		keyValue = "int32"
	case map[string]int64:
		keyType = "string"
		keyValue = "int64"
	case map[string]uint:
		keyType = "string"
		keyValue = "uint"
	case map[string]uint8:
		keyType = "string"
		keyValue = "uint8"
	case map[string]uint16:
		keyType = "string"
		keyValue = "uint16"
	case map[string]uint32:
		keyType = "string"
		keyValue = "uint32"
	case map[string]uint64:
		keyType = "string"
		keyValue = "uint64"
	case map[string]uintptr:
		keyType = "string"
		keyValue = "uintptr"
	case map[string]float32:
		keyType = "string"
		keyValue = "float32"
	case map[string]float64:
		keyType = "string"
		keyValue = "float64"
	case map[string]interface{}:
		keyType = "string"
		keyValue = "interface{}"
	case map[string]string:
		keyType = "string"
		keyValue = "string"
	}

	_ = keyType
	_ = keyValue

	buffer.WriteString("  map<" + keyType + ", " + keyValue + "> " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")

}
