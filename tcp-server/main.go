package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I heard you say : %s\n", line)
	}
	defer conn.Close()
	fmt.Println("Code here")
}
