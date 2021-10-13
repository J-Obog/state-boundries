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
	"MA": {"Massachusetts", "MA", []string{"NH", "VT", "NY", "CT", "RI"}},
	"MI": {"Michigan", "MI", []string{"OH", "IN", "WI", "IL", "MN"}},
	"MN": {"Minnesota", "MN", []string{"ND", "SD", "IA", "WI"}},
	"MS": {"Mississippi", "MS", []string{"TN", "AL", "LA", "AR"}},
	"MO": {"Missouri", "MO", []string{"IA", "IL", "KY", "TN", "AR", "OK", "KS", "NE"}},
	"MT": {"Montana", "MT", []string{"ND", "SD", "WY", "ID"}},
	"NE": {"Nebraska", "NE", []string{"SD", "IA", "MO", "KS", "CO", "WY"}},
	"NV": {"Nevada", "NV", []string{"OR", "ID", "UT", "AZ", "CA"}},
	"NH": {"New Hampshire", "NH", []string{"ME", "MA", "VT"}},
	"NJ": {"New Jersey", "NJ", []string{"NY", "DE", "PA"}},
	"NM": {"New Mexico", "NM", []string{"CO", "OK", "TX", "AZ"}},
	"NY": {"New York", "NY", []string{"PA", "NJ", "CT", "RI", "MA", "VT"}},
	"NC": {"North Carolina", "NC", []string{"VA", "GA", "SC", "TN"}},
	"ND": {"North Dakota", "ND", []string{"ND", "MN", "MT", "SD"}},
	"OH": {"Ohio", "OH", []string{"PA", "WV", "KY", "IN", "MI"}},
	"OK": {"Oklahoma", "OK", []string{"CO", "KS", "MO", "AR", "TX", "NM"}},
	"OR": {"Oregon", "OR", []string{"WA", "ID", "CA", "NV"}},
	"PA": {"Pennsylvania", "PA", []string{"NY", "NJ", "DE", "MD", "WV", "OH"}},
	"RI": {"Rhode Island", "RI", []string{"MA", "CT"}},
	"SC": {"South Carolina", "SC", []string{"NC", "GA"}},
}