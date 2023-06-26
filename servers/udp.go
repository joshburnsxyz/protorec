package servers

import (
	"fmt"
	"log"
	"net"

	"github.com/joshburnsxyz/protorec/handlers"
)

func StartUDPServer(host string, port int, messageHandler string) {
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

	log.Printf("UDP server listening on %s:%d\n", host, port)
	log.Printf("Using %s parser\n", messageHandler)

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
		h := handlers.SIACIDHandler{}
		h.Handle(message)
	case "CSV-IP":
		h := handlers.CSVIPHandler{}
		h.Handle(message)
	case "MQTT":
		h := handlers.MQTTHandler{}
		h.Handle(message)
	default:
		log.Println("Invalid message handler specified:", messageHandler)
	}
}
