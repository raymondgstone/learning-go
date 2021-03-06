package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink from its water bowl",
		"The dark bird of prey lands on a small tree after hunting for fish",
	}

	histogram := make(map[string]int)
	done := make(chan struct{})

	// splits and count
	go func() {
		defer close(done) // closes channel upon fn return
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				histogram[word]++
			}
		}
	}()

	<-done // blocks until closed

	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}
