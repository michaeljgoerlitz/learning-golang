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

	openRateInt := int(averageOpenRate)

	fmt.Println(openRateInt)

	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	fmt.Println("Plan: ", premiumPlanName)
	fmt.Println("Plan: ", basicPlanName)

	fmt.Printf("I am %v years old", 10)
	// I am 10 years old

	fmt.Printf("I am %v years old", "way too many")
	// I am way too many years old

	fmt.Printf("I am %s years old", "way too many")
	// I am way too many years old

	fmt.Printf("I am %d years old", 10)
	// I am 10 years old

	fmt.Printf("I am %f years old", 10.523)
	// I am 10.523000 years old

	// The ".2" rounds the number to 2 decimal places
	fmt.Printf("I am %.2f years old", 10.523)
	// I am 10.53 years old
}
