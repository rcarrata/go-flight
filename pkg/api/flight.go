package api

type Flight struct {
	Origin      string `json:"name"`
	Id          int    `json:"id"`
	Duration    int    `json:"duration"`
	Destination string `json:"destination"`
}

type Flights []Flight
