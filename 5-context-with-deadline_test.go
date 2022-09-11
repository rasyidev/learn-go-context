package learngocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterWithDeadline(ctx context.Context) chan int {
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
				// simulasi proses yang lambat
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	// cancel: function untuk mengirim sinyal context
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))

	destination := CreateCounterWithDeadline(ctx)
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
=== RUN   TestContextWithDeadline
Total Goroutine: 2
Counter 1
Counter 2
Counter 3
Counter 4
Counter 5
Total Goroutine: 2
--- PASS: TestContextWithDeadline (7.04s)
PASS
ok      learn-go-context        7.089s
*/
