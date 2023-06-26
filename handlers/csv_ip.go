package handlers

import (
	"strings"
)

// CSVIPHandler is a handler for CSV-IP messages
type CSVIPHandler struct{}

func (h CSVIPHandler) Handle(message string) string {
	// Parse CSV-IP message
	fields := strings.Split(message, ",")
	if len(fields) < 3 {
		return "Invalid CSV-IP message"
	}

	// Extract fields
	ipAddress := fields[0]
	port := fields[1]
	data := fields[2:]

	// Construct the result string
	result := "IP Address: " + ipAddress + "\n"
	result += "Port: " + port + "\n"
	result += "Data: " + strings.Join(data, ",") + "\n"

	return result
}
