module github.com/zhaochuninhefei/myproto-test

go 1.17

require (
	github.com/golang/protobuf v1.5.2
	github.com/zhaochuninhefei/myproto-go v0.0.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/google/go-cmp v0.5.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace github.com/zhaochuninhefei/myproto-go => ../myproto-go
