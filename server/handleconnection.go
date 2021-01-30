package server

import (
	"bufio"
	"chatApp/errorhandler"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var requestMap = make(map[string]string)

// HandleConnection handles the request made by the client
func HandleConnection(conn net.Conn) {
	fmt.Print("handling connection\n")
	fmt.Print("conn local address: ", conn.LocalAddr(), "\n")

	isValid := readerRequest(conn)
	if isValid {
		fmt.Print("request map after it has been set by the request: ", requestMap, "\n")
		PickResponse(requestMap, conn)

	} else {
		conn.Close()
	}
}

func readerRequest(conn net.Conn) bool {
	var err bool = false
	var isValid = false
	fmt.Print("reading client request\n")
	reader := bufio.NewReader(conn)
	arrayNumberOfBytesInRequest, readNumberOfBytesInRequestError := reader.ReadBytes('\n')
	err = errorhandler.RequestError(readNumberOfBytesInRequestError, "handleconnection.readerRequest() reading number of bytes", conn)

	if !err {
		fmt.Print("Number of bytes before conversion to int: ", string(arrayNumberOfBytesInRequest), "\n")

		strNumberOfBytesInRequest := string(arrayNumberOfBytesInRequest)
		trimmedStrNumberOfBytesInRequest := strings.TrimSpace(strNumberOfBytesInRequest)

		intNumberOfBytesInRequest, numberOfBytesError := strconv.ParseInt(trimmedStrNumberOfBytesInRequest, 10, 64)
		err = errorhandler.RequestError(numberOfBytesError, "handleconnection.readerRequest()", conn)
		if !err {
			fmt.Print("Number of bytes is string: ", intNumberOfBytesInRequest, "\n")

			var requestBytes []byte = make([]byte, intNumberOfBytesInRequest)
			_, readRestOfBytesError := reader.Read(requestBytes)
			err = errorhandler.RequestError(readRestOfBytesError, "handleconnection.readerRequest()", conn)

			if !err {
				fmt.Print("rest of message", string(requestBytes), "\n")

				var decrytpedMessage string = Decrypt(string(requestBytes))

				fmt.Print("decrypted message: \n", decrytpedMessage, "\n")

				isValid = ParseRequestMap(decrytpedMessage, requestMap)
			}

		}

	}
	return isValid
}
