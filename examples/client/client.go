package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	hub "github.com/Mihalic2040/Hub"
	"github.com/Mihalic2040/Hub/src/proto/api"
	"github.com/Mihalic2040/Hub/src/request"
	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
)

func MyHandler(input *api.Request) (response api.Response, err error) {
	// Do some processing with the input data
	// ...

	// Return the output data and no error
	return server.Response(input.Payload, 200), nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, app.Host.Peerstore().Peers())
}

var app hub.App

func main() {
	// ctx := context.Background()
	// //fake config
	config := types.Config{
		Host: "0.0.0.0",
		Port: "0",
		//Secret:     "SERVER",
		Rendezvous: "Hub",
		DHTServer:  true,
		ProtocolId: "/hub/0.0.1",
		Bootstrap:  "/ip4/0.0.0.0/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE",
	}

	//app := &hub.App{}

	app.Settings(config)

	//app.Handler(utils.GetFunctionName(MyHandler), MyHandler)

	app.Start(false)

	go func() {
		time.Sleep(time.Second * 5)
		peer := "12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4"

		data := api.Request{
			Payload: "gg",
			Handler: "MyHandler",
		}

		res, _ := request.New(app, peer, &data)

		log.Println(res)
	}()

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
	// for {
	// 	log.Println(app.Host.Peerstore().Peers())
	// }

}
