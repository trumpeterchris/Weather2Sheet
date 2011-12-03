//
//  wu - weather underground 
//

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"strconv"
	"json"
	"http"
	"flag"
)

type Config struct {
	Key     string
	Station string
}

var (
	help, version, doall, doalmanac, doalerts, doconditions, dolookup, doforecast, doastro bool
	conf                                                                                   Config
)

const defaultStation = "KLNK"

// GetVersion returns the version of the package
func GetVersion() string {
	return "2.1"
}

// GetConf returns the API key and weather station from
// the configuration file at $HOME/.condrc
func ReadConf() {

	var b []byte
	b, err := ioutil.ReadFile(os.Getenv("HOME") + "/.condrc")

	if err == nil {
		jsonErr := json.Unmarshal(b, &conf)
		CheckError(jsonErr)
	} else {
		fmt.Println("You must create a .condrc file in $HOME.")
		os.Exit(1)
	}
}

// Options handles commandline options and returns a 
// possibly updated weather station string
func Options() string {

	var station, sconf string

	if conf.Station == "" {
		sconf = defaultStation
	} else {
		sconf = conf.Station
	}

	flag.BoolVar(&doall, "all", false, "show all weather data")
	flag.BoolVar(&doconditions, "conditions", false, "Reports the current weather conditions")
	flag.BoolVar(&doalerts, "alerts", false, "Reports any active weather alerts")
	flag.BoolVar(&dolookup, "lookup", false, "Lookup the codes for the weather stations in a particular area")
	flag.BoolVar(&doastro, "astro", false, "Reports sunrise, sunset, and lunar phase")
	flag.BoolVar(&doforecast, "forecast", false, "Reports the current forecast")
	flag.BoolVar(&doalmanac, "almanac", false, "Reports average high, low and record temperatures")
	flag.BoolVar(&help, "h", false, "Print this message")
	flag.BoolVar(&version, "V", false, "Print the version number")
	flag.StringVar(&station, "s", sconf,
		"Weather station: \"city, state-abbreviation\", (US or Canadian) zipcode, 3- or 4-letter airport code, or LAT,LONG")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println("conditions " + GetVersion())
		fmt.Println("Copyright (C) 2011 by Stephen Ramsay")
		fmt.Println("Data courtesy of Weather Underground, Inc.")
		fmt.Println("is subject to Weather Underground Data Feed")
		fmt.Println("Terms of Service.  The program itself is free")
		fmt.Println("software, and you are welcome to redistribute")
		fmt.Println("it under certain conditions.  See LICENSE for")
		fmt.Println("details.")
		os.Exit(0)
	}

	// Trap for city-state combinations (e.g. "San Francisco, CA") and
	// make them URL-friendly (e.g. "CA/SanFranciso")

	cityStatePattern := regexp.MustCompile("([A-Za-z ]+), ([A-Za-z ]+)")
	cityState := cityStatePattern.FindStringSubmatch(station)

	if cityState != nil {
		station = cityState[2] + "/" + cityState[1]
		station = strings.Replace(station, " ", "_", -1)
	}
	return station
}

// BuildURL returns the URL required by the Weather Underground API
// from the query type, station id, and API key
func BuildURL(infoType string, stationId string) string {

	const URLstem = "http://api.wunderground.com/api/"
	const query = "/q/"
	const format = ".json"
	return URLstem + conf.Key + "/" + infoType + query + stationId + format
}

// Fetch does URL processing
func Fetch(url string) ([]byte, os.Error) {
	res, err := http.Get(url)
	CheckError(err)
	if res.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Bad HTTP Status: %d\n", res.StatusCode)
		return nil, err
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return b, err
}

// CheckError exits on error with a message
func CheckError(err os.Error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error\n%v\n", err)
		os.Exit(1)
	}
}

func init() {
	ReadConf()
}

type AlmanacConditions struct {
	Almanac Almanac
}

type Almanac struct {
	Temp_high Temp_high
	Temp_low  Temp_low
}

type Temp_high struct {
	Normal     Normal
	Record     Record
	Recordyear string
}

type Temp_low struct {
	Normal     Normal
	Record     Record
	Recordyear string
}

type Normal struct {
	F string
	C string
}

type Record struct {
	F string
	C string
}

// printAlmanac prints the Almanac for a given station to standard out
func printAlmanac(obs *AlmanacConditions, stationId string) {

	normalHighF := obs.Almanac.Temp_high.Normal.F
	normalHighC := obs.Almanac.Temp_high.Normal.C
	normalLowF := obs.Almanac.Temp_low.Normal.F
	normalLowC := obs.Almanac.Temp_low.Normal.C

	recordHighF := obs.Almanac.Temp_high.Record.F
	recordHighC := obs.Almanac.Temp_high.Record.C
	recordHYear := obs.Almanac.Temp_high.Recordyear
	recordLowF := obs.Almanac.Temp_low.Record.F
	recordLowC := obs.Almanac.Temp_low.Record.C
	recordLYear := obs.Almanac.Temp_low.Recordyear

	fmt.Printf("Normal high: %s\u00B0 F (%s\u00B0 C)\n", normalHighF, normalHighC)
	fmt.Printf("Record high: %s\u00B0 F (%s\u00B0 C) [%s]\n", recordHighF, recordHighC, recordHYear)
	fmt.Printf("Normal low : %s\u00B0 F (%s\u00B0 C)\n", normalLowF, normalLowC)
	fmt.Printf("Record low : %s\u00B0 F (%s\u00B0 C) [%s]\n", recordLowF, recordLowC, recordLYear)

}

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
func printAstro(obs *AstroConditions, stationId string) {

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

type AlertConditions struct {
	Alerts []Alerts
}

type Alerts struct {
	Date        string
	Expires     string
	Description string
	Message     string
}

// printAlerts prints the alerts for a given station to standard out
func printAlerts(obs *AlertConditions, stationId string) {
	if len(obs.Alerts) == 0 {
		fmt.Println("No active alerts")
	} else {
		fmt.Printf("Station: %s\n", stationId)
		for _, a := range obs.Alerts {
			fmt.Printf("### %s ###\n\nIssued at %s\nExpires at %s\n%s\n",
				a.Description, a.Date, a.Expires, a.Message)
		}
	}
}

type ForecastConditions struct {
	Forecast Forecast
}

type Forecast struct {
	Txt_forecast Txt_forecast
}

type Txt_forecast struct {
	Date        string
	Forecastday []Forecastday
}

type Forecastday struct {
	Title   string
	Fcttext string
}

// printForecast prints the forecast for a given station to standard out
func printForecast(obs *ForecastConditions, stationId string) {
	t := obs.Forecast.Txt_forecast
	fmt.Printf("Issued at %s\n", t.Date)
	for _, f := range t.Forecastday {
		fmt.Printf("%s: %s\n", f.Title, f.Fcttext)
	}
}

type Conditions struct {
	Current_observation Current
}

type Current struct {
	Observation_time     string
	Observation_location Location
	Station_id           string
	Weather              string
	Temperature_string   string
	Relative_humidity    string
	Wind_string          string
	Pressure_mb          string
	Pressure_in          string
	Pressure_trend       string
	Dewpoint_string      string
	Heat_index_string    string
	Windchill_string     string
	Visibility_mi        string
	Precip_today_string  string
}

type Location struct {
	Full string
}

// printConditions prints the conditions to standard output
func printConditions(obs *Conditions) {
	current := obs.Current_observation
	fmt.Printf("Current conditions at %s (%s)\n%s\n",
		current.Observation_location.Full, current.Station_id, current.Observation_time)
	fmt.Println("   Temperature:", current.Temperature_string)
	fmt.Println("   Sky Conditions:", current.Weather)
	fmt.Println("   Wind:", current.Wind_string)
	fmt.Println("   Relative humidity:", current.Relative_humidity)
	pstring := fmt.Sprintf("   Pressure: %s in (%s mb) and", current.Pressure_in, current.Pressure_mb)
	switch current.Pressure_trend {
	case "+":
		fmt.Println(pstring, "rising")
	case "-":
		fmt.Println(pstring, "falling")
	case "0":
		fmt.Println(pstring, "holding steady")
	}
	fmt.Println("   Dewpoint: ", current.Dewpoint_string)
	if current.Heat_index_string != "NA" {
		fmt.Println("   Heat Index: ", current.Heat_index_string)
	}
	if current.Windchill_string != "NA" {
		fmt.Println("   Windchill: ", current.Windchill_string)
	}
	fmt.Printf("   Visibility: %s miles\n", current.Visibility_mi)
	m, _ := regexp.MatchString("0.0", current.Precip_today_string)
	if !m {
		fmt.Println("   Precipitation today:", current.Precip_today_string)
	}
}

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
func printLookup(obs *Lookup) {
	station := obs.Location.Nearby_weather_stations.Airport.Station
	if len(station) == 0 {
		fmt.Println("No area stations")
	} else {
		for _, s := range station {
			fmt.Printf("%s: %s\n", s.City, s.Icao)
		}
	}
}

// weather prints various weather information for a specified station
func weather(operation string, station string) {
	url := BuildURL(operation, station)
	b, err := Fetch(url)
	CheckError(err)

	switch operation {
	case "almanac":
		var obs AlmanacConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		printAlmanac(&obs, station)
	case "astronomy":
		var obs AstroConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		printAstro(&obs, station)
	case "alerts":
		var obs AlertConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		printAlerts(&obs, station)
	case "conditions":
		var obs Conditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		printConditions(&obs)
	case "forecast":
		var obs ForecastConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		printForecast(&obs, station)
	case "geolookup":
		var l Lookup
		jsonErr := json.Unmarshal(b, &l)
		CheckError(jsonErr)
		printLookup(&l)
	}
}

func main() {
	stationId := Options()
	if doall {
		weather("alerts", stationId)
		weather("almanac", stationId)
		weather("astronomy", stationId)
		weather("conditions", stationId)
		weather("forecast", stationId)
		weather("geolookup", stationId)
		os.Exit(0)
	}
	if doalerts {
		weather("alerts", stationId)
	}
	if doalmanac {
		weather("almanac", stationId)
	}
	if doastro {
		weather("astronomy", stationId)
	}
	if doconditions {
		weather("conditions", stationId)
	}
	if doforecast {
		weather("forecast", stationId)
	}
	if dolookup {
		weather("geolookup", stationId)
	}
	if flag.NFlag() == 0 {
		weather("conditions", stationId)
	}
}
