
Wu
==========

Version 3.2.0.


_wu_ is a fast small command-line application that retrieves weather data from [Weather Underground](http://www.wunderground.com).

Description
-----------

To use _wu,_ you need to obtain an API key from Weather Underground [http://www.wunderground.com/weather/api/](http://www.wunderground.com/weather/api/).  You should then add that key and the name of your default weather station to $HOME/.condrc:

	{
	  "key": "YOUR_API_KEY",
	  "station": "Lincoln, NE"
	}

(the above is available in the wu root directory as "condrc")

wu has the following major options:

* _--conditions_ reports the current weather conditions.

* _--forecast_ gives the current forecast.

* _--alerts_ reports any active weather alerts.

* _--lookup_ [STATION] allows you to determine the codes for the various weather stations in a particular area.  The format for STATION is the same as that for the -s switch below.

* _--astronomy_ reports sunrise, sunset, and lunar phase.

* _--almanac_ reports average high and low temperatures, as well as record temperatures for the day.

* _--yesterday_ reports detailed alamanac information for the previous day.

* _--all_ generate all reports (useful for creating custom reports and for mollifying the truly weather-crazed).
	
All six options can be accompanied by the -s switch, which can be used to override the default location in .condrc.  The argument passed to -s can be a "city, state-abbreviation/country", a (U.S. or Canadian) zip code, a 3- or 4-letter airport code, or "lat,long").

wu also has two additional switches that provide information about the program:

* -h help
* -V version

Installing Wu 
-----------

The easiest way to install wu (assuming you have both [Git](http://git-scm.com/) and a Go compiler) is to type:

  GOPATH=[PATH] goinstall -u github.com/sramsay/wu/wu

where [PATH] is the directory you'd like it installed (e.g. /usr/local/bin).

If you don't have a Go compiler, you'll need to install one.  Detailed instructions are [here](http://golang.org/doc/install.html).  But in brief:

	hg clone -u release https://go.googlecode.com/hg/ go
	cd go/src
	./all.bash
	export GOROOT=/path/to/go
	export GOARCH=amd64
	export GOOS=linux
	export PATH=${GOROOT}/bin:$PATH

(substituting 386 for amd64, and darwin|freebsd for linux as appropriate).

Building wu from Source
-----------------------

To obtain the source code for wu:

  git clone git@github.com:sramsay/wu.git

To compile:

	cd wu/wu
	make
  GOPATH=/usr/local/bin make install

wu should work on any system that can compile Go programs.

License(s)
---------

Wu is written and maintained by [Stephen Ramsay](http://lenz.unl.edu/) (sramsay{dot}unl{at}gmail{dot}com) and [Anthony Starks](http://mindchunk.blogspot.com/).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see [http://www.gnu.org/licenses/](http://www.gnu.org/licenses/).

Data courtesy of Weather Underground, Inc. (WUI) is subject to the [Weather Underground API Terms and Conditions of Use](http://www.wunderground.com/weather/api/d/terms.html).  The author of this software is not affiliated with WUI, and the software is neither sponsored nor endorsed by WUI.

Thanks
------

Wu was heavily inspired by Jeremy Stanley's [weather](http://fungi.yuggoth.org/weather/).  This is a lovely Python script that has more-or-less the same output format as wu.  I reimplemented the system because Stanley's had stopped working (for me) and I wanted a program that was faster.  I also wanted a system that takes advantage of Weather Underground's rich, [JSON](http://www.json.org/) API.
