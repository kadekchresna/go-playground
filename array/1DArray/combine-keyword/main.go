package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkV2([]string{"hewhyyellowllsocatpapplebatgooodbye", "apple", "bat", "catt", "goodbye", "hello", "yellow", "why"}))
}

/**
 *
 */

func checkV2(arr []string) string {
	var aMap = make(map[string]bool)
	var aArr = make([]string, 0)
	var firstEl = arr[0]
	var final string
	var seperator int
	var multiple string
	var i = 0
	var j = 0

	// delete first index agar memudahkan untuk pengecekan deengan index pertama
	arr = append(arr[:0], arr[1:]...)
	for {
		// pengecekan huruf pertama dari kata di index pertama dan kata dalam dictionary index i
		if arr[i][0] == firstEl[j] {
			// pengecekan apakah index j + panjang dari kata dictionary index i
			if j+len(arr[i]) <= len(firstEl) {
				// flagging apakah kata index pertama dari j + panjang dari kata dictionary  sama dengan kata dictionary index i
				subCheck := true
				for k := 0; k < len(firstEl[j:j+len(arr[i])]); k++ {
					if firstEl[j:len(firstEl)][k] != arr[i][k] {
						subCheck = false
					}

				}
				// pengecekan apakah kata pernah muncul di index pertama atau di dictionary dan pengecekan apakan kata sama dengan yang ada di
				// index pertama dan kata dictionary index i
				if _, ok := aMap[arr[i]]; !ok && subCheck {
					// jika sama maka dimasuk kan ke temporary map dan di tambahkan array result
					aMap[arr[i]] = true
					aArr = append(aArr, arr[i])
					// pointer di index peretama akan pindah ke kata selanjutnya yang ada di index pertama
					j += len(arr[i])
					// pointer di dictionary akan di reset ke awal lagi
					i = 0
					// pengecekan apakah sudah ada kata yang sama sebelumnya. jika ada akan di masukan ke temporary string yang tadinya di pakai untuk pengecekan
					// di akhir.
				} else if ok && subCheck {
					multiple += arr[i]
					j += len(arr[i])
					i = 0
					// jika tidak sama maka pointer dictionary akan pindak ke kata selanjutnya
				} else if !subCheck {
					i++
				}
				// jika diakhir pengecekan index peretama dengan kata di dictionary melebihi length index pertama, loop akan di berhentikan
			} else {
				break
			}
			// jika pengecekan huruf pertama di dengan pointer j di index pertama dan huruf pertama di pointer i di kata dictionary tidak match
			// maka pointer dictionary akan di pindahkan ke selanjutnya
		} else {
			i++
		}

		// setiap kata yang match pointer dictionary akan di reset. jika pointer j di index pertama tidak ada yang match dengan
		// huruf di awal kata dictionary maka pointer dictionary akan di reset lagi dan pointer index pertama akan di pindahkan ke huruf
		// selanjutnya di index pertama
		if i >= len(arr) {
			i = 0
			j++
			seperator++
		}
		// jika pengecekan berhasil sampai akhir huruf index pertama loop akan di break
		if j > len(firstEl)-1 {
			break
		}
	}

	// pengecekan jika tidak ada kata yang match di index pertama dan dalam kata di dictiionary
	if len(aMap) == 0 {
		return "not possible"

		//jika ada kata yang match di index pertama dan dalam kata di dictiionary, maka akan di tampilkan menjadi string deengan comma seperated
	} else {
		for _, value := range aArr {
			final += fmt.Sprintf("%s,", value)
		}
		return final[0 : len(final)-1]
	}
}

func check(arr []string) string {
	var aMap = make(map[string]int)
	var tempFirst = arr[0]
	var tempMap = make(map[string]int)
	var res = make([]string, 0)
	var final string
	var lResTotal int

	for _, value := range arr[0] {
		if _, ok := aMap[string(value)]; ok {
			aMap[string(value)] += 1
		} else {
			aMap[string(value)] = 1
		}
	}
	arr = append(arr[:0], arr[1:]...)

	for _, str := range arr {
		c := len(str)
		for k, value := range aMap {
			tempMap[k] = value
		}

		for _, sub := range str {
			_, ok := tempMap[string(sub)]

			if ok && tempMap[string(sub)] > 0 {
				tempMap[string(sub)] -= 1
				c--
			}

			if c == 0 {
				res = append(res, str)
			}
		}
	}

	fmt.Println(aMap)
	fmt.Println(res)

	for _, value := range res {
		lResTotal += len(value)
	}

	if lResTotal != len(tempFirst) {
		return "not possible"
	} else {
		var i int
		var j int
		var k int
		var commas int

		for {
			if string(res[k][j]) == string(tempFirst[j+i]) {
				if len(final) > 0 {
					final += fmt.Sprintf(",%s", res[k])
					commas++
				} else {
					final += fmt.Sprintf("%s", res[k])
				}

				i += len(res[k])
			} else {
				k++
				if k > len(res)-1 {
					k = 0
				}

			}

			if len(final)-commas == len(tempFirst) {
				break
			}

		}
		return final

	}

}
