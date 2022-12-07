// Sample program to show how maps are not safe for concurrent use by default.
// The runtime will detect concurrent writes and panic.
package main

import (
	"fmt"
	"sync"
)

// scores holds values incremented by multiple goroutines.
var scores = make(map[string]int)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			scores["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			scores["B"]++
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final scores:", scores)
}

/*
fatal error: concurrent map writes

goroutine 18 [running]:
runtime.throw({0x1008a767c?, 0x0?})
        /opt/homebrew/Cellar/go@1.18/1.18.8/libexec/src/runtime/panic.go:992 +0x50 fp=0x14000040f20 sp=0x14000040ef0 pc=0x10084ab70
runtime.mapassign_faststr(0x0?, 0x0?, {0x1008a498b, 0x1})
        /opt/homebrew/Cellar/go@1.18/1.18.8/libexec/src/runtime/map_faststr.go:212 +0x3bc fp=0x14000040f90 sp=0x14000040f20 pc=0x10082ba9c
main.main.func1()
        /Users/berkay.akcay/Desktop/playground/go-playground/data-race/example5/example5.go:19 +0x54 fp=0x14000040fd0 sp=0x14000040f90 pc=0x1008a4674
runtime.goexit()
        /opt/homebrew/Cellar/go@1.18/1.18.8/libexec/src/runtime/asm_arm64.s:1270 +0x4 fp=0x14000040fd0 sp=0x14000040fd0 pc=0x100876504
created by main.main
        /Users/berkay.akcay/Desktop/playground/go-playground/data-race/example5/example5.go:17 +0x78

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x140000021a0?)
        /opt/homebrew/Cellar/go@1.18/1.18.8/libexec/src/runtime/sema.go:56 +0x2c
sync.(*WaitGroup).Wait(0x1400012c010)
        /opt/homebrew/Cellar/go@1.18/1.18.8/libexec/src/sync/waitgroup.go:136 +0x88
main.main()
        /Users/berkay.akcay/Desktop/playground/go-playground/data-race/example5/example5.go:31 +0xc4

Process finished with the exit code 2
*/
