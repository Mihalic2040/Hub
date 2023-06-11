package main

import (
	"context"
	"fmt"
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

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Get the peer ID from the request
	// curl http://localhost:8080/ -X POST -d "peer_id="

	peers := host.Host.Peerstore().Peers()
	peers.String()

	fmt.Fprintf(w, peers.String())
}

var (
	host types.Host
)

func main() {
	ctx := context.Background()
	//fake config
	config := types.Config{
		Host:       "0.0.0.0",
		Port:       "0",
		Secret:     "MIHALIC2040",
		Rendezvous: "Hub",
		ProtocolId: "/hub/0.0.1",
		Bootstrap:  "/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWCjZ7VQMu1jtJvisqpUcwqZUcUwJnikPbxqMijALZShCP",
	}

	// runing server
	handlers := server.HandlerMap{
		utils.GetFunctionName(MyHandler): MyHandler,
	}

	host = node.Server(ctx, handlers, config, true)

	// go func() {

	// 	for {
	// 		peer_id := "12D3KooWDGDG3iwx75D8qPdpW24EDB1ZqTQBuiFtPaEhU3azUXmn"
	// 		host.Dht.RefreshRoutingTable()
	// 		data := api.Request{
	// 			Payload: "Hello",
	// 			Handler: "MyHandler",
	// 		}
	// 		res, _ := request.New(host, peer_id, &data)

	// 		log.Println(res)
	// 	}
	// }()

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
