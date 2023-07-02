client:
	go build -o build/client examples/client/client.go
	./build/client

lib:
	go build -buildmode=c-shared -o build/library.so examples/server/server.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative src/proto/api/protocol.proto 

server:
	go build -o build/server examples/server/server.go
	./build/server

ws:
	go build -buildmode=c-shared -o examples/ws_server/library.so examples/ws_server/server.go
	
