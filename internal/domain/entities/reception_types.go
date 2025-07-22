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

type ProctologistData struct {
	Complaints         []string `json:"complaints"`
	DigitalExamination string   `json:"digital_examination"`
	Rectoscopy         string   `json:"rectoscopy"`
	Colonoscopy        string   `json:"colonoscopy"`
	Hemorrhoids        bool     `json:"hemorrhoids"`
	AnalFissure        bool     `json:"anal_fissure"`
	Paraproctitis      bool     `json:"paraproctitis"`
	Tumor              bool     `json:"tumor"`
	Diagnosis          string   `json:"diagnosis"`
	Recommendations    string   `json:"recommendations"`
}

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

type PsychiatristData struct {
	MentalStatus   string `json:"mental_status"`
	Mood           string `json:"mood"`
	Affect         string `json:"affect"`
	ThoughtProcess string `json:"thought_process"`
	ThoughtContent string `json:"thought_content"`
	Perception     string `json:"perception"`
	Cognition      string `json:"cognition"`
	Insight        string `json:"insight"`
	Judgment       string `json:"judgment"`

	RiskAssessment struct {
		Suicide  bool `json:"suicide"`
		SelfHarm bool `json:"self_harm"`
		Violence bool `json:"violence"`
	} `json:"risk_assessment"`

	DiagnosisICD string `json:"diagnosis_icd"`
	TherapyPlan  string `json:"therapy_plan"`
}

type UrologistData struct {
	Complaints []string `json:"complaints"`

	Urinalysis struct {
		Color        string `json:"color"`
		Transparency string `json:"transparency"`
		Protein      string `json:"protein"`
		Glucose      string `json:"glucose"`
		Leukocytes   string `json:"leukocytes"`
		Erythrocytes string `json:"erythrocytes"`
	} `json:"urinalysis"`

	Ultrasound          string `json:"ultrasound"`
	ProstateExamination string `json:"prostate_examination"`
	Diagnosis           string `json:"diagnosis"`
	Treatment           string `json:"treatment"`
}

type OtolaryngologistData struct {
	Complaints         []string `json:"complaints"`
	NoseExamination    string   `json:"nose_examination"`
	ThroatExamination  string   `json:"throat_examination"`
	EarExamination     string   `json:"ear_examination"`
	HearingTest        string   `json:"hearing_test"`
	Audiometry         string   `json:"audiometry"`
	VestibularFunction string   `json:"vestibular_function"`
	Endoscopy          string   `json:"endoscopy"`
	Diagnosis          string   `json:"diagnosis"`
	Recommendations    string   `json:"recommendations"`
}

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
