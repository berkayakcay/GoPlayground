// Cache Coherency and False Sharing
// - Thread memory access matters.
// - If your algorithm is not scaling look for false sharing problems.
//
// Links
// - Eliminate False Sharing - Herb Sutter http://www.drdobbs.com/parallel/eliminate-false-sharing/217500206
// This content is provided by Scott Meyers from his talk in 2014 at Dive:
//
// CPU Caches and Why You Care (30:09-38:30) https://youtu.be/WDIkqP4JbkE?t=1809
//
// go build -race
//
// Sample program to show how to create race conditions in
// our programs. We don't want to do this.
package main

import (
	"fmt"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int

func main() {

	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {

				// Capture the value of Counter.
				value := counter

				// Increment our local value of Counter.
				value++

				// Store the value back into Counter.
				counter = value
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

/*
==================
WARNING: DATA RACE
Read at 0x000102ac1960 by goroutine 8:
  main.main.func1()
      /Users/berkay.akcay/Desktop/playground/go-playground/data-race/example1.go:40 +0x40

Previous write at 0x000102ac1960 by goroutine 7:
  main.main.func1()
      /Users/berkay.akcay/Desktop/playground/go-playground/data-race/example1.go:46 +0x54

Goroutine 8 (running) created at:
  main.main()
      /Users/berkay.akcay/Desktop/playground/go-playground/data-race/example1.go:36 +0x58

Goroutine 7 (finished) created at:
  main.main()
      /Users/berkay.akcay/Desktop/playground/go-playground/data-race/example1.go:36 +0x58
==================
Final Counter: 4
Found 1 data race(s)
*/
