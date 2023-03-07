package main

import "fmt"

func main() {
	var sayBye = func(n string) {
		fmt.Printf("Bye %s", n)
	}

	sayBye("Bob")
}
