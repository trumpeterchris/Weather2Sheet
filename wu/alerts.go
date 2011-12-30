
package alerts

import (
  "fmt"
)

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
func PrintAlerts(obs *AlertConditions, stationId string) {
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
