package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	InputRender := bufio.NewReader(os.Stdin)

	fmt.Printf("Plase Enter a value")
	 
	input ,_ := InputRender.ReadString('\n')

	 data,err :=strconv.ParseFloat(strings.TrimSpace(input),64)
    
	 if err !=nil{
	fmt.Print(err)

	 } else{
	fmt.Printf("Plase Enter a value %T",data)

	 }

}