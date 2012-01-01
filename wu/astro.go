package astro

import (
	"fmt"
	"strconv"
)

type AstroConditions struct {
	Moon_phase Moon_phase
	Sunrise    Sunrise
	Sunset     Sunset
}

type Moon_phase struct {
	PercentIlluminated string
	AgeOfMoon          string
	Sunrise            Sunrise
	Sunset             Sunset
}

type Sunrise struct {
	Hour   string
	Minute string
}

type Sunset struct {
	Hour   string
	Minute string
}

// printAstro prints the lunar and solar informtion for a given station to standard out
func PrintAstro(obs *AstroConditions, stationId string) {

	var age, _ = strconv.Atoi(obs.Moon_phase.AgeOfMoon)
	var moonDesc string

	// Calculate traditional description of lunar phase

	switch {
	case age < 2:
		moonDesc = "New moon"
	case age < 6:
		moonDesc = "Waxing crescent"
	case age < 9:
		moonDesc = "First quarter"
	case age < 13:
		moonDesc = "Waxing gibbous"
	case age < 17:
		moonDesc = "Full moon"
	case age < 20:
		moonDesc = "Waning gibbous"
	case age < 24:
		moonDesc = "Last quarter"
	case age < 28:
		moonDesc = "Waning crescent"
	}
	sr := obs.Moon_phase.Sunrise
	ss := obs.Moon_phase.Sunset
	percent := obs.Moon_phase.PercentIlluminated
	fmt.Printf("Moon Phase: %s (%s%% illuminated)\n", moonDesc, percent)
	fmt.Printf("Sunrise   : %s:%s\n", sr.Hour, sr.Minute)
	fmt.Printf("Sunset    : %s:%s\n", ss.Hour, ss.Minute)
}
