package main

import "fmt"

func main() {
	var val interface{} = 7
	fmt.Println(val.(int) + 6)
}
