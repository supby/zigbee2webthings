package webthing

type Input struct {
	AtType     string `json:"@type,omitempty"`
	Type       string `json:"type,omitempty"`
	Minimum    int    `json:"minimum,omitempty"`
	Maximum    int    `json:"maximum,omitempty"`
	Unit       string `json:"unit,omitempty"`
	MultipleOf int    `json:"multipleOf,omitempty"`
}

type Action struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AtType      string `json:"@type,omitempty"`
	Links       []Link `json:"links"`
	Input       Input  `json:"input"`
}
