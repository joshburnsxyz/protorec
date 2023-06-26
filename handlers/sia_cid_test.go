package handlers

import (
	"testing"
)

func TestSIACIDHandler_Handle(t *testing.T) {
	// Create an instance of the SIACIDHandler
	handler := SIACIDHandler{}

	// Define a sample SIA-CID message
	message := "0010030051234567890123456789A"

	// Call the Handle method
	handler.Handle(message)
}
