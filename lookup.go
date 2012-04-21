/*
* lookup.go
*
* This file is part of wu.  It contains functions related to
* the -lookup switch (station lookup).
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Sat Apr 21 14:39:18 CDT 2012
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
