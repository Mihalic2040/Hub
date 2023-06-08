package types

type InputData struct {
	HandlerName string
	Input       interface{}
}

type Config struct {
	Host             string
	Port             string
	RendezvousString string
	ProtocolId       string
	Bootstrap        string
}
