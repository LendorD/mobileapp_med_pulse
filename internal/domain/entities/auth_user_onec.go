package entities

type AuthUser struct {
	ID       uint   `gorm:"primaryKey"`
	Login    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
