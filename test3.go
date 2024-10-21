package main

import (
	"fmt"
)

func judge(a int) bool {
	var b = int(a / 2)
	var c = false
	for i := 2; i <= b; i++ {
		if a%i == 0 {
			c = true
		}
	}
	return c
}
func main() {
	var number int
	fmt.Print("Enter an integer: ")
	fmt.Scanf("%d", &number)
	var c = judge(number)
	if c {
		fmt.Printf("%d不是素数", number)
	} else {
		fmt.Printf("%d是素数", number)
	}

}
