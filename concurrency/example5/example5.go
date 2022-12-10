// Sample program that implements a web request with a context that is
// used to timeout the request if it takes too long.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	// Create new request
	req, err := http.NewRequest("GET", "https://berkayakcay.com/", nil)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// Create a context with a timeout of 50 milliseconds
	ctx, cancel := context.WithTimeout(req.Context(), 500*time.Millisecond)
	defer cancel()

	// Declare a new transport and client for the call
	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := http.Client{
		Transport: &tr,
	}

	// Make the web call in a separate goroutine so it can be cancelled.
	ch := make(chan error, 1)

	go func() {
		log.Println("Starting Request")

		// Make the web call and return any error.
		resp, err := client.Do(req)
		if err != nil {
			log.Println("ERROR:", err)
			ch <- err
			return
		}

		// Close the response body on the return
		defer resp.Body.Close()

		// Write the response to stdout
		io.Copy(os.Stdout, resp.Body)
		ch <- nil
	}()

	// Wait the request or timeout
	select {
	case <-ctx.Done():
		fmt.Println("timeout, cancel work...")

		// Cancel the request and wait for it to complete.
		tr.CancelRequest(req)
		log.Println(<-ch)
	case err := <-ch:
		if err != nil {
			log.Println(err)
		}
	}

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------------------")
}
