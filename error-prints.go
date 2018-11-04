package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	_,err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
		log.Println("log err: ", err)
		log.Fatalln("log Fatal err:", err)
		log.Panicln("log Panic err", err)
	}
}