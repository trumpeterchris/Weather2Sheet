/*
* conditions.go
*
* This file is part of wu.  It contains functions related to
* the -conditions switch (current conditions).
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Sun Jan 08 16:47:00 CST 2012
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

package conditions

import (
	"fmt"
	"regexp"
)

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
func PrintConditions(obs *Conditions) {
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
