package main

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {

	return fmt.Sprintf("Hello, %s!", name)
}
