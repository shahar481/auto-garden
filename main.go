package main

import (
	"auto-garden/argparsing"
	"auto-garden/cli"
	"auto-garden/web_server"
)


func main() {
	a := argparsing.ParseArgs()
	cli.StartCli()
	web_server.StartServer(a)
}
