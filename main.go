package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Mihalic2040/Hub/src/node"
	"github.com/Mihalic2040/Hub/src/proto/api"
	"github.com/Mihalic2040/Hub/src/request"
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

	peers := host_one.Host.Peerstore().Peers()
	peers.String()

	fmt.Fprintf(w, peers.String())
}

var (
	host_one types.Host
	host_two types.Host
)

func main() {
	ctx := context.Background()
	//fake config
	config := types.Config{
		Host:       "0.0.0.0",
		Port:       "0",
		Secret:     "MIHALIC2040",
		Rendezvous: "rendezvous",
		ProtocolId: "/hub/0.0.1",
		Bootstrap:  "/ip4/0.0.0.0/udp/6666/quic/p2p/12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4",
	}

	config_two := types.Config{
		Host:       "0.0.0.0",
		Port:       "0",
		Secret:     "VICTOR",
		Rendezvous: "rendezvous",
		ProtocolId: "/hub/0.0.1",
		Bootstrap:  "/ip4/0.0.0.0/udp/6666/quic/p2p/12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4",
	}

	// runing server
	handlers := server.HandlerMap{
		utils.GetFunctionName(MyHandler): MyHandler,
	}

	host_one = node.Server(ctx, handlers, config, false)

	host_two = node.Server(ctx, handlers, config_two, false)

	go func() {

		for {
			time.Sleep(time.Second * 1)
			peer_id := "12D3KooWCjZ7VQMu1jtJvisqpUcwqZUcUwJnikPbxqMijALZShCP"
			data := api.Request{
				Payload: "Hello",
				Handler: "MyHandler",
			}
			res, err := request.New(host_one, peer_id, &data)
			if err != nil {
				log.Println(err)
			} else {
				log.Println(res)
			}

		}
	}()

	// go func() {

	// 	for {
	// 		time.Sleep(time.Second * 1)
	// 		node.Rendezvous(ctx, host_one.Host, host_one.Dht, config)
	// 	}
	// }()

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
