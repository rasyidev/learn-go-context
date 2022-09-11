## Context
- Data yang membawa value dan sinyal - sinyal (cancel, timeout, deadline)
- Biasanya dibuat per request
- Digunakan untuk mempermudah untuk meneruskan value dan sinyal antar proses
- Hampir semua bagian di Go Lang memanfaatkan context seperti:
  - Database
  - Http Server
  - Http Client
  - dll
- Diimplementasikan dalam sebuah interface Context
```go
type Contex interface{
  Deadline()(deadline time.Time, ok bool)
  Done() <- chan struct{}
  Err() error
  Value(key interface{}) interface{}
}
```

## Membuat Context
- `context.Background()`, Membuat context kosong. Tidak pernah dibatalkan, tidak pernah timeout, tidak memiliki value apapun. Biasanya digunakan di main function, test, atau dalam awal proses request
- `context.TODO()`, Sama seperti `context.Background()` tapi belum jelas context apa yang ingin digunakan.

## Parent dan Child Context
- Menganut konsep parent-child
- 1 parent - banyak child
- parent-child selalu terhubung, tidak bisa dilepas
- Pembatalan terhadap context parent, otomatis pembatalan juga terhadap context childnya
- Penyisipan data terhadap context parent, otomatis penyisipan juga terhadap context childnya
- Immutable, setelah dibuat value tidak dapat diubah lagi
- Ketika generate value baru terhadap context, akan dibuat child dan diisi dengan value tersebut

## Context with Value
- Context tidak memiliki value saat pertama kali dibuat
- Dapat menambahvalue dengan key-value dalam context
- `context.WithValue(parent, key, value)`

## Context with Cancel
- Menambahkan sinyal pembatalan dalam suatu context
- Biasanya dibutuhkan jika ingin memberi sinyal cancel ke suatu proses (biasanya goroutine)
- Harus diingat bahwa goroutine yang menggunakan context harus melakukan pengecekan terhadap contextnya. Jika tidak, tidak ada gunanya.
- `context.WithCancel(parent)`

## Context with Timeout
- Mengirimkan sinyal cancel atau pembatalan setelah jangka waktu tertentu
- Tidak perlu mengirim sinyal cancel secara manual
- `context.WithTimeout(parent, duration)`