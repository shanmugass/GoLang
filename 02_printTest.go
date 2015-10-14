package main

import "fmt"

func main(){
	var salary float64
	
	salary=234566.2345
	
	//Prints floats	
	fmt.Printf("%f \n", salary)
	
	//prints formated float value with 3 digits
	fmt.Printf("%.3f \n", salary)	
	
	//Prints data type of variable
	fmt.Printf("%T \n", salary)
	
	
	//prints integer value
	fmt.Printf("%d \n", 10)
	
	//prints boolen value
	fmt.Printf("%b \n", 13)
	fmt.Printf("%b \n", salary)
	
	//prints hexa decimal value
	fmt.Printf("%x \n", 11)
				
	//prints octa decimal value	
	fmt.Printf("%o \n", 11)
	
	//prints boolean
	fmt.Printf("%t \n", 20==10)
	fmt.Printf("%t \n", 20!=10)
	
}