package main

import (
	"context"
	"fmt"
	"net"
	"syscall"
)

func main() {
	cfg := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			var err error
			c.Control(func(fd uintptr) {
				// Set SO_REUSEPORT at the socket level
				// Note: 15 is SO_REUSEPORT on Linux. Use golang.org/x/sys/unix for cross-platform constants.
				err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, 15, 1)
			})
			return err
		},
	}

	// Start the listener. With SO_REUSEPORT, you can run this exact program 
	// in 5 different terminal windows without getting an "address already in use" error!
	listener, err := cfg.Listen(context.Background(), "tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Listening on :8080 with SO_REUSEPORT enabled.")
	// for { listener.Accept() ... }
}