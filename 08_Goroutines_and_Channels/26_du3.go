package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var sem = make(chan struct{}, 20)

func walkDir3(dir string, waitGroup *sync.WaitGroup, fileSizes chan<- int64) {
	sem <- struct{}{}
	defer func() { <-sem }()
	defer waitGroup.Done()
	for _, entry := range dirEnt3(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			waitGroup.Add(1)
			go walkDir3(subDir, waitGroup, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirEnt3(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		if err != nil {
			log.Println(err)
		}
		return nil
	}
	return entries
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, dir := range args {
		wg.Add(1)
		go walkDir3(dir, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage3(nfiles, nbytes)
}

func printDiskUsage3(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
