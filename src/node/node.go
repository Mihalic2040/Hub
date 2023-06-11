package node

import (
	"context"
	"fmt"
	"log"

	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/Mihalic2040/Hub/src/utils"
	"github.com/fatih/color"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/muxer/mplex"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/multiformats/go-multiaddr"
)

func Server(ctx context.Context, handlers server.HandlerMap, Config types.Config, serve bool) types.Host {
	host := Start_host(ctx, Config, handlers, serve)
	//fmt.Println(host)
	//server.Thread(handlers, input)
	return host
}

func Start_host(ctx context.Context, Config types.Config, handlers server.HandlerMap, serve bool) types.Host {

	// 0.0.0.0 will listen on any interface device.

	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s/", Config.Host, Config.Port))
	sourceMultiAddrQuic, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/udp/%s/quic", Config.Host, Config.Port))

	// TEST

	prvKey, err := utils.GeneratePrivateKeyFromString(Config.Secret)
	if err != nil {
		log.Fatal(err)
	}

	taranspors := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(quic.NewTransport),
	)

	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/mplex/", mplex.DefaultTransport),
	)

	addrs := libp2p.ListenAddrStrings(
		sourceMultiAddr.String(),
		sourceMultiAddrQuic.String(),
	)

	// CREATE HOST
	//var kademliaDHT *dht.IpfsDHT
	host, _ := libp2p.New(
		addrs,
		libp2p.Identity(prvKey),
		taranspors,
		muxers,

		//Cool stuff'
		libp2p.EnableHolePunching(),
		libp2p.NATPortMap(),
		libp2p.EnableNATService(),
		libp2p.EnableRelayService(),
	)

	//log.Printf("[*] Your Multiaddress Is: /ip4/%s/tcp/%v/p2p/%s\n", Config.Host, Config.Port, host.ID().Pretty())

	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("[*] Your ID is: %s", green(host.ID().Pretty()))
	log.Println("[*] Your Multiaddress is:", host.Addrs())

	// Set a function as stream handler.
	// This function is called when a peer initiates a connection and starts a stream with this peer.
	host.SetStreamHandler(protocol.ID(Config.ProtocolId), func(stream network.Stream) {
		//log.Println("DEBUG: new steam in handler")
		//log.Println(handlers)
		stream_handler(stream, handlers)
	})
	//Init KDHT
	kademliaDHT := init_DHT(ctx, host, Config)
	// boot from config
	bootstrap(ctx, kademliaDHT)
	boot(ctx, Config, host)
	//Rendezvous(ctx, host, kademliaDHT, Config)

	// Stating mdns service and bootstraping peers
	if serve == true {
		start_mdns(host, Config, ctx)
	} else {
		go start_mdns(host, Config, ctx)
	}

	server := types.Host{
		Host:   host,
		Dht:    kademliaDHT,
		Config: Config,
	}

	return server
}

// for {
// 	// Find a peer by its ID
// 	targetPeerID, err := peer.Decode("12D3KooWQXMKJFm6f3pWNHxHE8z7KqRaaDevnFKqyGahXQxj1CVN")
// 	if err != nil {
// 		fmt.Println("Invalid peer ID:", err)
// 		return
// 	}

// 	peerInfo, err := kademliaDHT.FindPeer(context.Background(), targetPeerID)
// 	if err != nil {
// 		fmt.Println("Failed to find peer:", err)
// 	}

// 	// Create a stream to the peer
// 	stream, err := host.NewStream(context.Background(), peerInfo.ID, protocol.ID(Config.ProtocolId))
// 	if err == nil {
// 		stream.Close()
// 	}

// 	// Use the stream for communication
// 	// ...

// 	// Remember to close the stream when done

// 	time.Sleep(2 * time.Second)
// }
