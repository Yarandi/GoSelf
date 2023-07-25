package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	//---------------------------
	//split by \n
	//---------------------------
	//s := "Hello World!\nHows everything?\nSo Good!!!"
	//scanner := bufio.NewScanner(strings.NewReader(s))
	//
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	fmt.Println(line)
	//}

	//-----------------------
	//split by words
	//-----------------------
	//s := "Hello World!\nHows everything?\nSo Good!!!"
	//scanner := bufio.NewScanner(strings.NewReader(s))
	//scanner.Split(bufio.ScanWords)
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	fmt.Println(line)
	//}

	//-----------------------
	//split by characters
	//-----------------------
	//s := "Hello World!\nHows everything?\nSo Good!!!"
	//scanner := bufio.NewScanner(strings.NewReader(s))
	//scanner.Split(bufio.ScanRunes)
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	fmt.Println(line)
	//}

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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	//we never got here
	fmt.Println("code got here")
}
