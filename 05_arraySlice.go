package main

import "fmt"

func main(){
	
	numSlice:=[] int {12, 345, 354,5666,22}
	
	fmt.Println(numSlice[2])	
	
	fmt.Println("Before slice")
	for _, value:= range numSlice {		
		fmt.Println(value);
	}
	
	
	fmt.Println("After slice")
	numSlice2:=numSlice[3:5]
	
	for _, value:= range numSlice2 {		
		fmt.Println(value);
	}
	
	
}