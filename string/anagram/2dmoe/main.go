package main

import (
	"fmt"
	"sort"
)

/**
 * check anagram dalam array of string dan kelompokan string dalam anagram yang sama
 */

type subS []rune

func (a subS) Len() int           { return len(a) }
func (a subS) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a subS) Less(i, j int) bool { return a[i] < a[j] }

func main() { // O(n^2*log(n))
	s := []string{"aku", "makan", "kuo", "kau", "muka", "kamu"}
	var res = make([]string, 0)
	var result = make([][]string, 0)

	for _, sub := range s { // O(n)
		r := []rune(sub)
		sort.Sort(subS(r)) // O(n*log(n))
		res = append(res, string(r))
	}

	var arrLMap = make(map[string][]int)

	for i := 0; i < len(res); i++ {
		arrLMap[res[i]] = append(arrLMap[res[i]], i)
	}

	for _, index := range arrLMap {
		arr := []string{}
		for _, i := range index {
			arr = append(arr, s[i])
		}
		result = append(result, arr)

	}

	fmt.Println(result)
}
