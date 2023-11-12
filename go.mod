module github.com/kunle001/go-grpc

go 1.21.4

require (
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
)

// You can group indirect dependencies together in a single 'require' block
require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
)
