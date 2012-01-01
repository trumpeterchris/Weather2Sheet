package yesterday

import "fmt"

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

func PrintYesterday(obs *YesterdayConditions, stationId string) {
	history := obs.History.Dailysummary[0]
	fmt.Print("Weather summary for yesterday: ")
	if history.Fog == "1" {
		fmt.Print("fog ")
	}
	if history.Rain == "1" {
		fmt.Print("rain ")
	}
	if history.Snow == "1" {
		fmt.Print("snow ")
	}
	if history.Hail == "1" {
		fmt.Print("hail ")
	}
	if history.Tornado == "1" {
		fmt.Print("tornado ")
	}
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
		fmt.Print("      Heating Degree Days: " + history.Heatingdegreedays)
		if history.Heatingdegreedaysnormal != "" {
			fmt.Printf(" (%s days normal)\n", history.Heatingdegreedaysnormal)
		}
		if history.Heatingdegreedaysnormal != "" && history.Heatingdegreedaysnormal != "0" {
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

	if history.Coolingdegreedaysnormal != "" && history.Coolingdegreedaysnormal != "0" {
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
