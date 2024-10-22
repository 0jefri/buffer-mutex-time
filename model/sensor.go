package model

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Sensor struct {
	Name string
}

func (s *Sensor) Sensore(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-timeout:
			fmt.Printf("%s: data diterima\n", s.Name)
			return
		case <-ticker.C:
			if rand.Float32() < 0.8 {
				ch <- fmt.Sprintf("%s: tidak ada respon\n", s.Name)
			}
		}
	}
}
