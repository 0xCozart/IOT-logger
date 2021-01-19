package utils

import "fmt"

// Must basic error handler
func Must(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
		panic(err)
	}
}
