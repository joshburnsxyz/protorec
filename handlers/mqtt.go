package handlers

import (
	"log"
)

// MQTTHandler is a handler for MQTT messages
type MQTTHandler struct{}

func (h MQTTHandler) Handle(message string) {
	if len(message) < 2 {
		log.Println("Invalid MQTT message")
	}

	// Extract the MQTT message topic length and topic
	topicLength := int(message[0])
	topic := message[1 : topicLength+1]

	// Extract the MQTT message payload
	payload := message[topicLength+1:]

	log.Printf("Topic: %s\nPayload: %s\n", string(topic), string(payload))
}
