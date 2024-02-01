//Range and Close

package main

import "fmt"

func fibonacci(capcity int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < capcity; i++ {
		channel <- x
		x, y = y, x+y
	}
	close(channel)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)

	for i := range ch {
		fmt.Println(i)
	}

}
