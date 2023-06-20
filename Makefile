client:
	go build -o build/client examples/client/client.go
	./build/client

lib:
	go build -buildmode=c-shared -o library.so main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative src/proto/api/protocol.proto 

server:
	go build -o build/server examples/server/server.go
	./build/server

