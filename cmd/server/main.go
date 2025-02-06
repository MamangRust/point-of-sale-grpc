package main

import "pointofsale/internal/app"

func main() {
	server, err := app.NewServer()

	if err != nil {
		panic(err)
	}

	server.Run()
}
