

package main

import "fmt"

func main() {
	fmt.Println("HI 😀")
 
// 	Map :=make(map[any]any)

// 	Map["H"]="HTML"
// 	Map["JS"]="JAVA"
// 	Map["REACT"]="REACT JS"
// 	Map["NODE"]="NODE JS"
// // 	delete(Map,"JS")
// // fmt.Println(Map)
//  for key, value :=range Map{
// 	fmt.Println("This is a Arry %v\n",key ,value)
//  }

Map :=map[any]string{
	"name":"abdullah",
	"id":"2",
}
fmt.Println(Map)
}