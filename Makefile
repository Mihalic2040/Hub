compile:
	go build -o build/hub
	./build/hub

lib:
	go build -buildmode=c-shared -o library.so main.go
