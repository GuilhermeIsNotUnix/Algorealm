package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			if i%3 == 0 {
				fmt.Println("Fizz")
			}

			if i%5 == 0 {
				fmt.Println("Buzz")
			}
		}

		time.Sleep(time.Second)
		i++
	}
}
