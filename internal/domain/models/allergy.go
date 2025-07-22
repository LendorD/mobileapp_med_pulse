package models

// AllergyResponse - ответ с информацией об аллергии
// @Description Ответ с названием аллергии
type AllergyResponse struct {
	Name string `json:"name" example:"Пыльца"` // Только название аллергии
}
