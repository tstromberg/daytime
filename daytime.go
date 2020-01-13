// a bad implementation of RFC867
//
// usage:
//    go build -o daytime
//    ./daytime --port 1313 &
package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"
)

var portFlag = flag.Int("port", 1313, "port to listen on")

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *portFlag))
	if err != nil {
		panic(fmt.Sprintf("listen: %v", err))
	}

	fmt.Printf("listening on port %d\n", *portFlag)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("accept failed: %v\n", err)
			continue
		}

		go daytime(conn)
	}
}

// daytime writes a date to the connection. No big deal.
func daytime(c io.WriteCloser) {
	defer c.Close()
	s := fmt.Sprintf("%s\n", time.Now())
	c.Write([]byte(s))
	c.Close()
}
