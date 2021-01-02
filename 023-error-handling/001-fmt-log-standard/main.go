package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//printing & logging
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("error fmt happened", err)
		log.Println("err log happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
	//println use standard output
	//ex : error fmt happened open no-file.txt: The system cannot find the file specified.
	//log/Println calls output to print to standard logger
	//ex : 2020/12/27 15:03:54 err log happened open no-file.txt: The system cannot find the file specified.
}
