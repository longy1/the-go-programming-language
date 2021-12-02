// TCP chatroom
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	ch   chan<- string
	name string
}

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
			// announce to new incoming client
			var curClients []string
			for cli := range clients {
				curClients = append(curClients, cli.name)
			}
			cli.ch <- fmt.Sprintf("Welcome, %q in chatroom", curClients)

			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		case msg := <-messages:
			for cli := range clients {
				go timeoutSend(cli.ch, msg)
			}
		}
	}
}

// timeoutSend will simply drop message if client timeout
func timeoutSend(dst chan<- string, msg string) {
	const SendTimeout = 1 * time.Second
	timeout := time.NewTimer(SendTimeout)
	defer func() { timeout.Stop() }() // terminate Timer goroutine

	select {
	case dst <- msg:
	case <-timeout.C:
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
	defer func() {
		if err := conn.Close(); err != nil {
			// do nothing
		}
	}()
	cli := make(chan string)
	messageMid := make(chan string)
	go clientWriter(conn, cli)
	go timeoutMessageCheck(conn, messageMid)

	who := conn.RemoteAddr().String()
	messageMid <- fmt.Sprintf("[%s] %s has arrived",
		time.Now().Format("Mon Jan 2 15:04:05"), who)
	entering <- client{cli, who}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		messageMid <- fmt.Sprintf("[%s] %s: %s",
			time.Now().Format("Mon Jan 2 15:04:05"), who, scanner.Text())
	}

	leaving <- client{cli, who}
	messageMid <- fmt.Sprintf("[%s] %s has left",
		time.Now().Format("Mon Jan 2 15:04:05"), who)
}

func clientWriter(conn net.Conn, cli <-chan string) {
	for msg := range cli {
		if _, err := fmt.Fprintln(conn, msg); err != nil {
			log.Println(err)
		}
	}
}

// timeoutMessageCheck will Close(conn) when timeout happens
func timeoutMessageCheck(conn net.Conn, src chan string) {
	const MessageTimeout = 10 * time.Second
	timeout := time.NewTimer(MessageTimeout)
	defer func() { timeout.Stop() }() // terminate Timer goroutine

	for {
		select {
		case msg := <-src:
			// safety reset timer
			if !timeout.Stop() {
				<-timeout.C
			}
			timeout.Reset(MessageTimeout)
			messages <- msg
		case <-timeout.C:
			if err := conn.Close(); err != nil {
				log.Println(err)
			}
			return
		}
	}
}
