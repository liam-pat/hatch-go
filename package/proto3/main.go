package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"packie/go-reborn/package/proto3/proto_obj"
)

func main() {
	stu1 := &proto_obj.Student{Name: "Lily", Male: true, Scores: []int32{90, 80, 100}}
	data, _ := proto.Marshal(stu1)

	stu2 := &proto_obj.Student{}
	_ = proto.Unmarshal(data, stu2)

	if stu1.Name == stu2.Name {
		fmt.Println("Same data")
	}
}
