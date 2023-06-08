package node

import (
	"log"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
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
	log.Println("[MDNS] Starting MDNS service")
	// register with service so that we get notified about peer discovery
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	// An hour might be a long long period in practical applications. But this is fine for us
	ser := mdns.NewMdnsService(peerhost, rendezvous, n)
	if err := ser.Start(); err != nil {
		log.Println("[MDNS] Failed to start MDNS service: ", err)
	}

	log.Println("[MDNS] Service started")
	return n.PeerChan
}

func start_mdns(host host.Host, config Config, ctx context.Context) {

	peerChan := initMDNS(host, config.RendezvousString)

	// This is test shit!!!!!
	// TODO: Rewrite to bootstrap to DHT
	for { // allows multiple peers to join
		peer := <-peerChan // will block untill we discover a peer
		log.Println("[MDNS] Found peer:", peer)

		if err := host.Connect(ctx, peer); err != nil {
			log.Println("[MDNS] Connection failed:", err)
			continue
		}

		//Bootstrap peer
	}
}
