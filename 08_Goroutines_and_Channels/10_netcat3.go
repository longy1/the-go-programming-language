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
	done := make(chan struct{})
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Print(err)
		}
		log.Println("Done")
		if tcpConn, ok := conn.(*net.TCPConn); ok {
			if err := tcpConn.CloseRead(); err != nil {
				log.Print(err)
			}
		}
		done <- struct{}{}
	}()
	if _, err := io.Copy(conn, os.Stdin); err != nil {
		log.Print(err)
	}
	// duplex close
	// if err := conn.Close(); err != nil {
	// 	log.Print(err)
	// }

	// only close write
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		if err := tcpConn.CloseWrite(); err != nil {
			log.Print(err)
		}
	}
	<-done
}
