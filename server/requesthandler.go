// https://github.com/mervick/aes-everywhere

package server

import (
	"fmt"
	"strings"
)

//ParseRequestMap parses the clients request into a map object
func ParseRequestMap(request string, requestMap map[string]string) bool {
	var isValidRequest bool = false
	var jsonObjArr []string = strings.Split(request, "\n")
	for i := 0; i < len(jsonObjArr); i++ {
		var jsonObj string = jsonObjArr[i]
		var keyAndValue []string = strings.Split(jsonObj, ":")
		fmt.Print(keyAndValue, "\n")
		if len(keyAndValue) == 2 {
			requestMap[keyAndValue[0]] = keyAndValue[1]
		}
	}
	fmt.Print("the request map: \n", requestMap, "\n")

	token := requestMap["authToken"]
	fmt.Print("Token: ", token, "\n")
	isValidRequest = verifyRequest(requestMap["authToken"])
	fmt.Print("isValidRequest: ", isValidRequest, "\n")
	return isValidRequest
}

func verifyRequest(token string) bool {
	fmt.Print("verifing request, token: ", token, " authToken: ", authToken, "\n")
	var isOk bool = false
	var sameAuthTok int = strings.Compare(authToken, token)
	if sameAuthTok == 0 {
		isOk = true
	}
	return isOk
}
