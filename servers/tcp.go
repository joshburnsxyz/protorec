package servers

import (
	"github.com/joshburnsxyz/protorec/handlers"
	"log"
	"fmt"
	"net"
)

func StartTCPServer(host string, port int, messageHandler string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal("Error starting TCP server:", err)
	}
	defer listener.Close()

	log.Printf("TCP server listening on %s:%d\n", host, port)
	log.Printf("Using %s parser\n", messageHandler)

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
		h := handlers.SIACIDHandler{}
		h.Handle(message)
	case "CSV-IP":
		h := handlers.CSVIPHandler{}
		h.Handle(message)
	case: "MQTT":
		h := handlers.MQTTHandler{}
		h.Handle(message)
	default:
		log.Println("Invalid message handler specified:", messageHandler)
	}
}
