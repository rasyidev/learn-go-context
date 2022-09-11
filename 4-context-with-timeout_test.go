package learngocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterWithTimeout(ctx context.Context) chan int {
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

func TestCreateCounterWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()
	// cancel: function untuk mengirim sinyal context
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	// mengirimkan sinyal cancel ke context, otomatis context.Done, semua proses dibatalkan
	defer cancel()

	destination := CreateCounterWithTimeout(ctx)
	for i := range destination {
		fmt.Println("Counter", i)
		if i == 10 {
			break
		}
	}

	// biar gak kecepetan. Nunggu semua goroutine selesai, karena goroutine asynchronous
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

/*
=== RUN   TestCreateCounterWithTimeout
Total Goroutine: 2
Counter 1
Counter 2
Counter 3
Counter 4
Counter 5
Total Goroutine: 2
--- PASS: TestCreateCounterWithTimeout (7.03s)
PASS
ok      learn-go-context        7.079s
*/
