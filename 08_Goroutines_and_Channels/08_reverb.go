package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
		go reverbConn(conn)
	}
}

func reverbConn(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Print(err)
		}
	}(conn)
	scan := bufio.NewScanner(conn)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		go echo(conn, scan.Text(), 1236*time.Millisecond)
	}
}

func echo(conn net.Conn, text string, delay time.Duration) {
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
