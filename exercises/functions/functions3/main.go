package main

import "fmt"

func main() {
	call_me(5)
}

func call_me(num int) {
	for n := 0; n <= num; n++ {
		fmt.Printf("Num is %d\n", n)
	}
}
