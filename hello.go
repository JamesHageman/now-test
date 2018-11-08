package main

import (
	"fmt"
	"net/http"
)

var (
	counter = make(chan func(count int))
)

func init() {
	fmt.Println("Starting now-test...")
	go func() {
		var count int
		for f := range counter {
			count++
			f(count)
		}
	}()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	counter <- func(count int) {
		fmt.Fprintf(w, "Hello, now-test! Invocation number %d.", count)
	}
}
