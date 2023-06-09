compile:
	go build -o build/hub
	./build/hub

lib:
	go build -buildmode=c-shared -o library.so main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative src/proto/api/protocol.proto 
