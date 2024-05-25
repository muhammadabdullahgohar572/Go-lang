package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi 😯")

	// // Initialize a reader for standard input
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Println("You are a good boy. Please answer with Yes or No.")

	// // Read user input
	// input, _ := reader.ReadString('\n')
	// input = strings.TrimSpace(input) // Remove newline and any surrounding whitespace

	// // Handle specific text inputs
	// if strings.EqualFold(input, "yes") || strings.EqualFold(input, "no") {
	// 	fmt.Println("Your answer is:", input)
	// 	return
	// }

	// // Try to parse the input as a float
	// parsedValue, err := strconv.ParseFloat(input, 64)
	// if err != nil {
	// 	fmt.Println("Error:", err, "Invalid value entered.")
	// } else {
	// 	fmt.Printf("Your answer as a float is: %f\n", parsedValue)
	// }

	 
	render :=bufio.NewReader(os.Stdin)
   fmt.Println("You are a good boy. Please answer with Yes or No.")
   
     input,_ :=render.ReadString('\n')
     input=strings.TrimSpace(input)
	
	 if strings.EqualFold(input,"yes") ||strings.EqualFold(input,"no") {
		fmt.Println("Your answer is:", input)
		return
	 }
	
	 data,err :=strconv.ParseFloat(input,64)
	if err != nil {
		fmt.Println("Error:", err, "Invalid value entered.")
	} else {
		fmt.Printf("Your answer as a float is: %f\n", data)
	}


}
