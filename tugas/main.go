package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sensor(name string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-timeout:
			fmt.Printf("%s: sensor timeout\n", name)
			return
		case <-ticker.C:
			if rand.Float32() < 0.8 {
				ch <- fmt.Sprintf("%s: data diterima", name)
			} else {
				fmt.Printf("%s: tidak ada respon\n", name)
			}
		}
	}
}

func main() {
	ch := make(chan string, 3)
	var wg sync.WaitGroup

	sensors := []string{"Sensor Suhu", "Sensor Kelembaban", "Sensor Tekanan"}
	for _, sensorName := range sensors {
		wg.Add(1)
		go sensor(sensorName, ch, &wg)
	}

	go func() {
		for data := range ch {
			fmt.Println(data)
		}
	}()

	wg.Wait()

	close(ch)

	fmt.Println("Selesai...")
}
