package servers

import (
	"github.com/joshburnsxyz/protorec/handlers"
	"log"
	"fmt"
	"net"
)

func StartTCPServer(host string, port int, messageHandler string) {
	// Start TCP server logic here
	// Use the 'host', 'port', and 'messageHandler' parameters as needed
	// Example code below demonstrates a simple TCP server that accepts connections and logs received messages

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal("Error starting TCP server:", err)
	}
	defer listener.Close()

	fmt.Printf("TCP server listening on %s:%d\n", host, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleTCPConnection(conn, messageHandler)
	}
}

func handleTCPConnection(conn net.Conn, messageHandler string) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading message:", err)
		return
	}

	message := string(buf[:n])

	switch messageHandler {
	case "SIA-CID":
		handlers.SIACIDHandler.Handle(message)
	default:
		log.Println("Invalid message handler specified:", messageHandler)
	}
}
