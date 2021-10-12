package api

type Location struct {
	Name         string   `json:"name"`
	Abbreviation string   `json:"abbreviation"`
	Borders      []string `json:"borders"`
}

var mapper = map[string]Location{
	"AL": {"Alabama", "AL", []string{"TN", "GA", "MS", "FL"}},
	"AK": {"Alaska", "AK", []string{}},
}