package api

type Location struct {
	Name         string   `json:"name"`
	Abbreviation string   `json:"abbreviation"`
	Borders      []string `json:"borders"`
}

var mapper = map[string]Location{
	"AL": {"Alabama", "AL", []string{"TN", "GA", "MS", "FL"}},
	"AK": {"Alaska", "AK", []string{}},
	"AZ": {"Arizona", "AZ", []string{"CA", "NV", "UT", "CO", "NM"}},
	"AR": {"Arkansas", "AR", []string{"MO", "TN", "MS", "LA", "TX", "OK"}},
	"CA": {"California", "CA", []string{"OR", "NV", "AZ"}},
	"CO": {"Colorado", "CO", []string{"WY", "NE", "KS", "OK", "NM", "UT"}},
	"CT": {"Connecticut", "CT", []string{"NY", "MA", "RI"}},
	"DE": {"Delaware", "DE", []string{"PA", "NJ", "MD"}},
	"FL": {"Florida", "FL", []string{"GA", "AL"}},
	"GA": {"Georgia", "GA", []string{"TN", "SC", "FL", "AL"}},
}