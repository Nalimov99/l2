package main

import "dev10/server"

func main() {
	ts := server.TelnetServer{}
	ts.RunServer()
}
