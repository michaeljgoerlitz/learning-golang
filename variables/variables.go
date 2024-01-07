package variables

import "fmt"

func SetVariables() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)
}
