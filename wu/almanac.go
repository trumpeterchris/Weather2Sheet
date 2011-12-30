
package almanac

import (
  "fmt"
)

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
func PrintAlmanac(obs *AlmanacConditions, stationId string) {

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
