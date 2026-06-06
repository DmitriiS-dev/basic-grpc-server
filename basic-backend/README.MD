## Structure:

cmd/project        -> app entrypoint
internal/service   -> business logic
internal/store     -> database/storage layer
proto              -> protobuf definitions


Example to Install protobuf plugins:
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest



#### CheatSheet Of Commands:
- installing: `go get ...`
- manage dependencies: `go mod tidy`

