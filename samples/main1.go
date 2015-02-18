package main

import (
	"fmt"
	tcm "CSCI130-Go/tmmath"
)

func main() {


	//mathy things
	numberSlice := []int{10,20,30}
	sumOfNumbers := tcm.Sum(numberSlice)
	fmt.Println(sumOfNumbers)
	avgOfNumbers := tcm.Average(numberSlice)
	fmt.Println(avgOfNumbers)
	

}
