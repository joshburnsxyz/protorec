package handlers

import (
	"testing"
)

func TestMQTTHandler_Handle(t *testing.T) {
	handler := MQTTHandler{}

	// Test case 1: Valid MQTT message
	validMessage := "\x08topic123payload"

	handler.Handle(validMessage)
}
