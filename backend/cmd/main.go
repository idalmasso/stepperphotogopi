package main

import "github.com/idalmasso/stepperphotogopi/backend/server"

func main() {
	server := server.PiServer{}
	server.Init()
	server.StartAndListen()
}
