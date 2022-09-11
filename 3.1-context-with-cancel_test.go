package learngocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterNoGoroutineLeaks(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		// tutup channel setelah goroutine selesai dijalankan
		defer close(destination)

		counter := 1

		for {
			select {
			// saat context selesai
			case <-ctx.Done():
				// break perulangan.
				return
				// break // kalau mau break select, bukan perulangannya

			// kalau context belum selesai
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func TestCreateCounterNoGoroutineLeaks(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	// cancel: function untuk mengirim sinyal context
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterNoGoroutineLeaks(ctx)
	for i := range destination {
		fmt.Println("Counter", i)
		if i == 10 {
			break
		}
	}
	// mengirimkan sinyal cancel ke context, otomatis context.Done, semua proses dibatalkan
	cancel()

	// biar gak kecepetan. Nunggu semua goroutine selesai, karena goroutine asynchronous
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

/*
=== RUN   TestCreateCounterNoGoroutineLeaks
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
Total Goroutine: 2
--- PASS: TestCreateCounterNoGoroutineLeaks (2.00s)
PASS
ok      learn-go-context        2.053s
*/
