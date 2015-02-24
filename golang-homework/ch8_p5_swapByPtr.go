/*
Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
 */

package main

import (
	"fmt"
)

func swap(x *int, y *int){
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	x := 1
	y := 2
	fmt.Println(x,y)
	swap(&x,&y)
	fmt.Println(x,y)
}
