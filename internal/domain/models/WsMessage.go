package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type Message struct {
	Header        string          `json:"header"`
	Text          string          `json:"text"`
	TypeID        uint            `json:"type_id"`
	Type          string          `json:"type"`           // ← добавь это!
	Data          json.RawMessage `json:"data,omitempty"` // ← позволяет вкладывать любой JSON
	Reference     string          `json:"reference"`
	ReferenceID   uint            `json:"reference_id"`
	BroadcastUUID uuid.UUID       `json:"broadcast_uuid"`
	CreatedAt     time.Time
}
