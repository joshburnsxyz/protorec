package servers

import (
	"fmt"
	"log"
	"net"

	"github.com/joshburnsxyz/protorec/handlers"
)

func StartUDPServer(host string, port int, messageHandler string) {
	// Start UDP server logic here
	// Use the 'host', 'port', and 'messageHandler' parameters as needed
	// Example code below demonstrates a simple UDP server that listens for incoming messages

	addr := fmt.Sprintf("%s:%d", host, port)
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatal("Error resolving UDP address:", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal("Error starting UDP server:", err)
	}
	defer conn.Close()

	fmt.Printf("UDP server listening on %s:%d\n", host, port)

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		message := string(buf[:n])
		go handleUDPMessage(message, addr, messageHandler)
	}
}

func handleUDPMessage(message string, addr *net.UDPAddr, messageHandler string) {
	switch messageHandler {
	case "SIA-CID":
		handlers.SIACIDHandler.Handle(message)
		// Example: Call the SIACIDHandler to handle the UDP message
	default:
		log.Println("Invalid message handler specified:", messageHandler)
	}
}
