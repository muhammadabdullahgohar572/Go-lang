package main

import "fmt"

func main() {
	fmt.Println("HI 😀")

	// data :=test(10 ,20)

	// fmt.Println(data)


	//  data :=test(1,23,1,4,456,2,3,1)
    //  fmt.Println(data)

	 data,ok :=test(1,23,1,4,456,2,3,1)
     fmt.Println(data,ok)
	}

// func test(a  int ,b int)int{
 
// 	return a+b
	 
// }



// func test(vals ... int)int{

// 	num :=0

// 	for _, n:=range vals{
//       num +=n
// 	}
//  return num

// }

func test(vals ... int)(int ,string){

	num :=0

	for _, n:=range vals{
      num +=n
	}
 return num ,"Ok 😀"

}