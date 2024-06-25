package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	var name string
// 	var userRating float64

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Plase enter your full name: ")
// 	name, _ = reader.ReadString('\n')
// 	name = strings.TrimSpace(name)
// 	fmt.Printf("\nHi, %v\n", name)

// 	reader = bufio.NewReader(os.Stdin)
// 	fmt.Print("Please rate our service between 1 and 5: ")
// 	inputRating, _ := reader.ReadString('\n')
// 	inputRating = strings.TrimSpace(inputRating)
// 	userRating, _ = strconv.ParseFloat(inputRating, 64)
// 	fmt.Printf("\nThankyou for rating us %.1f star rating.\n", userRating)

// 	// In Go, the 'else if' statement should be written without the
// 	// parentheses and 'else if' or 'else' should be on the same line
// 	// as the closing brace of the previous 'if' or 'else if' block.

// 	if userRating == 5 {
// 		fmt.Println("Bonus for team for a 5 star service.")
// 	} else if userRating < 5 && userRating >= 3 {
// 		fmt.Println("We are always improving.")
// 	} else if userRating < 3 {
// 		fmt.Println("Need serious work on our side.")
// 	}
// }
