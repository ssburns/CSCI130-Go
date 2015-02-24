package main

import "fmt"

// TYPE
type contact struct {
	name     string
	greeting string
}

// TYPES CAN CONTAIN METHODS & RETURN
func (c contact) sayHello() string {
	return "METHOD: " + c.name + " says hello."
}

func (c *contact) rename(newName string) {
	c.name = newName
}

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	// TYPES CAN CONTAIN DATA
	var friend = contact{"Marcus", "Hello!"}
	fmt.Println("DATA: " + friend.name)
	fmt.Println("DATA: " + friend.greeting)

	// TYPES CAN USE METHODS
	fmt.Println(friend.sayHello())
	friend.rename("Jenny")
	fmt.Println("DATA: " + friend.name)
	fmt.Println(friend.sayHello())


	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	fmt.Println((true && false) || (false && true) || !(false && false))

}
