# server
## Protocol Buffer Commands
Generate template code for serializing protocol buffer message and operating grpc service
- `$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/image_service.proto`