package reflect

import (
	"fmt"
	protoOld "github.com/golang/protobuf/proto"
	"github.com/zhaochuninhefei/myproto-go/owner"
	"google.golang.org/protobuf/proto"
	"reflect"
	"strings"
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

	fmt.Println("直接反射...")
	pb := proto.Message(owner1)

	rt := reflect.TypeOf(pb)
	mrt := rt.Elem()

	fmt.Printf("rt.Kind(): %s\n", rt.Kind())
	fmt.Printf("rt.Name(): %s\n", rt.Name())
	fmt.Printf("mrt.Kind(): %s\n", mrt.Kind())
	fmt.Printf("mrt.Name(): %s\n", mrt.Name())

	if mrt.Kind() != reflect.Struct {
		fmt.Println("owner1 is not a struct!")
		return
	}
	for i := 0; i < mrt.NumField(); i++ {
		f := mrt.Field(i)
		fmt.Printf("字段名: %s, 字段类型: %s, 字段标签: %s\n", f.Name, f.Type.Name(), f.Tag)
		tagProto := f.Tag.Get("protobuf")
		if tagProto != "" {
			fmt.Printf("protobuf: %s\n", tagProto)
			tmp := strings.Split(tagProto, ",")
			for _, word := range tmp {
				if strings.HasPrefix(word, "name=") {
					fmt.Println("对应proto消息的字段名:" + strings.TrimPrefix(word, "name="))
					break
				}
			}
		}

	}
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
	//goland:noinspection GoDeprecation
	protoProps := protoOld.GetProperties(mVal.Type())
	fmt.Println("----- oldReflect:")
	for _, prop := range protoProps.Prop {
		fieldName := prop.OrigName
		fieldValue := mVal.FieldByName(prop.Name)
		fieldTypeStruct, ok := mVal.Type().FieldByName(prop.Name)
		if !ok {
			fmt.Printf("programming error: proto does not have field advertised by proto package : %s\n", prop.Name)
			continue
		}
		fieldType := fieldTypeStruct.Type
		fmt.Printf("prop.Name: %s, fieldName: %s, fieldValue: %s, fieldType: %s \n", prop.Name, fieldName, fieldValue, fieldType)
	}
}

func newReflect(msg proto.Message) {
	pmVal := reflect.ValueOf(msg)
	pmType := reflect.TypeOf(msg)
	if pmVal.Kind() != reflect.Ptr {
		fmt.Println("error in oldReflect 1")
		return
	}
	if pmVal.IsNil() {
		fmt.Println("error in oldReflect 2")
		return
	}
	mVal := pmVal.Elem()
	mType := pmType.Elem()
	if mVal.Kind() != reflect.Struct {
		fmt.Println("error in oldReflect 3")
		return
	}

	fdNum := mType.NumField()
	fmt.Printf("fdNum: %d\n", fdNum)

	for i := 0; i < fdNum; i++ {
		f := mType.Field(i)
		fmt.Printf("go字段名: %s\n", f.Name)
	}

	m := msg.ProtoReflect()
	fds := m.Descriptor().Fields()
	fmt.Printf("fds.Len(): %d\n", fds.Len())

	fmt.Println("----- newReflect:")
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)
		//fmt.Println(fd.TextName())
		fv := m.Get(fd)
		//fmt.Println(fv)
		fieldName := fd.Name()
		fieldType := fd.Kind()
		fmt.Printf("fieldName: %s, fieldValue: %s, fieldType: %s \n", fieldName, fv, fieldType)
	}

}
