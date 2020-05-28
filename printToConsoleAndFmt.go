package main

import "fmt"

func main() {

	fmt.Printf("%T" ,10) //Prints type of the variable.
	fmt.Printf("Hello %T " ,10) //Appending hello to the type.

	fmt.Printf("%v" ,0000) //prints value of the variable.
		
	var x string  = fmt.Sprintf("%v%%" ,9898); //stores result varialbe x. two %% is one %
	fmt.Println(x)  
	
	fmt.Printf("this is %t" ,true) // %t for boolean value.
	fmt.Println("")

	fmt.Printf("Number with base is : %b" ,1025) 		// %b for binary 
	fmt.Println("")

	fmt.Printf("Number with octal is : %o" ,1025) 		// %o for octal 
	fmt.Println("")

	fmt.Printf("Number with decimal is : %d" ,1025) 	// %d for decimal 
	fmt.Println("")

	fmt.Printf("Number with xehadecimal is : %x" ,1025) // %x for xehadecimal 
	fmt.Println("")

	fmt.Printf("Number with xehadecimal is : %X" ,1025) // %X for xehadecimal displays in caps
	fmt.Println("")

	//Floating points representation.
	fmt.Printf("Number in exponential  %e" ,2.36589665) // %e scientific notation
	fmt.Println("")
	fmt.Printf("Number in exponential  %f" ,2.36589665) // %f decimal no exponential
	fmt.Println("")
	fmt.Printf("Number in exponential  %g" ,2.36589665) // %g large exponential
	fmt.Println("")

	//String
	fmt.Printf("String value is =  %s" ,"Abhishek") 					// %s prints string
	fmt.Println("")
	fmt.Printf("String value in double quote is =  %q" ,"Abhishek") 	// %q prints string in ""
	fmt.Println("")


	//Width and Precision.
	fmt.Printf("String %20q", "Abhay") 				//padding on left
	fmt.Println("")

	fmt.Printf("String %-20q is great.", "Abhay")	 // padding on right													//
	fmt.Println("")
	
	fmt.Printf("Round of number to 2 decimal point %.2f", 3.25698745) // Round of number to 2 decimal point
	fmt.Println("")

	fmt.Printf("Round of number to whole number %.f", 3.25698745) //prints the whole number.
	fmt.Println("")

	fmt.Printf("Number: %07d is result of padding \n" , 45) // padddigit 0 to len 7
	fmt.Printf("Number: %-7d is result of padding \n" , 45) // padddigit 0 to len 7

}
