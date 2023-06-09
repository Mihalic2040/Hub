package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mihalic2040/Hub/src/node"
	"github.com/Mihalic2040/Hub/src/proto/api"
	"github.com/Mihalic2040/Hub/src/request"
	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/Mihalic2040/Hub/src/utils"
)

func MyHandler(input interface{}) (response api.Response, err error) {
	// Do some processing with the input data
	// ...

	// Return the output data and no error
	return server.Response("Hello from handler", 200), nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Get the peer ID from the request
	// curl http://localhost:8080/ -X POST -d "peer_id="
	peerID := r.FormValue("peer_id")
	if peerID == "" {
		http.Error(w, "Missing peer ID", http.StatusBadRequest)
		return
	}

	response, err := request.New(host, peerID)
	if err != nil {
		log.Println("Error creating request: ", err)
	}
	//fmt.Println(response)

	fmt.Fprintf(w, response.Payload)
}

var (
	host types.Host
)

func main() {
	//fake config
	config := types.Config{
		Host:             "0.0.0.0",
		Port:             "0",
		RendezvousString: "Hub",
		ProtocolId:       "/hub/0.0.1",
		Bootstrap:        "/ip4/0.0.0.0/tcp/4344/p2p/12D3KooWMQB9RxQHng8ALnaKcLNKgCcrAMRRYtCr2mGrfUKTmBES",
	}

	// runing server
	handlers := server.HandlerMap{
		utils.GetFunctionName(MyHandler): MyHandler,
	}

	host = node.Server(handlers, config, false)

	http.HandleFunc("/", handleRequest)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8081", nil)
}

// for {
// 	// Find a peer by its ID
// 	targetPeerID, err := peer.Decode("12D3KooWStR9RTNfhHrGrPwph3itjZZFxHEu6wV4ph2G4oFtHARw")
// 	if err != nil {
// 		fmt.Println("Invalid peer ID:", err)
// 		return
// 	}

// 	peerInfo, err := dht.FindPeer(context.Background(), targetPeerID)
// 	if err != nil {
// 		fmt.Println("Failed to find peer:", err)
// 	}

// 	// Create a stream to the peer
// 	stream, err := host.NewStream(context.Background(), peerInfo.ID, protocol.ID(config.ProtocolId))
// 	if err == nil {
// 		stream.Close()
// 	}

// 	// Use the stream for communication
// 	// ...

// 	// Remember to close the stream when done

// 	time.Sleep(2 * time.Second)
// }
