package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	mycloudeventfunction "github.com/cloud-functions-test"
	"log"
	"os"

	// Blank-import the function package so the init() runs
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	// registering the functions

	t := os.Getenv("FUNCTION_TARGET")
	switch t {
	case "":
		log.Panicln("FUNCTION_TARGET variable is not defined")
	case "HelloWorld":
		// The http function case
		functions.HTTP("HelloWorld", mycloudeventfunction.HelloWorld)
	case "MyCloudEventFunction":
		//First the CloudEvent Function
		functions.CloudEvent("MyCloudEventFunction", mycloudeventfunction.MyCloudEventFunction)
		
	default:
		log.Panicln("Unknown function target: ", t)

	}

	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	// By default, listen on all interfaces. If testing locally, run with
	// LOCAL_ONLY=true to avoid triggering firewall warnings and
	// exposing the server outside of your own machine.
	hostname := ""
	if localOnly := os.Getenv("LOCAL_ONLY"); localOnly == "true" {
		hostname = "127.0.0.1"
	}
	if err := funcframework.StartHostPort(hostname, port); err != nil {
		log.Fatalf("funcframework.StartHostPort: %v\n", err)
	}

}
