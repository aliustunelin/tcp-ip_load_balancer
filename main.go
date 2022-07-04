package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// pre define values
var (
	// keep the server names
	counter int

	// TODO configurale
	listenAddr = "localhost:8080"

	server = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
	port = ""
)

/*
--> General Infos
first off all I need to accept incoming connections, this is "net" lib job from stdlib

*/
func main() {

	// func net.Listen(network string, address string) (net.Listener, error)
	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatalf("failed to listen: %s", err)

		// defer method run while close the main
	}
	defer listener.Close()

	for {
		//forever loop to accept new connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err)
		}

		backend := chooseBackend()
		// catch go routine solution if get error
		go func() {
			err := proxt_to_backend(backend, conn)
			if err != nil {
				log.Println("Warning: proxiying fail!!!", err)
			}
		}()

	}

}

func chooseBackend() string {
	// TODO return randomly
	s := server[counter%len(server)]
	counter++
	return s
}

func proxt_to_backend(backend string, c net.Conn) error {
	//dial is connection to any backend server using tcp
	backend_connection, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("failed to connection to backend %s:%d", backend, err)
	}

	//then roadmap:
	// c->backen_conn
	//backen_conn->c
	//we using this cycle but both need work same time and we choose go routine
	go io.Copy(backend_connection, c)

	go io.Copy(c, backend_connection)

	return nil
}
