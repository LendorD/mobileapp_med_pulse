package entities

type AllergologistData struct {
	Complaints      []string `json:"complaints"`
	AllergenHistory string   `json:"allergen_history"`

	SkinTests []struct {
		Allergen string `json:"allergen"`
		Reaction string `json:"reaction"`
	} `json:"skin_tests"`

	IgELevel        float32 `json:"ige_level"`
	Immunotherapy   bool    `json:"immunotherapy"`
	Diagnosis       string  `json:"diagnosis"`
	Recommendations string  `json:"recommendations"`
}

// ToDocumentWithValues для AllergologistData
func (a *AllergologistData) ToDocumentWithValues() SpecializationDataDocument {
	fields := []CustomField{
		{
			Name:         "complaints",
			Type:         "array",
			Required:     false,
			Description:  "Жалобы пациента",
			MaxItems:     intPtr(50),
			Format:       "[]string",
			MaxLength:    intPtr(200),
			Example:      []string{"зуд кожи", "ринит"},
			DefaultValue: []string{},
			Value:        a.Complaints,
		},
		{
			Name:         "allergen_history",
			Type:         "string",
			Required:     false,
			Description:  "Аллергоанамнез",
			MaxLength:    intPtr(1000),
			Example:      "Пыльца березы, домашняя пыль",
			DefaultValue: "",
			Value:        a.AllergenHistory,
		},
		{
			Name:         "skin_tests",
			Type:         "array",
			Required:     false,
			Description:  "Кожные пробы",
			MaxItems:     intPtr(20),
			Example:      []map[string]string{{"allergen": "пыльца", "reaction": "положительная"}},
			DefaultValue: []map[string]string{},
			Value:        a.SkinTests,
		},
		{
			Name:         "ige_level",
			Type:         "number",
			Required:     false,
			Description:  "Уровень IgE",
			MinValue:     intPtr(0),
			MaxValue:     intPtr(10000),
			Example:      150.5,
			DefaultValue: float32(0),
			Value:        a.IgELevel,
		},
		{
			Name:         "immunotherapy",
			Type:         "boolean",
			Required:     false,
			Description:  "Иммунотерапия",
			Example:      true,
			DefaultValue: false,
			Value:        a.Immunotherapy,
		},
		{
			Name:         "diagnosis",
			Type:         "string",
			Required:     false,
			Description:  "Аллергологический диагноз",
			MaxLength:    intPtr(1000),
			Example:      "Аллергический ринит",
			DefaultValue: "",
			Value:        a.Diagnosis,
		},
		{
			Name:         "recommendations",
			Type:         "string",
			Required:     false,
			Description:  "Рекомендации аллерголога",
			MaxLength:    intPtr(2000),
			Example:      "Избегать аллергенов, СЗП",
			DefaultValue: "",
			Value:        a.Recommendations,
		},
	}

	return SpecializationDataDocument{
		DocumentType: "allergologist_data",
		Fields:       fields,
	}
}
