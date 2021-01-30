package server

import (
	"fmt"
	"net"
)

//PickResponse is how the correct response in choosen
func PickResponse(requestMap map[string]string, conn net.Conn) {
	fmt.Print("picking response\n")

	switch requestMap["requestType"] {
	case "ping":
		break
	case "registration":
		fmt.Print("Sending Registration Response\n")
		SendRegistrationResponse(conn)
		break
	case "processMessage":
		break
	case "loadChat":
		break
	default:
		break
	}
}
