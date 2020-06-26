//protoc --go_opt=paths=source_relative --go_out=plugins=grpc:. *.proto
package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var t Test
	test(&t)
}

type Test struct {
	Gelada  string
	Piratas []string
	Complex
}

type Complex struct {
	Int int
}

func test(i interface{}) {
	var buffer bytes.Buffer

	elementName := reflect.ValueOf(i).Elem().Type().Name()
	element := reflect.ValueOf(i).Elem()

	buffer.WriteString(fmt.Sprintf("message %v {\n", elementName))

	for i := 0; i < element.NumField(); i += 1 {
		field := element.Field(i)
		nameOfField := element.Type().Field(i).Name

		err, localBuffer := ToScalarValue(field.Interface(), nameOfField, i)
		if err != nil {

			//nameOfField := field.Type().Name()
			for i := 0; i < field.NumField(); i += 1 {
				nameOfField := field.Type().Field(i).Name
				field := field.Field(i)
				err, localBuffer := ToScalarValue(field.Interface(), nameOfField, i)
				if err != nil {
					panic(err)
				}

				_, err = buffer.Write(localBuffer.Bytes())
				if err != nil {
					panic(err)
				}
			}

			continue
		}

		_, err = buffer.Write(localBuffer.Bytes())
		if err != nil {
			panic(err)
		}
	}

	buffer.WriteString("}\n")
	buffer.WriteString("\n")

	fmt.Printf("%v", buffer.String())
}

func ToScalarValue(
	element interface{},
	nameOfField string,
	i int,
) (
	err error,
	buffer bytes.Buffer,
) {

	switch element.(type) {
	default:
		err = errors.New(nameOfField + " is't a scalar value")
		return
	case string:
		buffer.WriteString("  string " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case uint:
		buffer.WriteString("  uint64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case uint8:
		buffer.WriteString("  uint8 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case uint16:
		buffer.WriteString("  uint16 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case uint32:
		buffer.WriteString("  uint32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case uint64:
		buffer.WriteString("  uint64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case int:
		buffer.WriteString("  int64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case int8:
		buffer.WriteString("  int8 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case int16:
		buffer.WriteString("  int16 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case int32:
		buffer.WriteString("  int32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case int64:
		buffer.WriteString("  int64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case float32:
		buffer.WriteString("  float " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case float64:
		buffer.WriteString("  double " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []string:
		buffer.WriteString("  repeated string " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []uint:
		buffer.WriteString("  repeated uint64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []uint8:
		buffer.WriteString("  repeated uint8 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []uint16:
		buffer.WriteString("  repeated uint16 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []uint32:
		buffer.WriteString("  repeated uint32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []uint64:
		buffer.WriteString("  repeated uint64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []int:
		buffer.WriteString("  repeated int64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []int8:
		buffer.WriteString("  repeated int8 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []int16:
		buffer.WriteString("  repeated int16 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []int32:
		buffer.WriteString("  repeated int32 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []int64:
		buffer.WriteString("  repeated int64 " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []float32:
		buffer.WriteString("  repeated float " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	case []float64:
		buffer.WriteString("  repeated double " + nameOfField + " = " + fmt.Sprintf("%v", i+1) + ";\n")
	}

	return
}
