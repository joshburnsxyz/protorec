package handlers

// MQTTHandler is a handler for MQTT messages
type MQTTHandler struct{}

func (h MQTTHandler) Handle(message string) string {
	if len(message) < 2 {
		return "Invalid MQTT message"
	}

	// Extract the MQTT message topic length and topic
	topicLength := int(message[0])
	topic := message[1 : topicLength+1]

	// Extract the MQTT message payload
	payload := message[topicLength+1:]

	// Construct the result string
	result := "Topic: " + string(topic) + "\n"
	result += "Payload: " + string(payload) + "\n"

	return result
}
