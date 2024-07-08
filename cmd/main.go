package main

import "templtest/internal/transport/web"

func main() {
	server := web.New()
	server.Serve()
}
