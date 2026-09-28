package main

import (
	"fmt"
	// "errors"
)
// ---------------------------------------------------------------------------------------------------
// Lession 3
// Complete the getMonthlyPrice function. It accepts a tier (string) as input and returns the monthly price for that 
// tier in pennies. Here are the prices in dollars:

// "basic" - $100.00
// "premium" - $150.00
// "enterprise" - $500.00
// Convert the prices from dollars to pennies. If the given tier doesn't match any of the above, return 0 pennies.

// func getMonthlyPrice(tier string) int {
// 	switch tier {
// 	case "basic":
// 		return 10000
// 	case "premium":
// 		return 15000
// 	case "enterprise":
// 		return 50000
// 	default:
// 		return 0.00
// 	}
// }
// ---------------------------------------------------------------------------------------------------
// Lesson 5
// monthlyBillIncrease: Should return the increase in the bill from the previous to the current month. If the bill 
// decreased, return a negative number.
// getBillForMonth: Should return the total cost for the number of messages sent.
// Fix the bugs in the monthlyBillIncrease and getBillForMonth functions. Looks like whoever wrote the functions 
// didn't know the getBillForMonth function's bill parameter would be passed by value. It's not actually updating the 
// lastMonthBill and thisMonthBill variables as intended so monthlyBillIncrease isn't returning the right result.

// Drop the bill parameter from getBillForMonth, so it only takes 2 parameters.
// Instead, simply return the total cost of the messages.
// monthlyBillIncrease should use the result of calling getBillForMonth to calculate the increase between months.

// func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
// 	return getBillForMonth(costPerSend, numThisMonth) -
// 	getBillForMonth(costPerSend, numLastMonth)
// }

// func getBillForMonth(costPerSend, messagesSent int) int {
// 	return costPerSend * messagesSent
// }

// Lesson 6
// Assignment
// Run the code as-is. You should get a compiler error.
// Fix getProductMessage to ignore the unused return value.

// func getProductMessage(tier string) string {
// 	quantityMsg, priceMsg, _ := getProductInfo(tier)
// 	return "You get " + quantityMsg + " for " + priceMsg + "."
// }

// // don't touch below this line

// func getProductInfo(tier string) (string, string, string) {
// 	if tier == "basic" {
// 		return "1,000 texts per month", "$30 per month", "most popular"
// 	} else if tier == "premium" {
// 		return "50,000 texts per month", "$60 per month", "best value"
// 	} else if tier == "enterprise" {
// 		return "unlimited texts per month", "$100 per month", "customizable"
// 	} else {
// 		return "", "", ""
// 	}
// }
// ---------------------------------------------------------------------------------------------------
// Lesson 7
// Assignment
// One of our clients likes us to send text messages reminding users of life events coming up.

// Fix the bug by adding named return values to the function signature – the bare return at the end is already a 
// naked return that will return them. The variables need to be automatically initialized. Order them as they appear 
// in the code. Do not alter the body of the function.

// func yearsUntilEvents(age int) (yearsUntilAdult int,yearsUntilDrinking int,yearsUntilCarRental int) {
// 	// don't touch below this line

// 	yearsUntilAdult = 18 - age
// 	if yearsUntilAdult < 0 {
// 		yearsUntilAdult = 0
// 	}
// 	yearsUntilDrinking = 21 - age
// 	if yearsUntilDrinking < 0 {
// 		yearsUntilDrinking = 0
// 	}
// 	yearsUntilCarRental = 25 - age
// 	if yearsUntilCarRental < 0 {
// 		yearsUntilCarRental = 0
// 	}
// 	return
// }
// ---------------------------------------------------------------------------------------------------
// Lesson 8
// Fix the bug in the code so that it returns the named values explicitly.

// func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
// 	yearsUntilAdult = 18 - age
// 	if yearsUntilAdult < 0 {
// 		yearsUntilAdult = 0
// 	}
// 	yearsUntilDrinking = 21 - age
// 	if yearsUntilDrinking < 0 {
// 		yearsUntilDrinking = 0
// 	}
// 	yearsUntilCarRental = 25 - age
// 	if yearsUntilCarRental < 0 {
// 		yearsUntilCarRental = 0
// 	}
// 	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
// }

// ---------------------------------------------------------------------------------------------------
// Lesson 11
// Assignment
// Complete the reformat function. It takes a message string and a formatter function as input:

// Apply the given formatter three times to the message
// Add a prefix of TEXTIO: to the result
// Return the final string
// For example, if the message is "General Kenobi" and the given formatter adds a period to the end of the string, the final result should be

// TEXTIO: General Kenobi...

// func reformat(message string, formatter func(string) string) string {
// 	return "TEXTIO: " + formatter(formatter(formatter(message)))
// }
// ---------------------------------------------------------------------------------------------------

// Lesson 12
// Assignment
// Complete the printReports function. It takes as input a sequence of messages, intro, body, outro. It should call 
// printCostReport once for each message by making three separate calls (you don't need to use loops or arrays for this).

// For each call of printCostReport, give it an anonymous function that returns the cost of a message as an integer. 
// Here are the costs:

// Intro: 2x the message length
// Body: 3x the message length
// Outro: 4x the message length
// Use the built-in len() function to get the length of a string:

// helloLen := len("hello")
// // helloLen = 5

// func printReports(intro, body, outro string) {
// 	printCostReport(func(a string) int {
// 		return len(a)* 2
// 	}, intro)
// 	printCostReport(func(a string) int {
// 		return len(a)* 3
// 	}, body)
// 	printCostReport(func(a string) int {
// 		return len(a)* 4
// 	}, outro)
// }

// // don't touch below this line

// func main() {
// 	printReports(
// 		"Welcome to the Hotel California",
// 		"Such a lovely place",
// 		"Plenty of room at the Hotel California",
// 	)
// }

// func printCostReport(costCalculator func(string) int, message string) {
// 	cost := costCalculator(message)
// 	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
// 	fmt.Println()
// }

// ---------------------------------------------------------------------------------------------------
// Lesson 13
// Assignment
// Complete the bootup function.

// Be sure to print the following string just before the bootup function returns:

// TEXTIO BOOTUP DONE

// Use defer so that you only have to write this message once instead of before each return statement. 
// The message should be printed on its own newline.

// func bootup() {
// 	defer fmt.Println("TEXTIO BOOTUP DONE")
// 	ok := connectToDB()
// 	if !ok {
// 		return
// 	}
// 	ok = connectToPaymentProvider()
// 	if !ok {
// 		return
// 	}
// 	fmt.Println("All systems ready!")
// }

// // don't touch below this line

// var shouldConnectToDB = true

// func connectToDB() bool {
// 	fmt.Println("Connecting to database...")
// 	if shouldConnectToDB {
// 		fmt.Println("Connected!")
// 		return true
// 	}
// 	fmt.Println("Connection failed")
// 	return false
// }

// var shouldConnectToPaymentProvider = true

// func connectToPaymentProvider() bool {
// 	fmt.Println("Connecting to payment provider...")
// 	if shouldConnectToPaymentProvider {
// 		fmt.Println("Connected!")
// 		return true
// 	}
// 	fmt.Println("Connection failed")
// 	return false
// }

// func test(dbSuccess, paymentSuccess bool) {
// 	shouldConnectToDB = dbSuccess
// 	shouldConnectToPaymentProvider = paymentSuccess
// 	bootup()
// 	fmt.Println("====================================")
// }

// func main() {
// 	test(true, true)
// 	test(false, true)
// 	test(true, false)
// 	test(false, false)
// }

// ---------------------------------------------------------------------------------------------------

// Lesson 14
// Assignment
// Run the code without changing anything: you should see a compilation error.
// Fix the scoping issue in the function so that it runs as you'd expect.

// func splitEmail(email string) (string, string) {
	
// 	username, domain := "", ""

// 	for i, r := range email {
// 		if r == '@' {
// 			username = email[:i]
// 			domain = email[i+1:]
// 			break
// 		}
// 	}
// 	return username, domain
// }

// ---------------------------------------------------------------------------------------------------
// Lesson 15

// Assignment
// Complete the placeOrder function.

// It returns a bool indicating whether the order was successful (true is a success) and a float64 representing the user's 
// balance after the order. The placeOrder function should always return the account balance regardless of whether it was 
// adjusted.

// The amountInStock and calcPrice functions can be used to look up the current stock and price of an item.

// If the quantity is greater than the amount in stock, the order should be rejected.
// If the user doesn't have enough money in their account, the order should be rejected.
// Otherwise, the order should be accepted and you should return the new balance.

// func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
// 	if quantity > amountInStock(productID) {
// 		return false, accountBalance
// 	} else if calcPrice(productID, quantity) > accountBalance {
// 		return false, accountBalance
// 	} else {
// 		return true, accountBalance - calcPrice(productID, quantity)
// 	}
// }

// // Don't touch below this line

// func calcPrice(productID string, quantity int) float64 {
// 	return priceList(productID) * float64(quantity)
// }

// func priceList(productID string) float64 {
// 	if productID == "1" {
// 		return 1.50
// 	} else if productID == "2" {
// 		return 2.25
// 	} else if productID == "3" {
// 		return 3.00
// 	} else if productID == "4" {
// 		return 1.00
// 	} else if productID == "5" {
// 		return 2.50
// 	} else if productID == "6" {
// 		return 8.99
// 	} else if productID == "7" {
// 		return 22.50
// 	} else if productID == "8" {
// 		return 50.00
// 	} else if productID == "9" {
// 		return 999.99
// 	} else {
// 		return 0.00
// 	}
// }

// func amountInStock(productID string) int {
// 	if productID == "1" {
// 		return 11
// 	} else if productID == "2" {
// 		return 25
// 	} else if productID == "3" {
// 		return 4
// 	} else if productID == "4" {
// 		return 6
// 	} else if productID == "5" {
// 		return 50
// 	} else if productID == "6" {
// 		return 2
// 	} else if productID == "7" {
// 		return 0
// 	} else if productID == "8" {
// 		return 99
// 	} else if productID == "9" {
// 		return 1
// 	} else {
// 		return 0
// 	}
// }

// ---------------------------------------------------------------------------------------------------

// Lesson 16
// Assignment
// Keeping track of how many texts we send is mission-critical at Textio. Complete the adder() enclosing function.

// Create an enclosed sum value inside the adder() function.
// Return a function from the adder() function that adds its input (an int) to the sum and returns the new value of sum. 
// (In other words, it keeps a running total of the sum variable within a closure.)

// func adder() func(int) int {
// 	sum := 0
// 	return func(ad int) int {
// 		sum += ad
// 		return sum
// 	}
// }

// ---------------------------------------------------------------------------------------------------

// Assignment
// The Textio API needs a very robust error-logging system so we can see when things are going awry in the back-end 
// system. We need a function that can create a custom "logger" (a function that prints to the console) given a 
// specific formatter.

// These errors are test data, not runtime failures.

// Complete the getLogger function. It should take as input a formatter function and return a new function. The 
// new logger function takes as input two strings and passes them to the formatter, then prints the result. Keep 
// the order of the strings.

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
// func getLogger(formatter func(string, string) string) func(string, string) {
// 	return func(a string, b string) {
// 		fmt.Println(formatter(a, b))
// 	}
// }

// // don't touch below this line

// func test(first string, errors []error, formatter func(string, string) string) {
// 	defer fmt.Println("====================================")
// 	logger := getLogger(formatter)
// 	fmt.Println("Logs:")
// 	for _, err := range errors {
// 		logger(first, err.Error())
// 	}
// }

// func colonDelimit(first, second string) string {
// 	return first + ": " + second
// }

// func commaDelimit(first, second string) string {
// 	return first + ", " + second
// }

// func main() {
// 	dbErrors := []error{
// 		errors.New("out of memory"),
// 		errors.New("cpu is pegged"),
// 		errors.New("networking issue"),
// 		errors.New("invalid syntax"),
// 	}
// 	test("Error on database server", dbErrors, colonDelimit)

// 	mailErrors := []error{
// 		errors.New("email too large"),
// 		errors.New("non alphanumeric symbols found"),
// 	}
// 	test("Error on mail server", mailErrors, commaDelimit)
// }

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------
