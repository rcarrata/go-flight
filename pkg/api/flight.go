package api

type Flight struct {
	Id          int    `json:"id"`
	Origin      string `json:"name"`
	Duration    int    `json:"duration"`
	Destination string `json:"destination"`
}

type Flights []Flight
