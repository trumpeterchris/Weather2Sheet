/*
* utils.go
*
* This file is part of wu.  It contains utility functions.
*
* Written and maintained by Stephen Ramsay <sramsay.unl@gmail.com>
* and Anthony Starks.
*
* Last Modified: Mon Aug 01 12:25:26 CDT 2016
*
* Copyright © 2010-2016 by Stephen Ramsay and Anthony Starks.
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
	"regexp"
)

func Convert(temp string) string {
	celsiusPattern := regexp.MustCompile("([0-9]+ F) \\(([0-9]+ C)\\)")

	pattern := celsiusPattern.FindStringSubmatch(temp)
	newTemp := pattern[2] + " " + "(" + pattern[1] + ")"

	return newTemp
}
