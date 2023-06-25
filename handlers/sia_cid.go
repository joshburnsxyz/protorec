package handlers

import (
	"fmt"
	"strconv"
	"strings"
)

// SIACIDHandler is a handler for SIA-CID messages
type SIACIDHandler struct{}

func (h SIACIDHandler) Handle(message string) {
	result := h.parseSIACIDMessage(message)
	fmt.Println(result)
}

func (h SIACIDHandler) parseSIACIDMessage(message string) string {
	// Define the field lengths and positions in the SIA-CID message
	const (
		accountNumberLength = 3
		eventCodeLength     = 3
		eventDataLength     = 3
		eventQualifierStart = 6
		eventQualifierEnd   = 10
		eventTimeStampStart = 10
		eventTimeStampEnd   = 20
		controlCodeStart    = 20
		controlCodeEnd      = 21
	)

	// Validate the message length
	if len(message) < controlCodeEnd {
		return "Invalid SIA-CID message"
	}

	// Extract the fields from the message
	accountNumber := message[:accountNumberLength]
	eventCode := message[accountNumberLength : accountNumberLength+eventCodeLength]
	eventData := message[accountNumberLength+eventCodeLength : accountNumberLength+eventCodeLength+eventDataLength]
	eventQualifier := message[eventQualifierStart:eventQualifierEnd]
	eventTimeStamp := message[eventTimeStampStart:eventTimeStampEnd]
	controlCode := message[controlCodeStart:controlCodeEnd]

	// Construct the human-readable string
	result := fmt.Sprintf("Account Number: %s\n", accountNumber)
	result += fmt.Sprintf("Event Code: %s\n", eventCode)
	result += fmt.Sprintf("Event Data: %s\n", eventData)
	result += fmt.Sprintf("Event Qualifier: %s\n", eventQualifier)
	result += fmt.Sprintf("Event Time Stamp: %s\n", eventTimeStamp)
	result += fmt.Sprintf("Control Code: %s\n", controlCode)

	return result
}
