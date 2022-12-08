package reflect

import (
	"fmt"
	protoOld "github.com/golang/protobuf/proto"
	"github.com/zhaochuninhefei/myproto-go/owner"
	"google.golang.org/protobuf/proto"
	"reflect"
)

func doReflect() {
	owner1 := &owner.Owner{
		OwnerId:   1,
		OwnerName: "owner1",
		OwnerDesc: "just test",
	}
	fmt.Printf("owner1: %s\n", owner1)

	oldReflect(owner1)
	newReflect(owner1)
}

func oldReflect(msg protoOld.Message) {
	pmVal := reflect.ValueOf(msg)
	if pmVal.Kind() != reflect.Ptr {
		fmt.Println("error in oldReflect 1")
		return
	}
	if pmVal.IsNil() {
		fmt.Println("error in oldReflect 2")
		return
	}
	mVal := pmVal.Elem()
	if mVal.Kind() != reflect.Struct {
		fmt.Println("error in oldReflect 3")
		return
	}
	protoProps := protoOld.GetProperties(mVal.Type())
	fmt.Println("----- oldReflect:")
	for _, prop := range protoProps.Prop {
		fmt.Println(prop)
	}
}

func newReflect(msg proto.Message) {
	m := msg.ProtoReflect()
	fds := m.Descriptor().Fields()
	fmt.Println("----- newReflect:")
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)
		fmt.Println(fd)
	}
}
