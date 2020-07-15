package interval

// A time interval where the lowest unit is a day
type DiurnalInterval struct {
	Interval DiurnalIntervalInterval `json:"interval"`
	Count    int64                   `json:"count"`
}
