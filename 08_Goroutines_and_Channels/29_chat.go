package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcast() {
	clients := make(map[client]bool) // set of existing clients
	for {
		select {
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:2021")
	if err != nil {
		log.Fatal(err)
	}

	go broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleChatConn(conn)
	}
}

func handleChatConn(conn net.Conn) {
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}(conn)
	cli := make(chan string)
	go clientWriter(conn, cli)

	who := conn.RemoteAddr().String()
	messages <- fmt.Sprintf("[%s] %s has arrived",
		time.Now().Format("Mon Jan 2 15:04:05"), who)
	entering <- cli

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		messages <- fmt.Sprintf("[%s] %s: %s",
			time.Now().Format("Mon Jan 2 15:04:05"), who, scanner.Text())
	}

	leaving <- cli
	messages <- fmt.Sprintf("[%s] %s has left",
		time.Now().Format("Mon Jan 2 15:04:05"), who)
}

func clientWriter(conn net.Conn, cli <-chan string) {
	for msg := range cli {
		if _, err := fmt.Fprintln(conn, msg); err != nil {
			log.Println(err)
		}
	}
}
