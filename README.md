# Hub

P2P local server for building p2p apps.

## Im building home lab and i need fast framework for creating datatransfer for my apps. So i create my own framework for it.

# Example
    package main

    import (
        "context"
        "fmt"
        "log"
        "net/http"

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
        peerID := r.FormValue("peer_id")
        if peerID == "" {
            http.Error(w, "Missing peer ID", http.StatusBadRequest)
            return
        }

        // create request
        data := api.Request{
            User:    "Hello",
            Payload: "Hello payload",
            Handler: "MyHandler",
        }

        response, err := request.New(host, peerID, &data)
        if err != nil {
            log.Println("Error creating request: ", err)
        }
        //fmt.Println(response)

        fmt.Fprintf(w, response.Payload)
    }

    var (
        host types.Host
    )

    func main() {
        ctx := context.Background()
        //fake config
        config := types.Config{
            Host:             "0.0.0.0",
            Port:             "0",
            RendezvousString: "Hub",
            ProtocolId:       "/hub/0.0.1",
            Bootstrap:        "/ip4/0.0.0.0/tcp/4344/p2p/12D3KooWMQB9RxQHng8ALnaKcLNKgCcrAMRRYtCr2mGrfUKTmBES",
        }

        // runing server
        handlers := server.HandlerMap{
            utils.GetFunctionName(MyHandler): MyHandler,
        }

        host = node.Server(ctx, handlers, config, false)

        http.HandleFunc("/", handleRequest)
        fmt.Println("Server started on http://localhost:8080")
        http.ListenAndServe(":8080", nil)
    }






    
