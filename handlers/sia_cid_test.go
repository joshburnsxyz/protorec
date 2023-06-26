package handlers

import (
	"testing"
	"log"
)

func TestSIACIDHandler_Handle(t *testing.T) {
	// Create an instance of the SIACIDHandler
	handler := SIACIDHandler{}

	// Define a sample SIA-CID message
	message := "0010030051234567890123456789A"

	// Call the Handle method
	output := handler.Handle(message)
	log.Println(output)
}
