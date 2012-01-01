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
	"json"
	"http"
	"flag"
	"./alerts"
	"./almanac"
	"./astro"
	"./conditions"
	"./forecast"
	"./lookup"
	"./yesterday"
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
	return "3.2.2"
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

// weather prints various weather information for a specified station
func weather(operation string, station string) {
	url := BuildURL(operation, station)
	b, err := Fetch(url)
	CheckError(err)

	switch operation {
	case "almanac":
		var obs almanac.AlmanacConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		almanac.PrintAlmanac(&obs, station)
	case "astronomy":
		var obs astro.AstroConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		astro.PrintAstro(&obs, station)
	case "alerts":
		var obs alerts.AlertConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		alerts.PrintAlerts(&obs, station)
	case "conditions":
		var obs conditions.Conditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		conditions.PrintConditions(&obs)
	case "forecast":
		var obs forecast.ForecastConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		forecast.PrintForecast(&obs, station)
	case "yesterday":
		var obs yesterday.YesterdayConditions
		jsonErr := json.Unmarshal(b, &obs)
		CheckError(jsonErr)
		yesterday.PrintYesterday(&obs, station)
	case "geolookup":
		var l lookup.Lookup
		jsonErr := json.Unmarshal(b, &l)
		CheckError(jsonErr)
		lookup.PrintLookup(&l)
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
