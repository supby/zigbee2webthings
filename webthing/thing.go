package webthing

type Thing struct {
	AtContext   string     `json:"@context,omitempty"`
	AtType      []string   `json:"@type,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Properties  []Property `json:"properties,omitempty"`
	Actions     []Action   `json:"actions,omitempty"`
	Events      []Event    `json:"events,omitempty"`
	Links       []Link     `json:"links,omitempty"`
}
