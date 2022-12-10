// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	//waitForTask()
	//waitForResult()
	//waitForFinished()
	//pooling()
	//fanOut()
	//fanOutSem()
	//drop()
	cancellation()
}

// waitForTask: In this pattern, the parent goroutine sends a signal to a
// child goroutine waiting to be told what to do.
func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("employee : received signal : ", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	ch <- "paper"

	fmt.Println("manager : send signal")

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}

// waitForResult: In this pattern, the parent goroutine waits for the child
// goroutine to finish some work to signal the result.
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee : send signal")
	}()

	p := <-ch
	fmt.Println("manager : received signal : ", p)

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}

func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		close(ch)
		fmt.Println("employee : send signal")
	}()

	_, wd := <-ch
	fmt.Println("manager : received signal : ", wd)

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}

func pooling() {
	ch := make(chan string)

	const emps = 2
	for e := 0; e < emps; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d : received signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : received shutdown signal", emp)
		}(e)
	}

	const work = 10
	for w := 0; w < 10; w++ {
		ch <- "paper"
		fmt.Println("manager : send signal :", w)
	}

	close(ch)
	fmt.Println("manager : send shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}

// fanOut: In this pattern, the parent goroutine creates 2000 child goroutines
// and waits for them to signal their results.
func fanOut() {

	emps := 20
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee : send signal : ", emp)
		}(e)
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manager : received signal : ", emps)
		emps--
	}
	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}

// fanOutSem: In this pattern, a semaphore is added to the fan out pattern
// to restrict the number of child goroutines that can be schedule to run.
func fanOutSem() {
	emps := 6
	ch := make(chan string, emps)

	const cap = 2

	sem := make(chan bool, cap)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("employee : send signal : ", emp)
			}
			<-sem
		}(e)
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manager : receiver signal :", emps)
		emps--
	}

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}

// drop: In this pattern, the parent goroutine signals 2000 pieces of work to
// a single child goroutine that can't handle all the work. If the parent
// performs a send and the child is not ready, that work is discarded and dropped.
func drop() {
	const cap = 5

	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : received signal :", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : sent signal :", w)
		default:
			fmt.Println("manager : dropped data :", w)
		}
	}

	close(ch)
	fmt.Println("manager : send shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}

// cancellation: In this pattern, the parent goroutine creates a child
// goroutine to perform some work. The parent goroutine is only willing to
// wait 150 milliseconds for that work to be completed. After 150 milliseconds
// the parent goroutine walks away.
func cancellation() {

	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee : sent signal")
	}()

	tc := time.After(100 * time.Millisecond)

	select {
	case p := <-ch:
		fmt.Println("manager : received signal :", p)
	case t := <-tc:
		fmt.Println("manager : timeout : ", t)
	}
	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}
