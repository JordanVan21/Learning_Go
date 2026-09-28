package main

import (
	"fmt"
	// "slices"
	"testing"
)
// --------------------------------------------------------------------------------------------------------
// Lesson 3

// func Test(t *testing.T) {
// 	type testCase struct {
// 		tier     string
// 		expected int
// 	}

// 	runCases := []testCase{
// 		{"basic", 10000},
// 		{"premium", 15000},
// 		{"enterprise", 50000},
// 	}

// 	submitCases := append(runCases, []testCase{
// 		{"invalid", 0},
// 		{"", 0},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}

// 	skipped := len(submitCases) - len(testCases)

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range testCases {
// 		output := getMonthlyPrice(test.tier)
// 		if output != test.expected {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v)
// Expecting:  %v
// Actual:     %v
// Fail
// `, test.tier, test.expected, output)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v)
// Expecting:  %v
// Actual:     %v
// Pass
// `, test.tier, test.expected, output)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true
// --------------------------------------------------------------------------------------------------------
// Lesson 5
// func Test(t *testing.T) {
// 	type testCase struct {
// 		costPerSend  int
// 		numLastMonth int
// 		numThisMonth int
// 		expected     int
// 	}

// 	runCases := []testCase{
// 		{2, 89, 102, 26},
// 		{2, 98, 104, 12},
// 	}

// 	submitCases := append(runCases, []testCase{
// 		{3, 50, 40, -30},
// 		{3, 60, 60, 0},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}
// 	skipped := len(submitCases) - len(testCases)

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range testCases {
// 		output := monthlyBillIncrease(test.costPerSend, test.numLastMonth, test.numThisMonth)
// 		_ = getBillForMonth(0, 0)
// 		if output != test.expected {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v, %v, %v)
// Expecting:  %v
// Actual:     %v
// Fail
// `, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v, %v, %v)
// Expecting:  %v
// Actual:     %v
// Pass
// `, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true
// --------------------------------------------------------------------------------------------------------
// Lesson 6
// func Test(t *testing.T) {
// 	type testCase struct {
// 		tier     string
// 		expected string
// 	}
// 	runCases := []testCase{
// 		{"basic", "You get 1,000 texts per month for $30 per month."},
// 		{"premium", "You get 50,000 texts per month for $60 per month."},
// 	}
// 	submitCases := append(runCases, []testCase{
// 		{"enterprise", "You get unlimited texts per month for $100 per month."},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}
// 	skipped := len(submitCases) - len(testCases)

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range testCases {
// 		output := getProductMessage(test.tier)
// 		if output != test.expected {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v)
// Expecting:  %v
// Actual:     %v
// Fail
// `, test.tier, test.expected, output)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v)
// Expecting:  %v
// Actual:     %v
// Pass
// `, test.tier, test.expected, output)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true

// func Test(t *testing.T) {
// 	type testCase struct {
// 		age                   int
// 		exYearsUntilAdult     int
// 		exYearsUntilDrinking  int
// 		exYearsUntilCarRental int
// 	}

// 	runCases := []testCase{
// 		{4, 14, 17, 21},
// 		{18, 0, 3, 7},
// 		{22, 0, 0, 3},
// 	}

// 	submitCases := append(runCases, []testCase{
// 		{25, 0, 0, 0},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}

// 	skipped := len(submitCases) - len(testCases)

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range testCases {
// 		yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(test.age)
// 		if yearsUntilAdult != test.exYearsUntilAdult ||
// 			yearsUntilDrinking != test.exYearsUntilDrinking ||
// 			yearsUntilCarRental != test.exYearsUntilCarRental {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v)
// Expecting:  (%v, %v, %v)
// Actual:     (%v, %v, %v)
// Fail
// `, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v)
// Expecting:  (%v, %v, %v)
// Actual:     (%v, %v, %v)
// Pass
// `, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true
// --------------------------------------------------------------------------------------------------------
// Lesson 8
// func Test(t *testing.T) {
// 	type testCase struct {
// 		age                   int
// 		exYearsUntilAdult     int
// 		exYearsUntilDrinking  int
// 		exYearsUntilCarRental int
// 	}
// 	runCases := []testCase{
// 		{4, 14, 17, 21},
// 		{18, 0, 3, 7},
// 		{22, 0, 0, 3},
// 	}
// 	submitCases := append(runCases, []testCase{
// 		{25, 0, 0, 0},
// 		{35, 0, 0, 0},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}

// 	skipped := len(submitCases) - len(testCases)
// 	passCount := 0
// 	failCount := 0

// 	for _, test := range testCases {
// 		yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(test.age)
// 		if yearsUntilAdult != test.exYearsUntilAdult ||
// 			yearsUntilDrinking != test.exYearsUntilDrinking ||
// 			yearsUntilCarRental != test.exYearsUntilCarRental {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v)
// Expecting:  (%v, %v, %v)
// Actual:     (%v, %v, %v)
// Fail
// `, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v)
// Expecting:  (%v, %v, %v)
// Actual:     (%v, %v, %v)
// Pass
// `, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true

// ---------------------------------------------------------------------------------------------------
// Lesson 11

// func Test(t *testing.T) {
// 	type testCase struct {
// 		message       string
// 		formatter     func(string) string
// 		formatterName string
// 		expected      string
// 	}

// 	runCases := []testCase{
// 		{"hello", addExclamation, "addExclamation", "TEXTIO: hello!!!"},
// 		{"hello there", addPeriod, "addPeriod", "TEXTIO: hello there..."},
// 	}

// 	submitCases := append(runCases, []testCase{
// 		{"moor der ehT", reverseString, "reverseString", "TEXTIO: The red room"},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}
// 	skipped := len(submitCases) - len(testCases)

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range testCases {
// 		output := reformat(test.message, test.formatter)
// 		if output != test.expected {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v, %v)
// Expecting:  %v
// Actual:     %v
// Fail
// `, test.message, test.formatterName, test.expected, output)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v, %v)
// Expecting:  %v
// Actual:     %v
// Pass
// `, test.message, test.formatterName, test.expected, output)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// func addPeriod(s string) string {
// 	return s + "."
// }

// func addExclamation(s string) string {
// 	return s + "!"
// }

// func reverseString(s string) string {
// 	r := []rune(s)
// 	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true

// ---------------------------------------------------------------------------------------------------
// Lesson 14

// func TestSplitEmail(t *testing.T) {
// 	type testCase struct {
// 		email    string
// 		username string
// 		domain   string
// 	}

// 	runCases := []testCase{
// 		{"drogon@dragonstone.com", "drogon", "dragonstone.com"},
// 		{"rhaenyra@targaryen.com", "rhaenyra", "targaryen.com"},
// 	}

// 	submitCases := append(runCases, []testCase{
// 		{"viserys@kingslanding.com", "viserys", "kingslanding.com"},
// 		{"aegon@stormsend.com", "aegon", "stormsend.com"},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}

// 	passCount := 0
// 	failCount := 0
// 	skipped := len(submitCases) - len(testCases)

// 	for _, test := range testCases {
// 		username, domain := splitEmail(test.email)
// 		if username != test.username || domain != test.domain {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v)
// Expecting:  (%v, %v)
// Actual:     (%v, %v)
// Fail
// `, test.email, test.username, test.domain, username, domain)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v)
// Expecting:  (%v, %v)
// Actual:     (%v, %v)
// Pass
// `, test.email, test.username, test.domain, username, domain)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true

// ---------------------------------------------------------------------------------------------------

// Lesson 15

// func Test(t *testing.T) {
// 	type testCase struct {
// 		productID      string
// 		quantity       int
// 		accountBalance float64
// 		expected_1     bool
// 		expected_2     float64
// 	}

// 	runCases := []testCase{
// 		{"1", 2, 226.95, true, 223.95},
// 		{"2", 25, 459, true, 402.75},
// 		{"3", 7, 1185.2, false, 1185.2},
// 		{"4", 5, 0, false, 0},
// 		{"5", 50, 195, true, 70},
// 	}

// 	submitCases := append(runCases, []testCase{
// 		{"6", 0, 100, true, 100},
// 		{"7", 1, 210.24, false, 210.24},
// 		{"1", 2, 2, false, 2},
// 		{"8", 55, 24.5, false, 24.5},
// 		{"9", 1, 999.99, true, 0},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}
// 	skipped := len(submitCases) - len(testCases)

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range testCases {
// 		output_1, output_2 := placeOrder(
// 			test.productID,
// 			test.quantity,
// 			test.accountBalance,
// 		)
// 		if output_1 != test.expected_1 || output_2 != test.expected_2 {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v, %v, %.2f)
// Expecting:  (%v, %.2f)
// Actual:     (%v, %.2f)
// Fail
// `, test.productID, test.quantity, test.accountBalance, test.expected_1, test.expected_2, output_1, output_2)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v, %v, %.2f)
// Expecting:  (%v, %.2f)
// Actual:     (%v, %.2f)
// Pass
// `, test.productID, test.quantity, test.accountBalance, test.expected_1, test.expected_2, output_1, output_2)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true

// ---------------------------------------------------------------------------------------------------
// Lesson 16
// func Test(t *testing.T) {
// 	type testCase struct {
// 		input    []int
// 		expected []int
// 	}

// 	runCases := []testCase{
// 		{
// 			input:    []int{1, 2, 3},
// 			expected: []int{1, 3, 6},
// 		},
// 		{
// 			input:    []int{1, 2, 3, 4, 5},
// 			expected: []int{1, 3, 6, 10, 15},
// 		},
// 	}
// 	submitCases := append(runCases, []testCase{
// 		{
// 			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
// 			expected: []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55},
// 		},
// 		{
// 			input:    []int{0, 0, 0, 0},
// 			expected: []int{0, 0, 0, 0},
// 		},
// 		{
// 			input:    []int{5, -3, -1},
// 			expected: []int{5, 2, 1},
// 		},
// 	}...)

// 	testCases := runCases
// 	if withSubmit {
// 		testCases = submitCases
// 	}

// 	skipped := len(submitCases) - len(testCases)
// 	passCount := 0
// 	failCount := 0

// 	for _, test := range testCases {
// 		f := adder()
// 		results := make([]int, len(test.input))
// 		for i, v := range test.input {
// 			results[i] = f(v)
// 		}
// 		if !slices.Equal(results, test.expected) {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     %v
// Expecting:  %v
// Actual:     %v
// Fail
// `, test.input, test.expected, results)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     %v
// Expecting:  %v
// Actual:     %v
// Pass
// `, test.input, test.expected, results)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	if skipped > 0 {
// 		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
// 	} else {
// 		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// 	}
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true
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