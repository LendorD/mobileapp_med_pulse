package entities

type TraumatologistData struct {
	InjuryType       string `json:"injury_type"`
	InjuryMechanism  string `json:"injury_mechanism"`
	Localization     string `json:"localization"`
	XRayResults      string `json:"xray_results"`
	CTResults        string `json:"ct_results"`
	MRIResults       string `json:"mri_results"`
	Fracture         bool   `json:"fracture"`
	Dislocation      bool   `json:"dislocation"`
	Sprain           bool   `json:"sprain"`
	Contusion        bool   `json:"contusion"`
	WoundDescription string `json:"wound_description"`
	TreatmentPlan    string `json:"treatment_plan"`
}

// ToDocumentWithValues для TraumatologistData
func (t *TraumatologistData) ToDocumentWithValues() SpecializationDataDocument {
	fields := []CustomField{
		{
			Name:         "injury_type",
			Type:         "string",
			Required:     true,
			Description:  "Тип травмы",
			MaxLength:    intPtr(200),
			Example:      "Перелом",
			DefaultValue: "",
			Value:        t.InjuryType,
		},
		{
			Name:         "injury_mechanism",
			Type:         "string",
			Required:     false,
			Description:  "Механизм получения травмы",
			MaxLength:    intPtr(500),
			Example:      "Падение с высоты",
			DefaultValue: "",
			Value:        t.InjuryMechanism,
		},
		{
			Name:         "localization",
			Type:         "string",
			Required:     true,
			Description:  "Локализация травмы",
			MaxLength:    intPtr(200),
			Example:      "Правая голень",
			DefaultValue: "",
			Value:        t.Localization,
		},
		{
			Name:         "xray_results",
			Type:         "string",
			Required:     false,
			Description:  "Результаты рентгена",
			MaxLength:    intPtr(1000),
			Example:      "Перелом малоберцовой кости",
			DefaultValue: "",
			Value:        t.XRayResults,
		},
		{
			Name:         "ct_results",
			Type:         "string",
			Required:     false,
			Description:  "Результаты КТ",
			MaxLength:    intPtr(1500),
			Example:      "Дополнительных повреждений не выявлено",
			DefaultValue: "",
			Value:        t.CTResults,
		},
		{
			Name:         "mri_results",
			Type:         "string",
			Required:     false,
			Description:  "Результаты МРТ",
			MaxLength:    intPtr(1500),
			Example:      "Повреждение мягких тканей",
			DefaultValue: "",
			Value:        t.MRIResults,
		},
		{
			Name:         "fracture",
			Type:         "boolean",
			Required:     false,
			Description:  "Перелом",
			Example:      true,
			DefaultValue: false,
			Value:        t.Fracture,
		},
		{
			Name:         "dislocation",
			Type:         "boolean",
			Required:     false,
			Description:  "Вывих",
			Example:      false,
			DefaultValue: false,
			Value:        t.Dislocation,
		},
		{
			Name:         "sprain",
			Type:         "boolean",
			Required:     false,
			Description:  "Растяжение",
			Example:      true,
			DefaultValue: false,
			Value:        t.Sprain,
		},
		{
			Name:         "contusion",
			Type:         "boolean",
			Required:     false,
			Description:  "Ушиб",
			Example:      false,
			DefaultValue: false,
			Value:        t.Contusion,
		},
		{
			Name:         "wound_description",
			Type:         "string",
			Required:     false,
			Description:  "Описание раны",
			MaxLength:    intPtr(1000),
			Example:      "Ранение мягких тканей без повреждения костей",
			DefaultValue: "",
			Value:        t.WoundDescription,
		},
		{
			Name:         "treatment_plan",
			Type:         "string",
			Required:     false,
			Description:  "План лечения",
			MaxLength:    intPtr(2000),
			Example:      "Иммобилизация, последующая реабилитация",
			DefaultValue: "",
			Value:        t.TreatmentPlan,
		},
	}

	return SpecializationDataDocument{
		DocumentType: "traumatologist_data",
		Fields:       fields,
	}
}
