package main

import (
	"flag"
	"log"

	"github.com/rtnhnd255/osau_stud_org_site.git/server"
	"github.com/rtnhnd255/osau_stud_org_site.git/storage"
)

var (
	serverConfigPath  string
	storageConfigPath string
)

func init() {
	flag.StringVar(&serverConfigPath, "servc", "server-config.json", "path to server config")
	flag.StringVar(&storageConfigPath, "storc", "storage-config.json", "path to storage config")
}

func main() {
	storc, err := storage.NewConfig(storageConfigPath)
	if err != nil {
		log.Println("Trouble opening storage config")
		log.Fatal(err)
	}
	servc, err := server.NewConfig(serverConfigPath, storc)
	if err != nil {
		log.Println("Trouble opening server config")
		log.Fatal(err)
	}

	s := server.NewServer(servc)
	if err := s.Run(); err != nil {
		log.Println("trouble with server init")
		log.Fatal(err)
	}

	if err = s.Run(); err != nil {
		log.Println("Error while starting server")
		log.Fatal(err)
	}
}
