// JSONBSchemaFieldWithValues описывает схему поля с его значением
package entities

type CustomField struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Required     bool        `json:"required"`
	Description  string      `json:"description"`
	Format       string      `json:"format,omitempty"`
	MinLength    *int        `json:"min_length,omitempty"`
	MaxLength    *int        `json:"max_length,omitempty"`
	MinValue     *int        `json:"min_value,omitempty"`
	MaxValue     *int        `json:"max_value,omitempty"`
	MinItems     *int        `json:"min_items,omitempty"`
	MaxItems     *int        `json:"max_items,omitempty"`
	Example      interface{} `json:"example,omitempty"`
	DefaultValue interface{} `json:"default_value,omitempty"`
	Value        interface{} `json:"value"`
	KeyFormat    string      `json:"key_format,omitempty"`
	ValueFormat  string      `json:"value_format,omitempty"`
}
