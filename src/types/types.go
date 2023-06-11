package types

import (
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
)

type InputData struct {
	HandlerName string
	Input       interface{}
}

type Config struct {
	Host       string
	Port       string
	Secret     string
	Rendezvous string
	ProtocolId string
	Bootstrap  string
	DHTServer  bool
}

type Host struct {
	Host   host.Host
	Dht    *dht.IpfsDHT
	Config Config
}
