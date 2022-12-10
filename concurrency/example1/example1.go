/*
	Context is value semantic

*/
// Sample program to show how to store and retrieve
// values from a context.
package main

import (
	"context"
	"fmt"
)

// TraceID is represents the trace id.
type TraceID string

// TraceIDKey is the type of value to use for the key. The key is
// type specific and only values of the same type will match.
type TraceIDKey int

func main() {

	// Create a traceID for this request.
	traceID := TraceID("aaaaaaaa-bbbb-cccc-dddd-000000000000")

	// Declare a key with the value of zero of type userKey.
	const traceIDKey TraceIDKey = 0

	// Store the traceID value inside the context with a value of
	// zero for the key type.
	ctx := context.WithValue(context.Background(), traceIDKey, traceID)

	// Retrieve that traceID value from the Context value bag.
	if uuid, ok := ctx.Value(traceIDKey).(TraceID); ok {
		fmt.Println("TraceID:", uuid)
	}

	// Retrieve that traceID value from the Context value bag not
	// using the proper key type.
	if _, ok := ctx.Value(0).(TraceID); !ok {
		fmt.Println("TraceID Not Found")
	}
}
