package main

import (
	"fmt"
)


//Printer function 1
func myTextPrinter ( print_message string ) {
	fmt.Println(print_message)
}

//No printer available error message
func noPrinter( print_message string ) {
	print_message = "No printer available!"
	fmt.Println( print_message )
}

//Function factory, return the correct printer type
func getPrinter ( printType int) (func(string)) {
	switch( printType ) {
	case 0:
		return myTextPrinter
	default:
		return noPrinter
	}
}


func main() {

	var myMsg = "print this!"

	var aPrinter = getPrinter(0)

	aPrinter(myMsg)
}
