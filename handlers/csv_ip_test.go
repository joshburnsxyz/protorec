package handlers

import (
	"testing"
)

func TestCSVIPHandler_Handle(t *testing.T) {
	// Create an instance of the CSVIPHandler
	handler := CSVIPHandler{}

	// Define a sample CSV-IP message
	message := "127.0.0.1,8080,data1,data2,data3"

	// Call the Handle method
	handler.Handle(message)
}
