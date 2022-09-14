package main

import (
	controller "snapp-trip/controller"
)

func main() {
	controller.ConnectToPostgre()
	controller.InitializeTables()
}