//Your function must be in a Go package at the root of your project.
//The package and its source files can have any name, except your function cannot be in package main.
//If you need a main package, for example for local testing, you can create one in a subdirectory:

package mycloudeventfunction

import (
	"context"

	"github.com/cloudevents/sdk-go/v2/event"
)

// Function myCloudEventFunction accepts and handles a CloudEvent object
func MyCloudEventFunction(ctx context.Context, e event.Event) error {
	// Your code here
	// Access the CloudEvent data payload via e.Data() or e.DataAs(...)

	// Return nil if no error occurred
	return nil
}
