package main

import "fmt"

func main() {

	//  Week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

     
	//  for i := 0; i < len(Week); i++ {
	// 	fmt.Println(Week[i])
	//  }


	//  for data :=range Week{
	// 	fmt.Println(Week[data])
	//  }
	
	// for index, data :=range Week{
	// 	fmt.Printf("This is an index %v and This is a Day %v\n",index,data)
	// }

	// for _, data :=range Week{
	// 	fmt.Printf("This is an index %v and This is a Day %v\n",_,data)
	// }

	//  num :=0

	//  for num <10 {
	// 	fmt.Println("This is a value :",num)
	// 	num++
	//  }

	//  num :=0

	//  for num <10 {

	// 	 if num==5 {
	// 		 num++
	// 		continue
	// 	 }
	// 	fmt.Println("This is a value :",num)
	// 	num++
	//  }


	// 	 num :=0

	//  for num <10 {

	// 	 if num==5 {
	// 		 num++
	// 		break
	// 	 }
	// 	fmt.Println("This is a value :",num)
	// 	num++
	//  }


		 num :=0

	 for num <10 {

		 if num ==2 {
			goto lco
		 }
		
		fmt.Println("This is a value :",num)
		num++
	 }

	 lco :
	 fmt.Print("Jump")
}