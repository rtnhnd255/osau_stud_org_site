package main

import "fmt"

func main() {
	s := server.NewServer()
	s.Start()

	fmt.Print("Hello world")
}
