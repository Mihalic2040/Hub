
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
    "context"

    "github.com/Mihalic2040/Hub/src/node"
    "github.com/Mihalic2040/Hub/src/proto/api"
    "github.com/Mihalic2040/Hub/src/server"
    "github.com/Mihalic2040/Hub/src/types"
    "github.com/Mihalic2040/Hub/src/utils"
)

func Echo(input *api.Request) (response api.Response, err error) {
    // Do some processing with the input data
    // ...

    // Return the output data and no error
    return server.Response(input.Payload, 200), nil
}

func main() {
    ctx := context.Background()
    // Config
    config := types.Config{
        Host:       "0.0.0.0",
        Port:       "6666",
        Secret:     "SERVER",
        Rendezvous: "Hub",
        DHTServer:  true,
        ProtocolId: "/hub/0.0.1",
        Bootstrap:  "/ip4/0.0.0.0/tcp/6666/p2p/12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4",
    }

    // Runing server
    handlers := server.HandlerMap{
        utils.GetFunctionName(Echo): Echo,
    }

    node.Server(ctx, handlers, config, true)

}

```
Request
```go
func main() {
    ctx := context.Background()
    // Config
    config := types.Config{
        Host:       "0.0.0.0",
        Port:       "6666",
        Secret:     "SERVER",
        Rendezvous: "Hub",
        DHTServer:  true,
        ProtocolId: "/hub/0.0.1",
        Bootstrap:  "/ip4/0.0.0.0/tcp/6666/p2p/12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4",
    }

    // Runing server
    handlers := server.HandlerMap{
        utils.GetFunctionName(Echo): Echo,
    }

    app := node.Server(ctx, handlers, config, false)


    req := api.Request{
        Payload: "Bla bla bla",
        Handler: "Echo",
    }

    user_id := "12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4"

    res, err := request.new(app, user_id, req)
    if err != nil {
        //ERROR
    }
    
}
```
## ToDo

- [ ] Encryption protobuf
- [ ] Relays
- [ ] Cool logo
## Feedback

Discord: Mihalic2040#6533


