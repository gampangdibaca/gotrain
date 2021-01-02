package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("error fmt happened")
		log.Println("err log happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
	defer f2.Close()
	//log.Println to a file
}
