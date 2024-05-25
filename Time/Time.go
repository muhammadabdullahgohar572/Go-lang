package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hi how are you 😙😃😄")

	time :=time.Now()
	foematdata :=time.Format("02-05-2006 3:04:PM Wednesday")
	fmt.Println(foematdata)
} 