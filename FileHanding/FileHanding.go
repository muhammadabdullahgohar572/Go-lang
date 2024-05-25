package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Print("HI 😀");

	content :="You Need This File and Deffer.go"

	File,err :=os.Create("./Test.text")

	if err !=nil{
		panic(err)
	}

	lenght,err :=io.WriteString(File,content)

	if err !=nil {
		panic(err)
	}

	fmt.Println("This is Lenght",lenght)

	defer File.Close()
	Readfile("./Test.text")
	
}

func Readfile(filename string){
 
	Deffer,err :=ioutil.ReadFile(filename)
	if err !=nil {
		panic(err)
	}

	fmt.Print("This is File ",string(Deffer))
}