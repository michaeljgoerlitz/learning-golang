package variables

import "fmt"

func SetVariables() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	empty := ""

	fmt.Printf(
		"%v %f %v %q %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
		empty,
	)

	penniesPerText := 2

	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	averageOpenRate, displayMessage := .23, "is the average open rate."

	fmt.Println(averageOpenRate, displayMessage)
}
