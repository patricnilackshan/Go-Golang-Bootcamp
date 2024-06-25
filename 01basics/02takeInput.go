package main

// import (
// 	"fmt"
// )

// // Can't get space seperated inputs and
// // save it as one variable in this method
// func main() {
// 	var name string
// 	fmt.Print("Enter your name here: ")
// 	fmt.Scanln(&name)
// 	fmt.Println(name)
// }

// // To get strings with space between
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter your name here: ")
// 	name, _ := reader.ReadString('\n')
// 	fmt.Println(name)
// }

// //Get Float as input
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter your rating: ")
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSpace(input)
// 	rating, _ := strconv.ParseFloat(input, 64)
// 	fmt.Printf("Rating: %v", rating)
// }
