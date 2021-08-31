package main

import (
	"fmt"

	"fatalerror.com/writer"
)

func main() {
	// Get a greeting message and print it.
	message := writer.Hello("Gladys")
	fmt.Println(message)
}
