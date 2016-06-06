package invdendpoint

import (
	"encoding/json"
)

type Event struct {
	Id   int64           `json:"id"`             //The event’s unique ID
	Type string          `json:"type"`           //Event type
	Data json.RawMessage `json:"data,omitempty"` //Contains an object property with the object that was subject of the event and an optional previous property for object.updated events that is a hash of the old values that changed during the event
}
