package node

import (
	"context"
	"log"

	"github.com/Mihalic2040/Hub/src/types"
	"github.com/fatih/color"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
)

func rendezvous(ctx context.Context, host host.Host, dht *dht.IpfsDHT, config types.Config) {
	log.Println("[DHT:Rendezvous]", "Announcing self!")
	routingDiscovery := drouting.NewRoutingDiscovery(dht)
	dutil.Advertise(ctx, routingDiscovery, config.Rendezvous)
	log.Println("[DHT:Rendezvous]", "Announcing done!")
	peers, err := dutil.FindPeers(ctx, routingDiscovery, config.Rendezvous)
	if err != nil {
		log.Println("[DHT:Rendezvous]", "Fail to find other peers:", err)
		return
	}
	go func() {
		blue := color.New(color.FgBlack).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		for _, peer := range peers {
			err := host.Connect(ctx, peer)
			if err != nil {
				log.Println("[DHT:Rendezvous]", "Fail to connect to peer:", yellow(peer.ID.String()))
				continue
			} else {
				log.Println("[DHT:Rendezvous]", "Connected to peer:", blue(peer.ID.String()))
			}

		}
	}()

}
