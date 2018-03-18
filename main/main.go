package main

import(

	"fmt"
	"strconv"
)
/************************************************************************
 *
 *  Golang Project
 *
 *  Steve Walsh
 *
 *  18/11/2018
 *
 *  Decimal to binary converter
 *
 *  Goal:
 *  prompts user to enter a decimal number and converts to binary
 *
 *
 *
 ************************************************************************/
func main() {

	var (
		inputValue  int64
		anotherNum   string
		isOver bool = false
	)

	fmt.Printf("\n    ----------------------------     ")
	fmt.Printf("\n                                     ")
	fmt.Printf("\n     Decimal to Binary Converter     ")
	fmt.Printf("\n                                     ")
	fmt.Printf("\n    ----------------------------     ")

	for isOver != true { 														// keep looping until user is winner

		fmt.Printf("\n\n") 												// line break

		fmt.Printf("Enter an integer in decimal number system: ") 		// ask user to input decimal value

		fmt.Scan(&inputValue) 													// stores character from user

		fmt.Printf("\n\n") 												// line break

		fmt.Printf("%d in binary number system is: ",inputValue) 		// print answer intro to user

		fmt.Printf(strconv.FormatInt(inputValue, 2)) 					    // convert and print value in binary

		fmt.Printf("\n\n") 												// line break

		fmt.Printf("Enter another value? ( Y / N ) ") 					// ask user to input decimal value

		fmt.Scan(&anotherNum)													// store users choice to continue

		if anotherNum == "y" || anotherNum == "Y" {								// Scenario 1: user types y

			isOver = false														// repeat for loop

		} else if anotherNum == "n" || anotherNum == "N"{ 						// Scenario 2: user types n

			isOver = true														// exit for loop
		}
	}
}