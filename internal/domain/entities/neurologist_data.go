package entities

// Структуры данных для специализаций

type NeurologistData struct {
	Reflexes         map[string]string `json:"reflexes"`
	MuscleStrength   map[string]int    `json:"muscle_strength"`
	Sensitivity      string            `json:"sensitivity"`
	CoordinationTest string            `json:"coordination_tests"`
	Gait             string            `json:"gait"`
	Speech           string            `json:"speech"`
	Memory           string            `json:"memory"`
	CranialNerves    string            `json:"cranial_nerves"`
	Complaints       []string          `json:"complaints"`
	Diagnosis        string            `json:"diagnosis"`
	Recommendations  string            `json:"recommendations"`
}

// ToDocumentWithValues для NeurologistData
func (n NeurologistData) ToDocumentWithValues() SpecializationDataDocument {
	fields := []CustomField{
		{
			Name:         "reflexes",
			Type:         "object",
			Required:     false,
			Description:  "Словарь рефлексов пациента",
			Format:       "map[string]string",
			MaxItems:     intPtr(20),
			KeyFormat:    "название рефлекса",
			ValueFormat:  "характер рефлекса",
			MaxLength:    intPtr(100),
			Example:      map[string]string{"knee": "увеличен"},
			DefaultValue: map[string]string{},
			Value:        n.Reflexes,
		},
		{
			Name:         "muscle_strength",
			Type:         "object",
			Required:     false,
			Description:  "Сила мышц по регионам",
			Format:       "map[string]int",
			MaxItems:     intPtr(20),
			KeyFormat:    "локализация",
			ValueFormat:  "шкала MRC 0-5",
			MinValue:     intPtr(0),
			MaxValue:     intPtr(5),
			Example:      map[string]int{"left_arm": 4},
			DefaultValue: map[string]int{},
			Value:        n.MuscleStrength,
		},
		{
			Name:         "sensitivity",
			Type:         "string",
			Required:     false,
			Description:  "Описание чувствительности",
			MaxLength:    intPtr(500),
			Example:      "снижена в distal отделах",
			DefaultValue: "",
			Value:        n.Sensitivity,
		},
		{
			Name:         "coordination_tests",
			Type:         "string",
			Required:     false,
			Description:  "Результаты координационных проб",
			MaxLength:    intPtr(1000),
			Example:      "пальце-носовая проба нарушена",
			DefaultValue: "",
			Value:        n.CoordinationTest,
		},
		{
			Name:         "gait",
			Type:         "string",
			Required:     false,
			Description:  "Характер походки",
			MaxLength:    intPtr(300),
			Example:      "атаксический",
			DefaultValue: "",
			Value:        n.Gait,
		},
		{
			Name:         "speech",
			Type:         "string",
			Required:     false,
			Description:  "Речевые нарушения",
			MaxLength:    intPtr(500),
			Example:      "дизартрия",
			DefaultValue: "",
			Value:        n.Speech,
		},
		{
			Name:         "memory",
			Type:         "string",
			Required:     false,
			Description:  "Состояние памяти",
			MaxLength:    intPtr(400),
			Example:      "краткосрочная память снижена",
			DefaultValue: "",
			Value:        n.Memory,
		},
		{
			Name:         "cranial_nerves",
			Type:         "string",
			Required:     false,
			Description:  "Состояние черепных нервов",
			MaxLength:    intPtr(500),
			Example:      "II, III пары нарушены",
			DefaultValue: "",
			Value:        n.CranialNerves,
		},
		{
			Name:         "complaints",
			Type:         "array",
			Required:     false,
			Description:  "Жалобы пациента",
			MaxItems:     intPtr(50),
			Format:       "[]string",
			MaxLength:    intPtr(200),
			Example:      []string{"головная боль"},
			DefaultValue: []string{},
			Value:        n.Complaints,
		},
		{
			Name:         "diagnosis",
			Type:         "string",
			Required:     false,
			Description:  "Неврологический диагноз",
			MaxLength:    intPtr(1000),
			Example:      "Мигрень с аурой",
			DefaultValue: "",
			Value:        n.Diagnosis,
		},
		{
			Name:         "recommendations",
			Type:         "string",
			Required:     false,
			Description:  "Рекомендации невролога",
			MaxLength:    intPtr(2000),
			Example:      "Повторный прием через 2 недели",
			DefaultValue: "",
			Value:        n.Recommendations,
		},
	}

	return SpecializationDataDocument{
		DocumentType: "neurologist_data",
		Fields:       fields,
	}
}
