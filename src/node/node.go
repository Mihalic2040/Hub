package node

import (
	"fmt"
	"log"

	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/multiformats/go-multiaddr"
)

type Config struct {
	Host       string
	Port       string
	ProtocolId string
	Bootstrap  string
}

func Start_host(Config Config, handlers server.HandlerMap, input types.InputData) host.Host {

	// 0.0.0.0 will listen on any interface device.

	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s", Config.Host, Config.Port))

	host, _ := libp2p.New(
		libp2p.ListenAddrs(sourceMultiAddr),
	)

	log.Printf("\n[*] Your Multiaddress Is: /ip4/%s/tcp/%v/p2p/%s\n", Config.Host, Config.Port, host.ID().Pretty())
	// Set a function as stream handler.
	// This function is called when a peer initiates a connection and starts a stream with this peer.
	//host.SetStreamHandler(protocol.ID(Config.ProtocolId), stream_handler)

	return host
}
