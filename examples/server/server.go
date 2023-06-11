package main

import (
	"context"

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
		Host:       "0.0.0.0",
		Port:       "6666",
		Secret:     "SERVER",
		Rendezvous: "Hub",
		DHTServer:  true,
		ProtocolId: "/hub/0.0.1",
		Bootstrap:  "/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4",
	}

	// runing server
	handlers := server.HandlerMap{
		utils.GetFunctionName(MyHandler): MyHandler,
	}
	node.Server(ctx, handlers, config, true)

}
