package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func toJPEG(r io.Reader, w io.Writer) error {
	img, kind, err := image.Decode(r)
	if err != nil {
		return err
	}
	fmt.Println("Input format:", kind)
	return jpeg.Encode(w, img, &jpeg.Options{Quality: 95})
}
