package main

import (
	"github.com/Mihalic2040/Hub/src/node"
	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/tests"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/Mihalic2040/Hub/src/utils"
)

func main() {
	//fake config
	config := types.Config{
		Host:             "0.0.0.0",
		Port:             "4344",
		RendezvousString: "Hub",
		ProtocolId:       "/hub/0.0.1",
		Bootstrap:        "/ip4/0.0.0.0/tcp/4344/p2p/12D3KooWMQB9RxQHng8ALnaKcLNKgCcrAMRRYtCr2mGrfUKTmBES",
	}

	//fake input
	input := types.InputData{
		HandlerName: "piska",
		Input:       50,
	}

	// runing server
	handlers := server.HandlerMap{
		utils.GetFunctionName(tests.MyHandlerFromApp): tests.MyHandlerFromApp,
		"piska": tests.MyHandlerFromAppp,
	}

	node.Server(handlers, input, config)
}
