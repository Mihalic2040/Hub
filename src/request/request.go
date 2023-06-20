package request

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"

	hub "github.com/Mihalic2040/Hub"
	"github.com/Mihalic2040/Hub/src/proto/api"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"google.golang.org/protobuf/proto"
)

func New(host hub.App, peerID string, data *api.Request) (*api.Response, error) {
	// Find a peer by its ID
	targetPeerID, err := peer.Decode(peerID)
	if err != nil {
		return nil, fmt.Errorf("Invalid peer ID: %v", err)
	}

	peerInfo, err := host.Dht.FindPeer(context.Background(), targetPeerID)
	if err != nil {
		return nil, fmt.Errorf("Fail to find peer: %v", err)
	}

	// Create a stream to the peer
	stream, err := host.Host.NewStream(context.Background(), peerInfo.ID, protocol.ID(host.Config.ProtocolId))
	if err != nil {
		return nil, fmt.Errorf("Failed to create stream: %v", err)
	}

	// Create a bufio ReadWriter using the stream
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create a request

	// Serialize the request to bytes
	bytes, err := proto.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to serialize request: %v", err)
	}

	// Write the request bytes to the stream
	if _, err := rw.Write(bytes); err != nil {
		return nil, fmt.Errorf("Failed to send request: %v", err)
	}

	// Flush the writer to ensure the data is sent
	if err := rw.Flush(); err != nil {
		return nil, fmt.Errorf("Error flusing writer: %v", err)
	}

	// Read the response from the stream
	responseBytes, err := ioutil.ReadAll(rw)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	// Create a response message to unmarshal the response bytes
	response := &api.Response{}

	// Unmarshal the response bytes
	if err := proto.Unmarshal(responseBytes, response); err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	//log.Println(response)
	// // Close the stream only if a response is received
	stream.Close()

	// Use the response message as needed
	// ...
	return response, nil

}
