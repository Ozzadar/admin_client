package main 

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	//connect to this socket
	conn, _ := net.Dial("tcp", "192.168.100.104:1337")

	
	//read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("What is your username?" )
		name, _ := reader.ReadString('\n')
		conn.Write([]byte(name)[0:len(name)-1])

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