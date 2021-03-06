package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// Error variable put inside the function to have access to f.
		ErrNorgateMath := fmt.
			Errorf(
				"norgate math again: square root of a negative number : %v",
				f)
		return 0, ErrNorgateMath
	}
	//implementation
	return 42, nil
}
