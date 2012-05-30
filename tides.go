/*
* tide.go
*
* This file is part of wu.  It contains functions related to
* the -tide switch (high and low tide data).
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Mon May 28 12:17:30 CDT 2012
*
* Copyright © 2010-2012 by Stephen Ramsay and Anthony Starks.
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
  "os"
  "strconv"
  "time"
)

type TideConditions struct {
  Tide Tide
}

type Tide struct {
  Tideinfo    []Tideinfo
  Tidesummary []Tidesummary
}

type Tideinfo struct {
  Tidesite string
}

type Tidesummary struct {
  Date Date // Defined in wu.go
  Data Data
}

type Data struct {
  Height string
  Type   string
}

// printTide prints the tidal data for given station to standard out
func PrintTides(obs *TideConditions, stationID string) {
  tide := obs.Tide
  info := tide.Tideinfo
  summary := tide.Tidesummary

  if len(summary) == 0 {
    fmt.Println("No tidal data available.")
    os.Exit(0)
  }

  fmt.Printf("Tidal data for %s\n", info[0].Tidesite)

  day := time.Now().Day()
  month := time.Now().Month()
  year := time.Now().Year()

  for c := 1; c <= 4; c++ {
    fmt.Printf("%d/%d/%d:\n", month, day, year)
    for _, s := range summary {
      if s.Date.Mday == strconv.Itoa(day) {
        if s.Data.Type == "Low Tide" || s.Data.Type == "High Tide" {
          fmt.Printf("  %s: %s:%s\n", s.Data.Type, s.Date.Hour, s.Date.Min)
        }
      }
    }
    // Increment dates
    if month == 12 && day < time.Now().AddDate(0, 0, c).Day() { // year change
      day = time.Now().AddDate(0, 0, c).Day()
      month = time.Now().AddDate(0, 1, 0).Month()
      year = time.Now().AddDate(1, 0, 0).Year()
    } else if day < time.Now().AddDate(0, 0, c).Day() { // month change
      day = time.Now().AddDate(0, 0, c).Day()
      month = time.Now().AddDate(0, 1, 0).Month()
    } else { // day change
      day = time.Now().AddDate(0, 0, c).Day()
    }
  }
}
