package entities

import "time"

// Analysis — данные анализа, привязанные к вызову (1:1)
type Analysis struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	ReceptionID uint `gorm:"uniqueIndex;not null" json:"-"`

	Code  string `json:"code"`
	Title string `json:"title"`
	Price uint   `json:"price"`
}
type AnalysisOrderItem struct {
	ID uint `gorm:"primarykey" json:"id"`

	OrderID uint           `gorm:"not null;index" json:"order_id"`
	Order   *AnalysisOrder `gorm:"foreignKey:OrderID" json:"-"`

	AnalysisID uint      `gorm:"not null;index" json:"analysis_id"`
	Analysis   *Analysis `gorm:"foreignKey:AnalysisID" json:"analysis"`

	// Статус конкретного анализа
	IsCompleted bool       `gorm:"default:false" json:"is_completed"` // Сдан или нет
	CompletedAt *time.Time `json:"completed_at,omitempty"`            // Когда сдан

	PriceAtAssignment uint `gorm:"not null" json:"price_at_assignment"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AnalysisOrder - направление на анализы (промежуточная структура)
type AnalysisOrder struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	OrderNumber string    `gorm:"not null;uniqueIndex" json:"order_number"` // Номер направления (уникальный)

	PatientID uint `gorm:"not null;index" json:"patient_id"`

	OrderItems []AnalysisOrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}
