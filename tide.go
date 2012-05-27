/*
* tide.go
*
* This file is part of wu.  It contains functions related to
* the -tide switch (high and low tide data).
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Sun May 27 16:39:29 CDT 2012
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

import "fmt"

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
	Date TideDate
	Data Data
}

type TideDate struct {
	Pretty string
}

type Data struct {
	Height string
	Type   string
}

func PrintTide(obs *TideConditions, stationID string) {
	fmt.Printf("Tidal data for %s\n", obs.Tide.Tideinfo[0].Tidesite)
}
