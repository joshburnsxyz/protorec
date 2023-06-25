package handlers

import (
	"strings"
	"log"
)

// CSVIPHandler is a handler for CSV-IP messages
type CSVIPHandler struct{}

func (h CSVIPHandler) Handle(message string) {
	log.Println("Handling CSV-IP message:", message)

	// Parse CSV-IP message
	fields := strings.Split(message, ",")
	if len(fields) < 3 {
		log.Println("Invalid CSV-IP message:", message)
		return
	}

	// Extract fields
	ipAddress := fields[0]
	port := fields[1]
	data := fields[2:]

	// Process the CSV-IP message as needed
	log.Println("IP Address:", ipAddress)
	log.Println("Port:", port)
	log.Println("Data:", data)
}
