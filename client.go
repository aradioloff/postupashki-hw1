package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
)

func main(){
	cnct, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer cnct.Close()
	ln, err := bufio.NewReader(cnct).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if ln != "OK\n" {
		log.Fatalf("unexpected response: %q", ln)
	}
	fmt.Println(ln)
}
