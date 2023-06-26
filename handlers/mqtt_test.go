package handlers

import (
	"testing"
	"log"
)

func TestMQTTHandler_Handle(t *testing.T) {
	handler := MQTTHandler{}

	// Test case 1: Valid MQTT message
	validMessage := "\x08topic123payload"

	output := handler.Handle(validMessage)
	log.Println(output)
}
