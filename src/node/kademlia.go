package node

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Mihalic2040/Hub/src/types"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func bootstrap(ctx context.Context, dht *dht.IpfsDHT) {
	//Bootstrap
	log.Println("[DHT] Running bootstrap thread")
	if err := dht.Bootstrap(ctx); err != nil {
		log.Println("[DHT] Fail to bootstrap DHT: ", err)
	}

}

func init_DHT(ctx context.Context, host host.Host, config types.Config) *dht.IpfsDHT {
	// init DHT

	if config.DHTServer == true {
		kademliaDHT, err := dht.New(ctx, host, dht.Mode(dht.ModeAutoServer))
		if err != nil {
			log.Println("[DHT] Fail to init DHT: ", err)
		}
		log.Println("[DHT:Server MODE] Init sucsesfull")
		return kademliaDHT
	} else {
		kademliaDHT, err := dht.New(ctx, host)
		if err != nil {
			log.Println("[DHT] Fail to init DHT: ", err)
		}
		log.Println("[DHT] Init sucsesfull")
		return kademliaDHT
	}

}

func boot(ctx context.Context, config types.Config, host host.Host) {
	if config.Bootstrap == "" {
		return
	}
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
			log.Println("[DHT:Bootstrap] Fail to connect:", err)
		} else {
			log.Println("[DHT:Bootstrap] Successfully connected to node: ", *peerinfo)
		}
	}()
	wg.Wait()

}
