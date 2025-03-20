package main

import "fmt"

// Perlu diketahui, tipe yang didefinisikan menggunakan type constraint ini hanya bisa dimanfaatkan pada generic.
// Tipe jenis ini tidak bisa digunakan di luar scope kode generic.
// Sebagai contoh,
// coba deklarasikan var s Number dalam fungsi main(), hasilnya akan muncul syntax error.

type Number interface {
	int64 | float64
}

func main() {
	// error
	// int does not satisfy Number (int missing in int64 | float64)
	// PrintSlice([]int{1,2,3,4})

	PrintSlice([]int64{1, 2, 3, 4})

}

func PrintSlice[T Number](t []T) {
	for _, v := range t {
		fmt.Println(v)
	}
}
