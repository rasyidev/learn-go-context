package learngocontext

import (
	"fmt"
	"runtime"
	"testing"
)

func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	destination := CreateCounter()
	for i := range destination {
		fmt.Println("Counter", i)
		if i == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

/*
=== RUN   TestContextWithCancel
Total Goroutine: 2
Counter 1
Counter 2
Counter 3
Counter 4
Counter 5
Counter 6
Counter 7
Counter 8
Counter 9
Counter 10
Total Goroutine: 3
--- PASS: TestContextWithCancel (0.00s)
PASS
ok      learn-go-context        0.048s

Masih ada 1 goroutine yang berjalan (goroutine leaks), ini bahaya karena bisa
memperlambat aplikasi jika diakses oleh banyak user.

Kok bisa?
karena di function CreateCounter ada infinite loop yang membuat goroutine tetap berjalan
*/
