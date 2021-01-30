package server

import (
	"fmt"
	"net"
	"strconv"
)

func sendOk(conn net.Conn) {
	var response string = "OK"
	conn.Write([]byte(response))
}

func sendNotOk(conn net.Conn) {
	var response string = "NOT OK"
	conn.Write([]byte(response))
	conn.Close()
}

func sendPingResponse(conn net.Conn) {
	var finish string = "done\n"
	conn.Write([]byte(finish))
	conn.Close()
}

//SendRegistrationResponse send response to a registration request from the client
func SendRegistrationResponse(conn net.Conn) {
	var requestType string = "requestType:registration\n"

	var unencryptedResponse string = authTokenStr + requestType + successResponse + finishedTrue
	encryptedResponse := Encrypt(unencryptedResponse)
	var numberOfBytesInResponse int = len(encryptedResponse)
	fmt.Print("number of bytes in response: ", numberOfBytesInResponse, "\n")
	var strNumberOfBytesInResponse string = strconv.Itoa(numberOfBytesInResponse)
	fmt.Print("str Number Of Bytes In Response: ", strNumberOfBytesInResponse, "\n")

	var response = strNumberOfBytesInResponse + "\n" + encryptedResponse

	fmt.Print("This is the encrypted response from the server: ", response, "\n")

	conn.Write([]byte(response))
	conn.Close()

}
