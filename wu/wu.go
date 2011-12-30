/*
* wu - a small, fast command-line application for retrieving weather
* data from Weather Underground
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Wed Dec 28 11:03:24 CST 2011
*
* Copyright © 2010-2011 by Stephen Ramsay and Anthony Starks.
*
* wu is free software; you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation; either version 3, or (at your option)
* any later version.
*
* wu is distributed in the hope that it will be useful, but WITHOUT
* ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
* or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public
* License for more details.
*
* You should have received a copy of the GNU General Public License
* along with wu; see the file COPYING.  If not see
* <http://www.gnu.org/licenses/>.
 */

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
	help, version, doall, doalmanac, doalerts, doconditions, dolookup, doforecast, doastro, doyesterday bool
	conf                                                                                                Config
)

const defaultStation = "KLNK"

// GetVersion returns the version of the package
func GetVersion() string {
	return "3.2.0"
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

	flag.BoolVar(&doconditions, "conditions", false, "Reports the current weather conditions")
	flag.BoolVar(&doalerts, "alerts", false, "Reports any active weather alerts")
	flag.BoolVar(&dolookup, "lookup", false, "Lookup the codes for the weather stations in a particular area")
	flag.BoolVar(&doastro, "astro", false, "Reports sunrise, sunset, and lunar phase")
	flag.BoolVar(&doforecast, "forecast", false, "Reports the current forecast")
	flag.BoolVar(&doalmanac, "almanac", false, "Reports average high, low and record temperatures")
	flag.BoolVar(&doyesterday, "yesterday", false, "Reports yesterday's weather data")
	flag.BoolVar(&help, "h", false, "Print this message")
	flag.BoolVar(&version, "V", false, "Print the version number")
	flag.BoolVar(&doall, "all", false, "Show all weather data")
	flag.StringVar(&station, "s", sconf,
		"Weather station: \"city, state-abbreviation\", (US or Canadian) zipcode, 3- or 4-letter airport code, or LAT,LONG")
	flag.Parse()

	// Check for correct usage of wu -lookup
	if dolookup {
		if len(os.Args) == 3 {
			station = os.Args[len(os.Args)-1]
		} else {
			fmt.Println("Usage: wu -lookup [station] where station is a \"city, state-abbreviation\", (US or Canadian) zipcode, 3- or 4-letter airport code, or LAT,LONG")
			os.Exit(0)
		}
	}

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println("conditions " + GetVersion())
		fmt.Println("Copyright  2011 by Stephen Ramsay and")
		fmt.Println("Anthony Starks. Data courtesy of Weather")
		fmt.Println("Underground, Inc. is subject to Weather")
		fmt.Println("Underground Data Feed Terms of Service.")
		fmt.Println("The program itself is free software, and")
		fmt.Println("you are welcome to redistribute it under")
		fmt.Println("certain conditions.  See LICENSE for details.")
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
		fmt.Println("   Precipitation today: ", current.Precip_today_string)
	}
}

//////////////////////

type YesterdayConditions struct {
	History History
}

type History struct {
	Dailysummary []Dailysummary
}

type Dailysummary struct {
	Fog                                string
	Rain                               string
	Snow                               string
	Snowfallm                          string
	Snowfalli                          string
	Monthtodatesnowfallm               string
	Monthtodatesnowfalli               string
	Since1julsnowfallm                 string
	Since1julsnowfalli                 string
	Snowdepthm                         string
	Snowdepthi                         string
	Hail                               string
	Thunder                            string
	Tornado                            string
	Meantempm                          string
	Meantempi                          string
	Meandewptm                         string
	Meandewpti                         string
	Meanpressurem                      string
	Meanpressurei                      string
	Meanwindspdm                       string
	Meanwindspdi                       string
	Meanwdire                          string
	Meanwdird                          string
	Meanvism                           string
	Meanvisi                           string
	Humidity                           string
	Maxtempm                           string
	Maxtempi                           string
	Mintempm                           string
	Mintempi                           string
	Maxhumidity                        string
	Minhumidity                        string
	Maxdewptm                          string
	Maxdewpti                          string
	Mindewptm                          string
	Mindewpti                          string
	Maxpressurem                       string
	Maxpressurei                       string
	Minpressurem                       string
	Minpressurei                       string
	Maxwspdm                           string
	Maxwspdi                           string
	Minwspdm                           string
	Minwspdi                           string
	Maxvism                            string
	Maxvisi                            string
	Minvism                            string
	Minvisi                            string
	Gdegreedays                        string
	Heatingdegreedays                  string
	Coolingdegreedays                  string
	Precipm                            string
	Precipi                            string
	Heatingdegreedaysnormal            string
	Monthtodateheatingdegreedays       string
	Monthtodateheatingdegreedaysnormal string
	Since1sepheatingdegreedays         string
	Since1sepheatingdegreedaysnormal   string
	Since1julheatingdegreedays         string
	Since1julheatingdegreedaysnormal   string
	Coolingdegreedaysnormal            string
	Monthtodatecoolingdegreedays       string
	Monthtodatecoolingdegreedaysnormal string
	Since1sepcoolingdegreedays         string
	Since1sepcoolingdegreedaysnormal   string
	Since1jancoolingdegreedays         string
	Since1jancoolingdegreedaysnormal   string
}

func printYesterday(obs *YesterdayConditions, stationId string) {
  history := obs.History.Dailysummary[0]
  fmt.Print("Weather summary for yesterday: ")
  if history.Fog == "1" {fmt.Print("fog ")}
  if history.Rain == "1" {fmt.Print("rain ")}
  if history.Snow == "1" {fmt.Print("snow ")}
  if history.Hail == "1" {fmt.Print("hail ")}
  if history.Tornado == "1" {fmt.Print("tornado ")}
  fmt.Print("\n")

  // if "month to date" is nil, it likely means that the station
  // doesn't report full almanac information (which is frequently
  // the case for non-U.S (NWS) station sources.  This may be the
  // case for several measurements in this section.

  // Snow

  if history.Snow == "1" && history.Monthtodatesnowfalli != "" {
    fmt.Println("   Snow:")
    if history.Snowfalli == "T" {
      fmt.Println("     Yesterday: trace")
    } else if history.Snowfalli != "" {
      fmt.Printf("     Yesterday: %s in (%s mm)\n", history.Snowfalli, history.Snowfallm)
    }
    fmt.Printf("     Snow depth: %s in (%s mm)\n", history.Snowdepthi, history.Snowdepthm)
    fmt.Printf("     Month to date: %s in (%s mm)\n", history.Monthtodatesnowfalli, history.Monthtodatesnowfallm)
    fmt.Printf("     Since July 1st: %s in (%s mm)\n", history.Since1julsnowfalli, history.Since1julsnowfallm)
  }

  // Precipitation

  if history.Rain == "1" {
    fmt.Printf("   Precipitation: %s in (%s mm)\n", history.Precipi, history.Precipm)
  }

  // Temperature

  fmt.Println("   Temperature:")
  fmt.Printf("      Mean Temperature: %s F (%s C)\n", history.Meantempi, history.Meantempm)
  fmt.Printf("      Max Temperature: %s F (%s C)\n", history.Maxtempi, history.Maxtempm)
  fmt.Printf("      Min Temperature: %s F (%s C)\n", history.Mintempi, history.Mintempm)

  // Degree Days

  fmt.Println("   Degree Days:")
  if history.Heatingdegreedays != "" {
    fmt.Println("      Heating Degree Days: " + history.Heatingdegreedays)
    if history.Heatingdegreedaysnormal != "" {
      fmt.Printf(" (%s days normal)\n", history.Heatingdegreedaysnormal)
    }
    if history.Heatingdegreedaysnormal != "" {
      fmt.Printf("         HDG month to date: %s (%s days normal)\n", history.Monthtodateheatingdegreedays, history.Monthtodateheatingdegreedaysnormal)
      if history.Since1julheatingdegreedaysnormal == "" {
        fmt.Printf("         HDG since Sept 1st: %s (%s days normal)\n", history.Since1sepheatingdegreedays, history.Since1sepheatingdegreedaysnormal)
      } else {
        fmt.Printf("         HDG since July 1st: %s (%s days normal)\n", history.Since1julheatingdegreedays, history.Since1julheatingdegreedaysnormal)
      }
    } else {
      fmt.Print("\n")
    }
  }
  if history.Coolingdegreedaysnormal != "" {
    fmt.Print("      Cooling Degree Days: " + history.Coolingdegreedays)
    if history.Coolingdegreedaysnormal != "" {
      fmt.Printf(" (%s days normal)\n", history.Coolingdegreedaysnormal)
    } else {
      fmt.Print("\n")
    }
    if history.Coolingdegreedaysnormal != "" {
      fmt.Printf("         CDG month to date: %s (%s days normal)\n", history.Monthtodatecoolingdegreedays, history.Monthtodatecoolingdegreedaysnormal)
      if history.Since1jancoolingdegreedaysnormal == "" {
        fmt.Printf("         CDG since Sept 1st: %s (%s days normal)\n", history.Since1sepcoolingdegreedays, history.Since1sepcoolingdegreedaysnormal)
      } else {
        fmt.Printf("         CDG since Jan 1st: %s (%s days normal)\n", history.Since1jancoolingdegreedays, history.Since1jancoolingdegreedaysnormal)
      }
    } else {
      fmt.Print("\n")
    }
  }

  // Moisture

  fmt.Println("   Moisture:")
  fmt.Printf("      Mean Dew Point: %s (%s C)\n", history.Meandewpti, history.Meandewptm)
  fmt.Printf("      Max Dew Point: %s (%s C)\n", history.Maxdewpti, history.Maxdewptm)
  fmt.Printf("      Min Dew Point: %s (%s C)\n", history.Mindewpti, history.Mindewptm)
  if history.Humidity != "" {
    fmt.Printf("      Humidity: %s%%\n", history.Humidity)
  }
  fmt.Printf("      Max Humidity: %s%%\n", history.Maxhumidity)
  fmt.Printf("      Min Humidity: %s%%\n", history.Minhumidity)

  // Pressure

  fmt.Println("   Pressure:")
  fmt.Printf("      Mean Pressure: %s in (%s mb)\n", history.Meanpressurei, history.Meanpressurem)
  fmt.Printf("      Max Pressure: %s in (%s mb)\n", history.Maxpressurei, history.Maxpressurem)
  fmt.Printf("      Min Pressure: %s in (%s mb)\n", history.Minpressurei, history.Minpressurem)

  // Wind

  fmt.Println("   Wind:")
  fmt.Printf("      Mean Wind Speed: %s mph (%s kph)\n", history.Meanwindspdi, history.Meanwindspdm)
  fmt.Printf("      Max Wind Speed: %s mph (%s kph)\n", history.Maxwspdi, history.Maxwspdm)
  fmt.Printf("      Min Wind Speed: %s mph (%s kph)\n", history.Minwspdi, history.Minwspdm)
  fmt.Printf("      Mean Wind Direction: %s°\n", history.Meanwdird)

  // Visibility

  fmt.Println("   Visibility:")
  fmt.Printf("      Mean Visibility %s mi (%s km)\n", history.Meanvisi, history.Meanvism)
  fmt.Printf("      Max Visibility %s mi (%s km)\n", history.Maxvisi, history.Maxvism)
  fmt.Printf("      Min Visibility %s mi (%s km)\n", history.Minvisi, history.Minvism)

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
	case "yesterday":
		var obs YesterdayConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		printYesterday(&obs, station)
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
		weather("conditions", stationId)
		weather("forecast", stationId)
		weather("alerts", stationId)
		weather("almanac", stationId)
		weather("yesterday", stationId)
		weather("astronomy", stationId)
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
	if doyesterday {
		weather("yesterday", stationId)
	}
	if dolookup {
		weather("geolookup", stationId)
	}
	if flag.NFlag() == 0 {
		weather("conditions", stationId)
	}
}
