package handlers

import (
	"fmt"
)

// SIACIDHandler is a handler for SIA-CID messages
type SIACIDHandler struct{}

func (h SIACIDHandler) Handle(message string) {
	fmt.Println("Handling SIA-CID message:", message)
}
