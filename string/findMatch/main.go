package main

import (
	"fmt"
)

/**
 * Given the words in the magazine and the words in the ransom note,
 * print Yes if he can replicate his ransom note exactly using whole words from the magazine; otherwise, print No.
 * The words in his note are case-sensitive and he must use only whole words available in the magazine.
 * He cannot use substrings or concatenation to create the words he needs.
 *
 * Sample Input 0
 * 6 4
 * give me one grand today night
 * give one grand today
 *
 * Sample Output 0
 * Yes
 */

func main() {
	checkMagazine([]string{"apgo", "clm", "w", "lxkvg", "mwz", "elo", "elo", "lxkvg", "bg", "elo", "apgo", "apgo", "w", "elo", "bg"}, []string{"elo", "lxkvg", "bg", "mwz", "clm", "w"})
}

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
	var count = len(note)
	var magMap = make(map[string]int)
	for _, word := range magazine {
		if _, ok := magMap[word]; ok {
			magMap[word] += 1
		} else {
			magMap[word] = 1
		}
	}
	for _, word := range note {
		if val, ok := magMap[word]; ok && val > 0 {
			magMap[word] -= 1
			count--
		}
	}
	if count > 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

// // Complete the checkMagazine function below.
// func checkMagazine(magazine []string, note []string) string {
// 	var j = 0
// 	var lenMag = len(magazine)
// 	for i := 0; i < len(note); {
// 		if note[i] == magazine[j] {
// 			if j+1 >= len(magazine) {
// 				magazine = append(magazine[:j], magazine[len(magazine):]...)
// 				j = 0
// 			} else {
// 				magazine = append(magazine[:j], magazine[j+1:]...)
// 				j = 0
// 			}

// 			i++
// 		} else {
// 			j++
// 			if j > len(magazine)-1 {
// 				break
// 			}
// 		}
// 	}
// 	if len(note) == lenMag-len(magazine) {
// 		return fmt.Sprintf("YES")
// 	} else {
// 		return fmt.Sprintf("NO")
// 	}
// }
