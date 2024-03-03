package content

import "golang.org/x/net/html"

// CallbackFunc represents a function that operates on an HTML node.
type CallbackFunc func(*html.Node)


// Event represents an event to be triggered on a specific HTML node type.
type Event struct {
	nodeType     html.NodeType // NodeType specifies the type of HTML node for the event.
	callbackFunc CallbackFunc  // callbackFunc is the function to be triggered when the event occurs.
}

// NewEvent creates a new Event with the specified node type and callback function.
func NewEvent(nodeType html.NodeType, callbackFunc CallbackFunc) Event {
	return Event{
		callbackFunc: callbackFunc,
		nodeType:     nodeType,
	}
}