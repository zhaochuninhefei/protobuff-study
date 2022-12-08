package main

import (
	"fmt"
	"github.com/zhaochuninhefei/myproto-go/owner"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("myproto-test start...")

	// 创建pb消息
	owner1 := &owner.Owner{
		OwnerId:   1,
		OwnerName: "owner1",
		OwnerDesc: "just test",
	}
	fmt.Printf("owner1: %s\n", owner1)

	// 序列化pb消息
	outOwner1File := "testdata/owner1.pb"
	outOwner1, err := proto.Marshal(owner1)
	if err != nil {
		log.Fatalln("Failed to encode owner1:", err)
	}
	if err := ioutil.WriteFile(outOwner1File, outOwner1, 0644); err != nil {
		log.Fatalln("Failed to write owner1:", err)
	}

	// 反序列化pb消息
	inOwner1, err := ioutil.ReadFile(outOwner1File)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	owner1New := &owner.Owner{}
	if err := proto.Unmarshal(inOwner1, owner1New); err != nil {
		log.Fatalln("Failed to parse owner1:", err)
	}
	fmt.Printf("owner1New: %s\n", owner1New)

}
