package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(group_anagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func group_anagrams(input []string) [][]string {
	// []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	var res = [][]string{}

	var resMap = map[string][]string{}
	for _, v := range input {

		arrString := strings.Split(v, "")
		sort.Strings(arrString)
		joinString := strings.Join(arrString, "")

		resMap[joinString] = append(resMap[joinString], v)

	}

	for _, v := range resMap {
		res = append(res, v)
	}

	return res
}
