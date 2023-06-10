package main

import (
	"context"
	"log"

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

func main() {
	ctx := context.Background()
	//fake config
	config := types.Config{
		Host:             "0.0.0.0",
		Port:             "0",
		Secret:           "SERVER",
		RendezvousString: "Hub",
		ProtocolId:       "/hub/0.0.1",
		Bootstrap:        "/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWCjZ7VQMu1jtJvisqpUcwqZUcUwJnikPbxqMijALZShCP",
	}

	// runing server
	handlers := server.HandlerMap{
		utils.GetFunctionName(MyHandler): MyHandler,
	}

	host := node.Server(ctx, handlers, config, true)

	for {
		log.Println(host.Dht.RoutingTable().ListPeers())
	}

}
