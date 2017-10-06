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
			//	log.Printf("%v", r)
			c <- fmt.Errorf("%v", r)
		}
	}()

	c <- rec(n)
}

func main() {
	c := make(chan error)
	d := uint64(1e7)
	log.Printf("trying recursion depth %d", d)
	go tryRec(d, c)
	err := <-c
	if err != nil {
		log.Printf("recursion failed for depth: %v", d, err)
	} else {
		log.Printf("success for depth %v", d)
	}
}
