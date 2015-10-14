package main

import "fmt"

func main(){
	
	intarray:=[] int {12, 345, 354,5666,22}
	
	fmt.Println(intarray[2])	
	
	//With Index
	for i, value:= range intarray {		
		fmt.Println(i, value);
	}
	
	//With out index
	for _, value:= range intarray {		
		fmt.Println(value);
	}
	
}