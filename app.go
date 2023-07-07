package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jplein/chatbot/chat"
	"github.com/jplein/chatbot/config"
	"github.com/jplein/chatbot/storage"
)

func main() {
	var err error

	var dir storage.Dir
	if err = dir.Init(); err != nil {
		log.Fatal(err)
	}

	var f config.File
	f.StorageDir = dir

	if err = f.Init(); err != nil {
		log.Fatal(err)
	}

	var c config.Config
	if c, err = f.Read(); err != nil {
		log.Fatal(err)
	}
	if c.Model == "" {
		c.Model = chat.DefaultChatModel
	}

	msg := strings.Join(os.Args[1:], " ")

	var response string
	var tokenUsage int
	if response, tokenUsage, err = chat.Send(&dir, c.APIKey, msg, c.Model); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", response)

	if c.LogTokenUsage {
		fmt.Printf("(tokens used: %d)\n", tokenUsage)
	}
}
