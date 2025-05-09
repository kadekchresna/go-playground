package main

import (
	"encoding/base64"
	"fmt"
	"time"
)

var shortURL = map[string]string{}

func main() {

	// a := []string{"a", "b", "c", "d", "e"}
	// b := []string{"f", "g", "h", "i", "j"}

	// fmt.Println(append(a, ...b))

	genShortURL("https://google.com")
	fmt.Println(shortURL)

}

func genShortURL(url string) string {

	num := time.Now().Unix() + int64(len(shortURL)+1)
	shortLink := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", num)))

	shortURL[shortLink] = url
	return shortLink
}
