package node

import (
	"context"
	"fmt"
	"log"
	"sync"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func init_DHT(ctx context.Context, host host.Host) *dht.IpfsDHT {
	// init DHT
	kademliaDHT, err := dht.New(ctx, host)
	if err != nil {
		log.Println("[DHT] Fail to init DHT: ", err)
	}
	log.Println("[DHT] Init sucsesfull")
	//Bootstrap
	log.Println("[DHT] Running bootstrap thread")
	if err := kademliaDHT.Bootstrap(ctx); err != nil {
		log.Println("[DHT] Fail to bootstrap DHT: ", err)
	}

	return kademliaDHT
}

func boot(ctx context.Context, config Config, host host.Host) {
	log.Println("[DHT:Bootstrap] Running bootstrap from config: ", config.Bootstrap)
	// Parse configuration
	sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf(config.Bootstrap))
	if err != nil {
		log.Println("[DHT:Bootstrap] Fail to parse multiaddr: ", err)
	}

	// Conver the multiaddr to peerinfo
	var wg sync.WaitGroup
	peerinfo, err := peer.AddrInfoFromP2pAddr(sourceMultiAddr)
	if err != nil {
		log.Println("[DHT:Bootstrap] Fail to parse multiaddr: ", err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := host.Connect(ctx, *peerinfo); err != nil {
			log.Println("[DHT:Bootstrap] Fail to connect!!!")
		} else {
			log.Println("[DHT:Bootstrap] Successfully connected to node: ", *peerinfo)
		}
	}()
	wg.Wait()

}
