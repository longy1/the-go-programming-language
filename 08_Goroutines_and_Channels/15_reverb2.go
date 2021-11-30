package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:2021")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go reverbConn2(conn)
	}
}

func reverbConn2(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Print(err)
		}
	}(conn)
	scan := bufio.NewScanner(conn)
	scan.Split(bufio.ScanLines)

	var wg sync.WaitGroup
	for scan.Scan() {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			echo2(conn, text, 1236*time.Millisecond)
		}(scan.Text())
	}

	// closer
	wg.Wait()
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		if err := tcpConn.CloseWrite(); err != nil {
			log.Print(err)
		}
	}
}

func echo2(conn net.Conn, text string, delay time.Duration) {
	time.Sleep(delay)
	if _, err := fmt.Fprintf(conn, "  %s~~~\n", strings.ToUpper(text)); err != nil {
		log.Print(err)
		return
	}
	time.Sleep(delay)
	if _, err := fmt.Fprintf(conn, "  %s~~\n", text); err != nil {
		log.Print(err)
		return
	}
	time.Sleep(delay)
	if _, err := fmt.Fprintf(conn, "  %s~\n", strings.ToLower(text)); err != nil {
		log.Print(err)
		return
	}
}
