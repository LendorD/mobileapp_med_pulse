package entities

import (
	"time"
)

// OneCReception — заявка на вызов скорой из 1С
type OneCReception struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	CallID      string     `gorm:"uniqueIndex;type:varchar(255)" json:"call_id"`
	Address     string     `json:"address"`
	Phone       string     `json:"phone"`
	Status      CallStatus `gorm:"not null;default:'received'" json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	MedServices []byte     `gorm:"type:jsonb" json:"med_services,omitempty"`

	// Отношения (только прямые ссылки — без обратных!)
	Patients   []PatientData `gorm:"foreignKey:ReceptionID" json:"patients,omitempty"`
	Doctor     DoctorData    `gorm:"foreignKey:ReceptionID" json:"doctor"`
	Receptions []Receptions  `gorm:"foreignKey:ReceptionID" json:"receptions,omitempty"`
	Flg        Flg           `gorm:"foreignKey:ReceptionID" json:"flg"`
	Analysis   Analysis      `gorm:"foreignKey:ReceptionID" json:"analysis"`
}

// CallStatus — статус вызова
type CallStatus string

const (
	CallStatusCompleted CallStatus = "completed" // исправлено: было "compleated"
	CallStatusWork      CallStatus = "process"   // исправлено: было "proccess"
)

// Receptions — заключения врача по вызову (1:N)
type Receptions struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	ReceptionID uint `gorm:"index;not null" json:"-"`

	Data []byte `gorm:"type:jsonb" json:"data,omitempty"`
}
