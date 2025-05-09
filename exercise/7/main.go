package main

// Untuk n = 14400, keluarannya harus menjadi solution(n) = 4.

// Karena telah berlalu 14400 detik, maka waktu saat ini adalah 04:00. Digit-digitnya jumlahnya adalah 0 + 4 + 0 + 0 = 4, yang merupakan jawaban.

// Untuk n = 48480, keluarannya harus menjadi solution(n) = 14.

// 48480 detik berarti sekarang adalah pukul 13:28, sehingga jawabannya harus 1 + 3 + 2 + 8 = 14.

func main() {

}

func solution(n int) int {
	total := (n / 60) / 60
	return total

}
