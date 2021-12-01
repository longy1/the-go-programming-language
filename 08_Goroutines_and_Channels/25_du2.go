package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func walkDir2(dir string, fileSizes chan<- int64) {
	for _, entry := range dirEnt2(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir2(subDir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirEnt2(dir string) []os.FileInfo {
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
	go func() {
		for _, dir := range args {
			walkDir2(dir, fileSizes)
		}
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage2(nfiles, nbytes)
		}
	}
	printDiskUsage2(nfiles, nbytes)
}

func printDiskUsage2(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
