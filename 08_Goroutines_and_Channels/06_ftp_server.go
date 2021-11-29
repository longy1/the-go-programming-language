package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var sessions map[net.Addr]Session

type Session struct {
	dataCon net.Conn
}

func main() {
	listener, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handlerFTPConn(conn)
	}
}

// handlerFTPConn call dispatch to delivery command, defer clear session
// and close connect
func handlerFTPConn(conn net.Conn) {
	defer func(conn net.Conn) {
		// clear session
		if session, ok := sessions[conn.RemoteAddr()]; !ok {
			err := session.dataCon.Close()
			if err != nil {
				log.Print(err)
			}
		}
		delete(sessions, conn.RemoteAddr())
		log.Print(fmt.Sprintf("[Close] RemoteAddr: %v, connection closed\n", conn.RemoteAddr()))
		err := conn.Close()
		if err != nil {
			log.Print(err)
		}
	}(conn)
	dispatch(conn)
}

// dispatch delivery client's command
func dispatch(conn net.Conn) {
	var input = bufio.NewScanner(conn)
	input.Split(bufio.ScanLines)

	for input.Scan() {
		split := strings.Split(strings.TrimSpace(input.Text()), " ")
		switch len(split) {
		// do nothing
		case 0:
			continue
		// unary op
		case 1:
			switch strings.TrimSpace(split[0]) {
			case "ls":
				listToClient(conn)
			case "close":
				break
			}
		// binary op
		case 2:
			switch strings.TrimSpace(split[0]) {
			case "get":
				serverHandleGet(conn, strings.TrimSpace(split[1]))
			case "send":
				serverHandleSend()
			case "cd":
				serverHandleCd()
			}
		// unknown op
		default:
			_, err := io.WriteString(conn, fmt.Sprintf("Unknown command: %v\n", input.Text()))
			if err != nil {
				log.Print(err)
				return
			}
		}
	}
}

func listToClient(conn net.Conn) {
	// todo listToClient
}

func serverHandleGet(conn net.Conn, p string) {
	// todo serverHandleGet
}

func serverHandleSend() {
	// todo serverHandleSend
}

func serverHandleCd() {
	// todo serverHandleCd
}
