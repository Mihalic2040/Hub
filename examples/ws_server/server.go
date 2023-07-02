package wserver

import "C"
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	hub "github.com/Mihalic2040/Hub"
	"github.com/Mihalic2040/Hub/src/proto/api"
	"github.com/Mihalic2040/Hub/src/request"
	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/Mihalic2040/Hub/src/utils"
	"github.com/gorilla/websocket"
)

func MyHandler(input *api.Request) (response api.Response, err error) {
	// Do some processing with the input data
	// ...

	// Return the output data and no error
	return server.Response(input.Payload, 200), nil
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow connections from any origin
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		// Read incoming messages from the WebSocket client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var req api.Request
		err = json.Unmarshal(message, &req)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			break
		}
		res, err := request.New(app, req.User, &req)
		json_data, _ := json.Marshal(res)
		err = conn.WriteMessage(websocket.TextMessage, json_data)
		if err != nil {
			log.Println("Error sending response:", err)
			break
		}
	}
}

var app hub.App

//export start
func start() {
	// ctx := context.Background()
	// //fake config
	config := types.Config{
		Host:       "0.0.0.0",
		Port:       "0",
		Secret:     "SERVER_WS",
		Rendezvous: "Hub",
		DHTServer:  true,
		ProtocolId: "/hub/0.0.1",
		//Bootstrap:  "/ip4/0.0.0.0/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE",
	}

	app.Settings(config)

	app.Handler(utils.GetFunctionName(MyHandler), MyHandler)

	app.Start(false)

	http.HandleFunc("/ws", wsHandler)

	fmt.Println("WebSocket server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// for {
	// 	log.Println(app.Host.Peerstore().Peers())
	// }

}

// func main() {
// 	// ctx := context.Background()
// 	// //fake config
// 	config := types.Config{
// 		Host:       "0.0.0.0",
// 		Port:       "0",
// 		Secret:     "SERVER_WS",
// 		Rendezvous: "Hub",
// 		DHTServer:  true,
// 		ProtocolId: "/hub/0.0.1",
// 		//Bootstrap:  "/ip4/0.0.0.0/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE",
// 	}

// 	app.Settings(config)

// 	app.Handler(utils.GetFunctionName(MyHandler), MyHandler)

// 	app.Start(false)

// 	http.HandleFunc("/ws", wsHandler)

// 	fmt.Println("WebSocket server listening on port 8080...")
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// 	// for {
// 	// 	log.Println(app.Host.Peerstore().Peers())
// 	// }

// }
