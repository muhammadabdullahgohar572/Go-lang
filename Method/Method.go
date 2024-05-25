package main

import "fmt"

func main() {

	fmt.Println("Hi 😀")

	Userdata :=Data{"Abdullah","abdullah1232gmail.com",133}

	fmt.Println(Userdata)

	fmt.Printf("This is a Details %+v\n",Userdata)

	fmt.Printf("This is Name %v and This a Email %v\n",Userdata.userName ,Userdata.Password)
   fmt.Print("----------------")
	 Userdata.NewData()
}

type Data struct {
	userName string
	Email    string
	Password int
}

func (D Data) NewData(){
 D.Email="Test.gmail.com"
 fmt.Printf("This is a New Email %v",D.Email)
}

