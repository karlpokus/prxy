package main

import (
	"log"
	"os"

	"prxy"
)

func main() {
	log.Println("prxy starting")
	err := prxy.Start(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	log.Println("prxy exited")
}
