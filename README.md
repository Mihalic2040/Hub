# Hub

P2P local server for building p2p apps.

## Im building home lab and i need fast framework for creating datatransfer for my apps. So i create my own framework for it.

# Example
    package main

    import (
        "github.com/Mihalic2040/Hub/src/node"
        "github.com/Mihalic2040/Hub/src/server"
        "github.com/Mihalic2040/Hub/src/types"
        "github.com/Mihalic2040/Hub/src/utils"
    )

    func MyHandler(input interface{}) (output interface{}, err error) {
        // Do some processing with the input data
        // ...

        // Return the output data and no error
        return input, nil
    }

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
            HandlerName: "MyHandler",
            Input:       50,
        }

        // runing server
        handlers := server.HandlerMap{
            utils.GetFunctionName(MyHandler): MyHandler,
        }

        node.Server(handlers, input, config)
    }




    
