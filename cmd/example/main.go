package main

import (
	"time"

	"github.com/prasanththumma/gopulse/pkg/gopulse"
)

func main() {

	gopulse.Start()

	for {
		go func() {
			time.Sleep(5 * time.Second)
		}()

		time.Sleep(1 * time.Second)
	}
}
