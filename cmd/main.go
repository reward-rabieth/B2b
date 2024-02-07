package main

import (
	"github.com/reward-rabieth/b2b/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
