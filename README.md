
# Hub

Hello, my name is Mykhailo, one evening I decided to make a home ecosystem. And for that I needed to create a network between all the devices in my house. Therefore, I set out to make a small framework that will help me do this easily. Since I really like p2p technologies, I want to build a p2p network.



## Documentation

I've created a custom handler system that runs on top of libp2p:Stream. I used protobuf for data transport.

#### Request



| Parameter | Type     | Description                                     |
| :-------- | :------- | :-------------------------                      |
| `User`    | `string` | **Required**. Sender ID                         |
| `Payload` | `string` | **Required**. Data (I think i will use json)    |
| `Handler` | `string` | **Required**. Handler name                      |

#### Response



| Parameter | Type     | Description                            |
| :-------- | :------- | :--------------------------------      |
| `Payload` | `string` | **Required**. Id of item to fetch      |
| `Status` | `int64` | **Required**. Status code "Like http :з" |

#### Example

Server
```go
package main

import (
	hub "github.com/Mihalic2040/Hub"
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
	// ctx := context.Background()
	// //fake config
	config := types.Config{
		Host: "0.0.0.0",
		Port: "0",
		//Secret:     "SERVER",
		Rendezvous: "Hub",
		DHTServer:  true,
		ProtocolId: "/hub/0.0.1",
		Bootstrap:  "/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE",
	}

	app := &hub.App{}

	app.Settings(config)

	app.Handler(utils.GetFunctionName(MyHandler), MyHandler)

	app.Start(true)


}


```
Request
```go
package main

import (
	"fmt"
	"net/http"

	hub "github.com/Mihalic2040/Hub"
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
		Bootstrap:  "/ip4/141.145.193.111/tcp/6666/p2p/12D3KooWQd1K1k8XA9xVEzSAu7HUCodC7LJB6uW5Kw4VwkRdstPE",
	}

	//app := &hub.App{}

	app.Settings(config)

	app.Handler(utils.GetFunctionName(MyHandler), MyHandler)

	app.Start(false)

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)

}
```
## ToDo

- [ ] Encryption protobuf
- [ ] Relays
- [ ] Cool logo
## Feedback

Discord: Mihalic2040#6533


