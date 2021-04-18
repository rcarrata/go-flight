package api

var currentId int

var flights Flights

func dbFindFlight(id int) Flight {

	// If id of flightId is equal to Id from value DB
	// return f
	for _, f := range flights {
		if f.Id == id {
			return f
		}
	}

	return Flight{}
}

func dbCreateFlight(f Flight) Flight {
	// Increment in 1 the flightId
	currentId += 1

	// Assign/Update the currentId to the Id in flight
	f.Id = currentId

	// Append to a slice to add new flights
	flights = append(flights, f)

	return f
}

// init() is always called, regardless if there's main or not,
// so if you import a package that has an init function, it will be executed.
func init() {
	dbCreateFlight(
		Flight{
			Origin:      "Vancouver",
			Duration:    10,
			Destination: "Rome",
		})
	dbCreateFlight(
		Flight{
			Origin:      "Madrid",
			Duration:    3,
			Destination: "Amsterdam",
		})
}
