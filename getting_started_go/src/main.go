package main

import (
	"example_module/src/greetings"
	"fmt"
)

func main() {
	fmt.Printf("%s, world!\n", greetings.Greeting())
}
