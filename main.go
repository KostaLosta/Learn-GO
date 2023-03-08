package main

import (
	"fmt"
)

func main() {

	type Parent struct{}

	(func(c *Parent) Print)()
	{
		fmt.Println("parent")
	}

	type Child struct {
		Parent
	}
	(func(p *Child) Print)()
	{
		fmt.Println("child")

	}
	var x Child

	x.Print()
}
