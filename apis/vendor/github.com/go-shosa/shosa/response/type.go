package response

// Page defines the response type for REST hateoas.
// This type is used REST API outputs index data.
type Page struct {
	Href       string        `json:"href"`
	Offset     int           `json:"offset"`
	Limit      int           `json:"limit"`
	First      *string       `json:"first"`
	Previous   *string       `json:"previous"`
	Next       *string       `json:"next"`
	Last       *string       `json:"last"`
	Items      []interface{} `json:"items"`
	TotalItems int           `json:"total"`
}

// Link defines the information to navigate the service.
type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

// Links bundles Link structure.
type Links []*Link
