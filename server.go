package main

import (
	"net"
	"log"
)

func main() {
	lstn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lstn.Close()

	for {
		cnct, err := lstn.Accept()
		if err != nil {
			log.Println("accept:", err)
			continue
		}
		handle(cnct)
	}
}

func handle(cnct net.Conn) {
	defer cnct.Close()
	if _, err := cnct.Write([]byte("OK\n")); err != nil {
		log.Println("write:", err)
}
}
