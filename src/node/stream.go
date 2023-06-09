package node

import (
	"bufio"
	"fmt"
	"log"

	"github.com/Mihalic2040/Hub/src/proto/api"
	"github.com/Mihalic2040/Hub/src/server"
	"github.com/libp2p/go-libp2p/core/network"
	"google.golang.org/protobuf/proto"
)

func stream_handler(stream network.Stream, handlers server.HandlerMap) {
	log.Println("New stream!!")
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create a data channel to receive the information
	dataCh := make(chan []byte)

	// Goroutine to read information from the stream and send it to the data channel
	go func() {
		for {
			// Read a chunk of data from the stream
			data := make([]byte, 1024) // Adjust the buffer size as per your needs
			n, err := rw.Read(data)
			if err != nil {
				//log.Println("Error reading data from stream:", err)
				close(dataCh)
				break
			}

			// Send the received data to the data channel
			dataCh <- data[:n]
		}
	}()

	// Read the information from the data channel
	data := <-dataCh

	// Create a new instance of your protobuf message
	req := &api.Request{}

	// Decode the protobuf data
	if err := proto.Unmarshal(data, req); err != nil {
		log.Println("Error decoding protobuf data:", err)
		return
	}

	// Start data processing thread

	log.Println(req)

	// Send the response back to the client
	response := api.Response{
		Payload: "Responklgdfmklgkldfgmdfse",
	}

	// Send the response to the stream
	response_b, err := proto.Marshal(&response)
	if err != nil {
		log.Println("Error encoding protobuf data:", err)

	}

	// Write the request bytes to the stream
	if _, err := rw.Write(response_b); err != nil {
		fmt.Println("Error writing protobuf response:", err)

	}

	// Flush the writer to ensure the data is sent
	if err := rw.Flush(); err != nil {
		log.Println("Error flush cahanel:", err)

	}

	// Close the stream
	stream.Close()

}

// func stream_handler(stream network.Stream, handlers server.HandlerMap) {
// 	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

// 	// Read data from the stream
// 	data, err := ioutil.ReadAll(rw)
// 	if err != nil {
// 		log.Println("Error reading data from stream:", err)
// 		return
// 	}

// 	// Create a new instance of your protobuf message
// 	req := &api.Request{}

// 	// Decode the protobuf data
// 	if err := proto.Unmarshal(data, req); err != nil {
// 		log.Println("Error decoding protobuf data:", err)
// 		return
// 	}

// 	log.Println(req)
// 	stream.Close()
// }
