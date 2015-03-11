package main

import (
	"fmt"
	"os"
)

func main() {
	envv := os.Environ()

	for key, value := range(envv) {
		fmt.Println(key, ":", value)
	}
}
