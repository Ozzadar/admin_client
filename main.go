package main 

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	//connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:1337")

	
	//read in input from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Oasis!" )

	go ListenForResponse(conn)

	for {
		text, _ := reader.ReadString('\n')

		//send to socket
		conn.Write([]byte(text)[0:len(text)-1])

	}
}

func ListenForResponse(conn net.Conn) {

	for {
		//listen for reply
		buffer := make([]byte,1024)
		conn.Read(buffer)

		fmt.Println(string(buffer))
	}
}