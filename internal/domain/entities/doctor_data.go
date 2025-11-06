package entities

// DoctorData — данные врача, назначенные на вызов (1:1)
type DoctorData struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	ReceptionID    uint   `gorm:"uniqueIndex;not null" json:"-"`
	Name           string `json:"name"`
	Specialization string `json:"specialization"`
}
