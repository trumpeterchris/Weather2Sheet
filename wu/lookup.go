package lookup

import "fmt"

type Lookup struct {
	Location SLocation
}

type SLocation struct {
	Nearby_weather_stations Nearby_weather_stations
}

type Nearby_weather_stations struct {
	Airport Airport
}

type Airport struct {
	Station []Station
}

type Station struct {
	City string
	Icao string
}

// printLookup prints nearby stations
func PrintLookup(obs *Lookup) {
	station := obs.Location.Nearby_weather_stations.Airport.Station
	if len(station) == 0 {
		fmt.Println("No area stations")
	} else {
		for _, s := range station {
			fmt.Printf("%s: %s\n", s.City, s.Icao)
		}
	}
}
