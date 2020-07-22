package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const totalFile = 20000
const contentLength = 1000

var tempPath = filepath.Join("../", "pipeline-temp")

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	log.Println("Start")

	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println(" crated for ", duration.Seconds())

}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := randomString(contentLength)
		err := ioutil.WriteFile(filename, []byte(content), os.ModePerm)

		if err != nil {
			log.Printf("error create file %s\n", err)
		}

		if i%100 == 0 && i > 0 {
			log.Printf("%d has been created\n", i)
		}

	}
	log.Printf("%d created\n", totalFile)
}
