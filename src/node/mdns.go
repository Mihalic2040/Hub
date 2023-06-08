package node

import (
	"bufio"
	"log"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"golang.org/x/net/context"
)

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

// interface to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

// Initialize the MDNS service
func initMDNS(peerhost host.Host, rendezvous string) chan peer.AddrInfo {
	// register with service so that we get notified about peer discovery
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	// An hour might be a long long period in practical applications. But this is fine for us
	ser := mdns.NewMdnsService(peerhost, rendezvous, n)
	if err := ser.Start(); err != nil {
		panic(err)
	}
	return n.PeerChan
}

func start_mdns(host host.Host, config Config, ctx context.Context) {

	peerChan := initMDNS(host, config.RendezvousString)
	// This is test shit!!!!!
	// TODO: Rewrite to bootstrap to DHT
	for { // allows multiple peers to join
		peer := <-peerChan // will block untill we discover a peer
		log.Println("Found peer:", peer, ", connecting")

		if err := host.Connect(ctx, peer); err != nil {
			log.Println("Connection failed:", err)
			continue
		}

		// open a stream, this stream will be handled by handleStream other end
		stream, err := host.NewStream(ctx, peer.ID, protocol.ID(config.ProtocolId))

		if err != nil {
			log.Println("Stream open failed", err)
		} else {
			rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

			log.Println(rw)
			log.Println("Connected to:", peer)
		}
	}
}
