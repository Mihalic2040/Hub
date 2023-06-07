package node

import (
	"bufio"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
)

func stream_handler(stream network.Stream) {
	// fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	fmt.Println(rw)

	// 'stream' will stay open until you close it (or the other side closes it).
}
