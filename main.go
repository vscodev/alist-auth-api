package main

import (
	"flag"

	"github.com/vscodev/alist-auth-api/api"
)

var (
	_host string
	_port int
)

func init() {
	flag.StringVar(&_host, "host", "0.0.0.0", "listen host")
	flag.IntVar(&_port, "port", 5243, "listen port")
}

func main() {
	flag.Parse()

	api.Run(_host, _port)
}
