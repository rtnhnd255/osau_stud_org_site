package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rtnhnd255/osau_stud_org_site.git/app/src/server"
)

var (
	cfgPath string
)

func init() {
	flag.StringVar(&cfgPath, "config", "./server-config.json", "path to server config")
}

func main() {
	c, err := server.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(c)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Hello world")
}
