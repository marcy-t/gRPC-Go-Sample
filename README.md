# gRPC Sample
- Client
  - go run cmd/client/main.go
- Server
  - go run cmd/server/main.go

# command
- protoc -I ./pkg/proto ./proto/helloworld.proto --go_out=plugins=grpc:./pkg/proto
