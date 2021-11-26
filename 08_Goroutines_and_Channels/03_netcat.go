package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2021")
	if err != nil {
		log.Fatal(err)
	}
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Print(err)
		}
	}(conn)

	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
