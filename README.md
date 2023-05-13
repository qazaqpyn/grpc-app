# Simple gRPC client and server

### RUN
1) ```go run cmd/server/main.go```
2) ```go run cmd/client/main.go```

### Generation of .pg.go files
```protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/notification.proto```