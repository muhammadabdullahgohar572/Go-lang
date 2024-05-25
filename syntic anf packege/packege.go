package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	render :=bufio.NewReader(os.Stdin)
	fmt.Printf("Plase Enter a value ")

	input,_ :=render.ReadString('\n')
	fmt.Printf("This is your value: %s", input)

}