package main

import (
	"fmt"
	// "runtime"
	"sync"
	"time"

	"github.com/buffer-mutex-time/model"
)

func main() {

	var wg sync.WaitGroup
	// var product model.Product

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go model.Reduce(5, &wg)
	}
	wg.Wait()
	// fmt.Printf("sisa product %d\n", product.Qty)
	fmt.Println("Proses seleai....")

	// numCPU := runtime.NumCPU()
	// fmt.Printf("Jumlah CPU: %d\n", numCPU)

	// runtime.GOMAXPROCS(8)
	// fmt.Printf("Mengatur GOMAXPROCS menjadi: %d\n", runtime.GOMAXPROCS(0))

	// for i := 0; i <= 5; i++ {
	// 	wg.Add(1)
	// 	go ganjil(10, &wg)
	// }

	// wg.Wait()
	// fmt.Println("selesai...")
	// wg.Wait()
	// fmt.Println("pekerjaan selesai")
	// ch := make(chan string)
	// ch1 := make(chan string)
	// ch2 := make(chan string)
	// msg1("hello, saya pesan 1", ch)
	// msg2("hello, saya pesan 2", ch1)
	// msg3("hello, saya pesan 3", ch2)
	// result1 := <-ch
	// fmt.Println(result1)

	// result2 := <-ch1
	// fmt.Println(result2)

	// result3 := <-ch2
	// fmt.Println(result3)

	// TestSendData()
	// TestSendData2()
	// CountDown(10)
}

func TestSendData() {
	ch := make(chan string)

	go func() {
		fmt.Println("mengirim data 1 tanpa buffer")
		ch <- "data1"
		fmt.Println("data 1 telah dikirim")
	}()

	// time.Sleep(2 * time.Second)

	msg := <-ch
	fmt.Println("datanya adalah :", msg)
}

func TestSendData2() {
	ch := make(chan string, 5)

	go func() {
		fmt.Println("mengirim data 2 dengan buffer")
		ch <- "data1"
		fmt.Println("data 2 telah dikirim")

		fmt.Println("mengirim data 2 dengan buffer")
		ch <- "data1"
		fmt.Println("data 2 telah dikirim")
	}()

	time.Sleep(2 * time.Second)

	msg := <-ch
	fmt.Println("datanya adalah :", msg)

	msg2 := <-ch
	fmt.Println("datanya adalah :", msg2)
}

func CountDown(value int) {
	ch := make(chan int, 2)

	go func() {
		for i := value; i >= 0; i-- {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()

	// result := <-ch
	// fmt.Println(result)

	for num := range ch {
		fmt.Println(num)
	}

	fmt.Println("Countdown selesai!")
}

// 3 func di goroutine berbeda

func msg1(pesan string, ch chan<- string) {
	// ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- pesan
	}()

	// result := <-ch
	// fmt.Println(result)
}

func msg2(pesan string, ch chan<- string) {
	// ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- pesan
	}()

	// result := <-ch
	// fmt.Println(result)
}
func msg3(pesan string, ch chan<- string) {
	// ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- pesan
	}()

	// result := <-ch
	// fmt.Println(result)
}

// Nilai Ganjil
func ganjil(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= n; i++ {
		if i%2 != 0 {
			// wg.Add(1)
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}
	}
}

// kurangin quantity struct
