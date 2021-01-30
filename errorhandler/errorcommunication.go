package errorhandler

import (
	"fmt"
	"net"
)

//RequestError handles the closing of the client connection if there is an error
func RequestError(err error, functionName string, conn net.Conn) bool {
	var wasError bool = false
	if err != nil {
		wasError = true
		fmt.Print("ERROR in function: ", functionName, " the error was: ", err, "\n")
		conn.Close()
	}
	return wasError

}
