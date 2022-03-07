package main

import (
	"fmt"
)

func main() {
	fmt.Print("My App...")
	DoAction()
}

func DoAction() {
	fmt.Print("foo")
}
