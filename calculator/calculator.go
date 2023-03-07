package calculator

import "fmt"

func Calculator() {
	var x, y, z int
	var name string
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Scanln(&name)
	switch name {
	case "+":
		z = x + y
		fmt.Println(z)
	case "-":
		z = x - y
		fmt.Println(z)
	case "*":
		z = x * y
		fmt.Println(z)
	case "/":
		z = x / y
		fmt.Println(z)
	}

}
