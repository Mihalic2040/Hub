package main

import (
	"github.com/Mihalic2040/Hub/src/node"
	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/tests"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/Mihalic2040/Hub/src/utils"
)

func Server(handlers server.HandlerMap, input types.InputData, Config node.Config) {
	host := node.Start_host(Config, handlers, input)

	host.Addrs()
	//fmt.Println(host)
	//server.Thread(handlers, input)
}

func main() {
	//fake config
	config := node.Config{
		Host:       "0.0.0.0",
		Port:       "4441",
		ProtocolId: "/hub/0.0.1",
		Bootstrap:  "0.0.0.0",
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

	Server(handlers, input, config)
}
