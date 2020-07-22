package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	tempPath = filepath.Join("../", "pipeline-temp")
)

func main() {

	log.Println("rename start")
	start := time.Now()

	rename()

	duration := time.Since(start)
	log.Println("finished for", duration.Seconds(), "seconds")

}

func rename() {
	countRenamed := 0
	countTotal := 0

	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		countTotal++
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		if err == io.EOF {
			log.Println("EOF ", err)
		}
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		destination := filepath.Join(tempPath, fmt.Sprintf("file-%s-%d.txt", sum, countRenamed))
		errRename := os.Rename(path, destination)
		if errRename != nil {
			return err
		}
		countRenamed++
		return nil
	})

	if err != nil {
		log.Println("err ", err)
	}

	log.Printf("%d/%d files renamed", countRenamed, countTotal)
}
