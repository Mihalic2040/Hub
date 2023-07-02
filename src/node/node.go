package node

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"strconv"

	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/Mihalic2040/Hub/src/utils"
	color "github.com/fatih/color"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/muxer/mplex"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/libp2p/go-libp2p/p2p/transport/websocket"
	"github.com/multiformats/go-multiaddr"
)

func Start_host(ctx context.Context, config types.Config, handlers server.HandlerMap, serve bool) types.App {

	// 0.0.0.0 will listen on any interface device.

	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s/", config.Host, config.Port))
	sourceMultiAddrQuic, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/udp/%s/quic", config.Host, config.Port))
	var sourceMultiAddrWs multiaddr.Multiaddr
	if config.Port != "0" {
		port, _ := strconv.Atoi(config.Port)
		portNumber := strconv.Itoa(port + 1)
		sourceMultiAddrWs, _ = multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s/ws", config.Host, portNumber))
	} else {
		sourceMultiAddrWs, _ = multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s/ws", config.Host, config.Port))
	}

	// TEST
	var prvKey crypto.PrivKey
	var err error
	if config.Secret == "" {
		prvKey, _, err = crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
		if err != nil {
			log.Fatalln("[*] Error generating key pair:", err)
		}
	} else {
		prvKey, err = utils.GeneratePrivateKeyFromString(config.Secret)
		if err != nil {
			log.Fatalln("[*] Error generating key pair:", err)
		}
	}

	taranspors := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(quic.NewTransport),

		//For js nodes
		libp2p.Transport(websocket.New),
	)

	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/mplex/", mplex.DefaultTransport),

		//For js nodes
		libp2p.Muxer("/yamux/", yamux.DefaultTransport),
	)

	addrs := libp2p.ListenAddrStrings(
		sourceMultiAddr.String(),
		sourceMultiAddrQuic.String(),
		sourceMultiAddrWs.String(),
	)

	// CREATE HOST
	host, _ := libp2p.New(
		addrs,
		libp2p.Identity(prvKey),
		taranspors,
		muxers,

		//Cool stuff
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
	host.SetStreamHandler(protocol.ID(config.ProtocolId), func(stream network.Stream) {
		//log.Println("DEBUG: new steam in handler")
		//log.Println(handlers)
		stream_handler(stream, handlers)
	})

	//WORK IN PROGRESS
	// Data stream proto for streaming big amount of data
	host.SetStreamHandler(protocol.ID(config.ProtocolId+"/stream/"), func(stream network.Stream) {
		data_stream(stream, handlers)
	})

	//Init KDHT
	kademliaDHT := init_DHT(ctx, host, config)
	// boot from config
	bootstrap(ctx, kademliaDHT)
	boot(ctx, config, host)
	//Rendezvous(ctx, host, kademliaDHT, Config)

	// Stating mdns service and bootstraping peers
	if serve == true {
		start_mdns(host, config, ctx)
	} else {
		go start_mdns(host, config, ctx)
	}

	server := types.App{
		Host:   host,
		Dht:    kademliaDHT,
		Config: config,
	}

	return server
}
