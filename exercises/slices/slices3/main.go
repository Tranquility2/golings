package main

import "fmt"

func main() {
	names := []string{"John", "Maria", "Carl", "Peter"}
	names = append(names, "Alex")
	fmt.Println(names)
}
