package entities

// PatientData — данные пациента, связанные с вызовом
type PatientData struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	ReceptionID uint `gorm:"index;not null" json:"-"`

	PatientID   string      `json:"patient_id" gorm:"column:patient_id;type:varchar(255)"`
	FullName    string      `json:"full_name" gorm:"column:full_name;type:varchar(255)"`
	Age         int         `json:"age" gorm:"column:age;type:int4"`
	BirthDate   string      `json:"birth_date" gorm:"column:birth_date;type:varchar(10)"`
	MobilePhone string      `json:"mobile_phone" gorm:"column:mobile_phone;type:varchar(20)"`
	Policy      Policy      `json:"policy" gorm:"type:jsonb"`      // или embedded
	Certificate Certificate `json:"certificate" gorm:"type:jsonb"` // или embedded
}
