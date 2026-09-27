package main

import (
	"fmt"
	"io"
	"net"
)

func handleProxyConnection(clientConn net.Conn, targetAddr string) {
	defer clientConn.Close()

	// Connect to the target destination
	backendConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		fmt.Println("Backend connection failed:", err)
		return
	}
	defer backendConn.Close()

	fmt.Printf("Proxying traffic: %s <-> %s\n", clientConn.RemoteAddr(), targetAddr)

	// Channel to wait for both streams to close
	done := make(chan struct{})

	// Stream: Client -> Backend
	go func() {
		io.Copy(backendConn, clientConn)
		done <- struct{}{}
	}()

	// Stream: Backend -> Client
	go func() {
		io.Copy(clientConn, backendConn)
		done <- struct{}{}
	}()

	<-done // Block until one of the streams closes
}

func main() {
	targetServer := "google.com:80" // Forward traffic here
	localAddr := ":8081"

	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("L4 TCP Proxy listening on %s -> forwarding to %s\n", localAddr, targetServer)

	// Uncomment to run:
	/*
	for {
		conn, err := listener.Accept()
		if err == nil {
			go handleProxyConnection(conn, targetServer)
		}
	}
	*/
}