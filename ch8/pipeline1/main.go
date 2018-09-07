// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 228.

// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import "fmt"

//!+
func main() {
	naturals := make(chan int64)
	squares := make(chan int64)

	// Counter
	go func() {
		for x := int64(0); ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			fmt.Println("naturals == ",naturals)
			squares <- x * x *x * x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

//!-
