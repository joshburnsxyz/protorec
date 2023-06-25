package handlers

// Handler defines the interface for message handlers
type Handler interface {
	Handle(message string)
}
