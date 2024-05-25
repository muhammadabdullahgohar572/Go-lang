package main

import "fmt"

type User struct {
	name     string
	password int
	status   bool
	Age      float64
}
func main() {
	fmt.Println("HI 😀")

	UserData :=User{"Abdullah",10,false,10.10}

	fmt.Println(UserData)

	fmt.Printf("This is a my details %+v\n", UserData)

}

