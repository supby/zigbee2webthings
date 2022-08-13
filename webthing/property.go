package webthing

type Property struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	AtType      string   `json:"@type,omitempty"`
	Minimum     int      `json:"minimum,omitempty"`
	Maximum     int      `json:"maximum,omitempty"`
	Unit        string   `json:"unit,omitempty"`
	MultipleOf  int      `json:"multipleOf,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	ReadOnly    bool     `json:"readOnly,omitempty"`
	Links       []Link   `json:"links,omitempty"`
}
