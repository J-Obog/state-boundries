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
	"HI": {"Hawaii", "HI", []string{}},
	"ID": {"Idaho", "ID", []string{"MT", "WY", "NV", "UT", "WA", "OR"}},
	"IL": {"Illinois", "IL", []string{"WI", "IN", "KY", "MO", "IA"}},
	"IN": {"Indiana", "IN", []string{"MI", "OH", "KY", "IL"}},
	"IA": {"Iowa", "IA", []string{"WI", "IL", "MO", "NE", "SD", "MN"}},
	"KS": {"Kansas", "KS", []string{"NE", "MO", "OK", "CO"}},
	"KY": {"Kentucky", "KY", []string{"WV", "VA", "TN", "MO", "IL", "IN", "OH"}},
	"LA": {"Louisiana", "LA", []string{"AR", "MS", "TX"}},
	"ME": {"Maine", "ME", []string{"NH"}},
	"MD": {"Maryland", "MD", []string{"PA", "DE", "VA", "WV"}},
}