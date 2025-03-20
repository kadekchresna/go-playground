package main

import (
	"fmt"
	"strings"
)

type Player struct {
	Name  string
	Score int
}

func main() {

	p := []Player{
		Player{
			Name:  "aleksa",
			Score: 150,
		},
		Player{
			Name:  "amy",
			Score: 100,
		},
		Player{
			Name:  "david",
			Score: 100,
		},
		Player{
			Name:  "aakansha",
			Score: 75,
		},
		Player{
			Name:  "alan",
			Score: 75,
		},
		Player{
			Name:  "heraldo",
			Score: 50,
		},
		Player{
			Name:  "ass",
			Score: 60,
		},
	}

	fmt.Println(MergeSort(p))

}

// func compare(p []Player) []Player {

// 	for {

// 	}

// }

func helper(arrLeft []Player, arrRight []Player) []Player {
	var result []Player
	var i = 0
	var j = 0

	for i < len(arrLeft) && j < len(arrRight) {
		if arrLeft[i].Score == arrRight[j].Score {
			if strings.ToLower(arrLeft[i].Name) > strings.ToLower(arrRight[j].Name) {
				result = append(result, arrRight[j])
				j++
			} else {
				result = append(result, arrLeft[i])
				i++
			}

		} else if arrLeft[i].Score < arrRight[j].Score {
			result = append(result, arrRight[j])
			j++
		} else {
			result = append(result, arrLeft[i])
			i++
		}
	}

	if i < len(arrLeft) {
		result = append(result, arrLeft[i:]...)
	}
	if j < len(arrRight) {
		result = append(result, arrRight[j:]...)
	}
	return result
}

func MergeSort(arr []Player) []Player {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:])
	return helper(left, right)

}
