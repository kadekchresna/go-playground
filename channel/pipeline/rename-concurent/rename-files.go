package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileInfo struct {
	FilePath  string // file location
	Content   []byte // file content
	Sum       string // md5 sum of content
	IsRenamed bool   // indicator whether the particular file is renamed already or not
}

var (
	tempPath = filepath.Join("../", "pipeline-temp")
)

func main() {

	log.Println("rename start")
	start := time.Now()

	chanFileContent := readFile()

	chanFileSum1 := getSum(chanFileContent)
	chanFileSum2 := getSum(chanFileContent)
	chanFileSum3 := getSum(chanFileContent)
	chanFileSum := mergeChan(chanFileSum1, chanFileSum2, chanFileSum3)

	chanRename1 := rename(chanFileSum)
	chanRename2 := rename(chanFileSum)
	chanRename3 := rename(chanFileSum)
	chanRename4 := rename(chanFileSum)
	chanRename := mergeChan(chanRename1, chanRename2, chanRename3, chanRename4)

	// print output
	counterRenamed := 0
	counterTotal := 0
	for fileInfo := range chanRename {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)

	duration := time.Since(start)
	log.Println("finished for", duration.Seconds(), "seconds")

}

func readFile() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			if err == io.EOF {
				log.Println("EOF ", err)
			}
			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}
			return nil

		})

		if err != nil {
			log.Println("error when reading ", err)
		}
		close(chanOut)

	}()

	return chanOut
}
func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for val := range chanIn {
			val.Sum = fmt.Sprintf("%x", md5.Sum(val.Content))
			chanOut <- val
		}
		close(chanOut)
	}()
	return chanOut

}

func mergeChan(chans ...<-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	var wg sync.WaitGroup
	wg.Add(len(chans))
	for _, chanE := range chans {
		go func(chanIn <-chan FileInfo) {
			for chanData := range chanIn {
				chanOut <- chanData
			}
			wg.Done()
		}(chanE)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut

}

func rename(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for val := range chanIn {
			newPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", val.Sum))
			err := os.Rename(val.FilePath, newPath)
			val.IsRenamed = err == nil
			chanOut <- val
		}
		close(chanOut)
	}()

	return chanOut
}
