package main

import (
	"chatApp/dbmanager"
)

func main() {

	// go server.CreateListeningSocket("5124", "localhost")
	// go server.CreateListeningSocket("5125", "localhost")

	// for !server.EndOfLife {

	// }
	dbmanager.CreateNewDbSession()
}
