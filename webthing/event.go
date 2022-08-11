package webthing

type Event struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	AtType      string `json:"@type,omitempty"`
	Minimum     int    `json:"minimum,omitempty"`
	Maximum     int    `json:"maximum,omitempty"`
	Unit        string `json:"unit,omitempty"`
	MultipleOf  int    `json:"multipleOf,omitempty"`
	Links       []Link `json:"links,omitempty"`
}
