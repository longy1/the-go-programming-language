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

var sem4 = make(chan struct{}, 20)
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done: // will return immediately after closing done
		return true
	default:
		return false
	}
}

func walkDir4(dir string, waitGroup *sync.WaitGroup, fileSizes chan<- int64) {
	defer waitGroup.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirEnt4(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			waitGroup.Add(1)
			go walkDir4(subDir, waitGroup, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirEnt4(dir string) []os.FileInfo {
	select {
	case <-done:
		return nil
	case sem4 <- struct{}{}:
	}
	defer func() { <-sem4 }()
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
	// canceler, trigger when press enter
	go func() {
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
		}
		close(done)
	}()

	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, dir := range args {
		wg.Add(1)
		go walkDir4(dir, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range fileSizes {
			} // do nothing to clear fileSizes
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		}
	}
	printDiskUsage4(nfiles, nbytes)
}

func printDiskUsage4(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
