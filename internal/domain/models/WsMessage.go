package models

import "github.com/gofrs/uuid"

type Message struct {
	Header string `json:"header"`
	Text   string `json:"text"`
	TypeID uint   `json:"type_id"`

	Reference   string `json:"reference"`
	ReferenceID uint   `json:"reference_id"`

	GroupIDs []uint `json:"-"`

	BroadcastUUID uuid.UUID `json:"broadcast_uuid"`
}
