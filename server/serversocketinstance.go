package server

import (
	"fmt"
	"net"
	"os"
)

//CreateListeningSocket creates new serversocket that clients can use
func CreateListeningSocket(port string, host string) {
	fmt.Print("Creating new socket where client can connect\n")
	var url string = host + ":" + port
	serverSocket, socketErr := net.Listen("tcp", url)
	if socketErr != nil {
		fmt.Println("Error listening:", socketErr.Error())
		os.Exit(1)
	}
	for {
		conn, connErr := serverSocket.Accept()

		if connErr != nil {
			fmt.Println("Error listening:", connErr.Error())
			os.Exit(1)
		}
		fmt.Print("This is the socket that was used this time, Local Address:", conn.LocalAddr(), " Remote address: ", conn.RemoteAddr())
		go HandleConnection(conn)
	}
}
