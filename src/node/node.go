package node

import (
	"context"
	"fmt"
	"log"

	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
)

type Config struct {
	Host             string
	Port             string
	RendezvousString string
	ProtocolId       string
	Bootstrap        string
}

func Start_host(Config Config, handlers server.HandlerMap, input types.InputData) host.Host {
	ctx := context.Background()

	// 0.0.0.0 will listen on any interface device.

	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s", Config.Host, Config.Port))

	host, _ := libp2p.New(
		libp2p.ListenAddrs(sourceMultiAddr),
	)

	log.Printf("[*] Your Multiaddress Is: /ip4/%s/tcp/%v/p2p/%s\n", Config.Host, Config.Port, host.ID().Pretty())

	// Set a function as stream handler.
	// This function is called when a peer initiates a connection and starts a stream with this peer.
	host.SetStreamHandler(protocol.ID(Config.ProtocolId), stream_handler)

	//Init KDHT
	kademliaDHT := init_DHT(ctx, host)
	kademliaDHT.Context()
	// boot from config
	boot(ctx, Config, host)

	// Stating mdns service and bootstraping peers
	start_mdns(host, Config, ctx)

	return host
}
