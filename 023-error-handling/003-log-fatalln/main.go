package main

import (
	"log"
	"os"
)

func main() {
	//printing & logging
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("error fmt happened", err)
		//log.Println("err log happened", err)
		log.Fatalln(err)
		//panic(err)
	}
	//log.Fatalln is equivalent to Println() followed by a call to os.Exit(1)
}
