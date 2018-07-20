package main

import "fmt"

func main() {
	numbers := make([]int, 100)
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i)
	}

	c := factorial(numbers)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n []int) chan int {
	for _, j := range n {
		out := make(chan int)
		go func() {
			total := 1
			for i := j; i > 0; i-- {
				total *= i
			}
			out <- total
			close(out)
		}()
	}
	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently
and in parallel.
-- Use the "pipeline" pattern to accomplish this

-- What realizations did you have about working with concurrent code when
building your solution?
-- eg, what insights did you have which helped you understand working with
concurrency?
*/
