package main

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
}

func main() {
	Info.Println("Info...")
	Error.Println("Error...")
	log.Println("Start log demo...")
	//log.Panicln("Panic")
	log.Fatalln("Fatal...")
	println("Never executed...")
}
