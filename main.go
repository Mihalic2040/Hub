package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Mihalic2040/Hub/src/node"
	"github.com/Mihalic2040/Hub/src/proto/api"
	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/Mihalic2040/Hub/src/utils"
)

func MyHandler(input *api.Request) (response api.Response, err error) {
	// Do some processing with the input data
	// ...

	// Return the output data and no error
	return server.Response(input.Payload, 200), nil
}

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	// Get the peer ID from the request
// 	// curl http://localhost:8080/ -X POST -d "peer_id="

// 	peers := host_one.Host.Peerstore().Peers()

// 	fmt.Fprintf(w, peers.String())
// }

// var (
// 	host_one types.Host
// )

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, app.Host.Network().Peers())
}

var app types.Host

func main() {
	ctx := context.Background()
	//fake config
	config := types.Config{
		Host: "0.0.0.0",
		Port: "0",
		// Secret:     "MIHALIC2040", // If Secret is "" genereting random Prvkey!!
		Rendezvous: "Hub",
		ProtocolId: "/hub/0.0.1",
		Bootstrap:  "/ip4/141.145.193.111/udp/6666/quic/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE",
	}

	// runing server
	handlers := server.HandlerMap{
		utils.GetFunctionName(MyHandler): MyHandler,
	}

	app = node.Server(ctx, handlers, config, false)

	log.Println("HTTP server started!!")
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
