package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// A
	host := "localhost"
	port := 8080

	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Server can't listening on port %d: %v", port, err)
	}
	defer server.Close()

	log.Printf("TCP is listening on %s:%d", host, port)

	c := make(chan net.Conn, 10)
	go handleClient(c)

	for {
		// B
		conn, err := server.Accept()
		if err != nil {
			log.Fatalf("Accept failed: %v", err)
		}

		log.Printf("new client accepted: %s", conn.RemoteAddr().String())
		c <- conn
	}
}

// C
func handleClient(c chan net.Conn) {
	for {
		conn := <-c
		defer conn.Close()
		clientAddr := conn.RemoteAddr().String()
		log.Printf("Handle Request from [%s]", clientAddr)

		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Printf("Client [%s[] Error: %v", clientAddr, err)
				break
			}
			data := string(buf[:n])
			log.Printf("Received data [%v] from Client [%s]", data, clientAddr)
			_, err = conn.Write([]byte("OK\n"))
			if err != nil {
				log.Printf("Reply to Client [%s] failed: %v", clientAddr, err)
				break
			}
		}
	}

}
