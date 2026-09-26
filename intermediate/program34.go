package main

import "fmt"

type IPAddress [4]byte

// String implements the fmt.Stringer interface
func (ip IPAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

type Server struct {
	Name string
	IP   IPAddress
}

func main() {
	srv := Server{
		Name: "Database Node 1",
		IP:   IPAddress{192, 168, 1, 100},
	}

	// Go automatically calls the String() method on the IPAddress field
	fmt.Printf("Server %s is running at %s\n", srv.Name, srv.IP)
	fmt.Println(srv.IP) 
}