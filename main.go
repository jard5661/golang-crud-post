package main

import (
	"golang-test/connection"
	"golang-test/handlers"
)

func main() {
	connection.Connect()

	handlers.HandleReq()
}
