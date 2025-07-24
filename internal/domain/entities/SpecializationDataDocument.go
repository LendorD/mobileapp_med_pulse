package entities

type SpecializationDataDocument struct {
	DocumentType string        `json:"document_type"`
	Fields       []CustomField `json:"fields"`
}

// Вспомогательная функция для создания указателя на int
func intPtr(i int) *int {
	return &i
}
