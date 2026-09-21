package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	// Set 5 second timeout for reading and writing
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	writer := bufio.NewWriter(conn)
	writer.WriteString("Welcome to the secure TCP server.\n")
	writer.Flush()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Received:", text)
		
		// Reset the deadline after a successful read
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		writer.WriteString("ACK: " + text + "\n")
		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Connection dropped or timed out:", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("TCP Server listening on port 9090")
	// Uncomment to run:
	/*
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
	*/
}