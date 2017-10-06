package main

import (
	"fmt"
	"log"
)

func rec(n uint64) error {
	if n == 0 {
		return nil
	}
	return rec(n - 1)
}

func tryRec(n uint64, c chan<- error) {
	defer func() {
		if r := recover(); r != nil {
			c <- fmt.Errorf("%v", r)
		}
	}()

	c <- rec(n)
}

func main() {
	c := make(chan error)
	d := uint64(1e3)
	log.Printf("trying recursion depth %d", d)
	go tryRec(1000, c)
	err := <-c
	if err != nil {
		log.Printf("recursion failed for depth: %v", d, err)
	} else {
		log.Printf("success for depth %v", d)
	}
}
