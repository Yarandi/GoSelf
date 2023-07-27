package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	t := time.Now().Add(20 * time.Second)
	err := conn.SetReadDeadline(t) //the connection will close after 20 seconds
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say:%s\n", ln)
	}
	defer conn.Close()
	//now we get here
	//the connection will timeout
	//that breaks us out of the scanner loop
	fmt.Printf("Connection was closed at %v", t)
}
