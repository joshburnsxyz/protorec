package servers

import (
	"github.com/joshburnsxyz/protorec/handlers"
	"log"
	"fmt"
	"net"
)

func StartTCPServer(host string, port int, messageHandler string, messageBufferLen int) {
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

		go handleTCPConnection(conn, messageHandler, messageBufferLen)
	}
}

func handleTCPConnection(conn net.Conn, messageHandler string, messageBufferLen int) {
	defer conn.Close()

	buf := make([]byte, messageBufferLen)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading message:", err)
		return
	}

	message := string(buf[:n])

	var handlerOutput string

	switch messageHandler {
	case "SIA-CID":
		h := handlers.SIACIDHandler{}
		handlerOutput = h.Handle(message)
	case "CSV-IP":
		h := handlers.CSVIPHandler{}
		handlerOutput = h.Handle(message)
	case "MQTT":
		h := handlers.MQTTHandler{}
		handlerOutput = h.Handle(message)
	default:
		log.Fatal("Invalid message handler specified:", messageHandler)
	}

	log.Println(handlerOutput)
}
