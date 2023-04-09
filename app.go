package main

import (
	"log"

	"github.com/jplein/chatbot/config"
)

func main() {
	var err error

	var f config.File
	if err = f.Init(); err != nil {
		log.Fatal(err)
	}
}
