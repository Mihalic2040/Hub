package node

import (
	"context"
	"log"
	"time"

	"github.com/Mihalic2040/Hub/src/types"
	"github.com/fatih/color"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
)

func Rendezvous(ctx context.Context, host host.Host, dht *dht.IpfsDHT, config types.Config) *drouting.RoutingDiscovery {
	log.Println("[DHT:Rendezvous]", "Announcing self!")
	routingDiscovery := drouting.NewRoutingDiscovery(dht)
	dutil.Advertise(ctx, routingDiscovery, config.Rendezvous)
	log.Println("[DHT:Rendezvous]", "Announcing done!")

	blue := color.New(color.FgBlack).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	// log.Println(peers)
	go func() {
		for {
			time.Sleep(time.Second * 10)
			peers, err := dutil.FindPeers(ctx, routingDiscovery, config.Rendezvous)
			if err != nil {
				log.Println("[DHT:Rendezvous]", "Fail to find other peers:", err)
				return
			}
			log.Println(peers)
			for _, peer := range peers {
				err := host.Connect(ctx, peer)
				if err != nil {
					log.Println("[DHT:Rendezvous]", "Fail to connect to peer:", yellow(peer.ID.String()))
					continue
				} else {
					log.Println("[DHT:Rendezvous]", "Connected to peer:", blue(peer.ID.String()))
				}

			}
		}

	}()

	log.Println("[DHT:Rendezvous]", "Successfully")

	return routingDiscovery

}
