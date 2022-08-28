package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type TelnetServer struct {
	Listener net.Listener
}

func (t *TelnetServer) RunServer() {
	listener, err := net.Listen("tcp", ":8020")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	t.Listener = listener

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		for {
			msg, err := bufio.NewReader(con).ReadString('\n')
			if err != nil {
				fmt.Println("connection end")
				break
			}

			trim := strings.TrimSuffix(msg, "\n")
			fmt.Println(trim)
			fmt.Fprintf(con, "message recieved: %s", msg)
		}
	}
}
