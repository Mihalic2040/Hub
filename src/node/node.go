package node

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
)

func Server(handlers server.HandlerMap, input types.InputData, Config types.Config) {
	Start_host(Config, handlers, input)
	//fmt.Println(host)
	//server.Thread(handlers, input)
}

func printPeers(peer peer.AddrInfo) {
	fmt.Println("Connected peers:")
	fmt.Println(peer.ID)
	for _, addr := range peer.Addrs {
		fmt.Println("  ", addr)
	}
}

func findAllPeers(dhtInstance *dht.IpfsDHT) (peer.AddrInfo, error) {
	// Find all peers in the DHT
	peers, err := dhtInstance.FindPeer(context.Background(), "")
	if err != nil {
		log.Println("fff")
	}

	return peers, nil
}

func Start_host(Config types.Config, handlers server.HandlerMap, input types.InputData) {
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
	go start_mdns(host, Config, ctx)

	for {
		// Find a peer by its ID
		targetPeerID, err := peer.Decode("12D3KooWQXMKJFm6f3pWNHxHE8z7KqRaaDevnFKqyGahXQxj1CVN")
		if err != nil {
			fmt.Println("Invalid peer ID:", err)
			return
		}

		peerInfo, err := kademliaDHT.FindPeer(context.Background(), targetPeerID)
		if err != nil {
			fmt.Println("Failed to find peer:", err)
		}

		// Create a stream to the peer
		stream, err := host.NewStream(context.Background(), peerInfo.ID, protocol.ID(Config.ProtocolId))
		if err == nil {
			stream.Close()
		}

		// Use the stream for communication
		// ...

		// Remember to close the stream when done

		time.Sleep(2 * time.Second)
	}

}
