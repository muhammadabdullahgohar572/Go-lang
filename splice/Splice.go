package main

import (
	"fmt"
	"sort"
)

func main() {

	var name = []string{"Abdullah", "Ali", "Umer"}

	fmt.Printf("Type of data %T\n", name)
	name = append(name, "Jhone", "Test")

	fmt.Print(name)

	num :=make([]int,4)

	num [0]=45;
	num [1]=5;
	num [2]=536;
	num [3]=98;

	sort.Ints(num)

 name= append(name[2:]  )

	fmt.Println(num)
	fmt.Println(name)

}
