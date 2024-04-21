package main

import (
	"bufio"
	"log"
	"os"

	"github.com/lallenfrancisl/lsp-from-scratch/rpc"
)

func main() {
	logger := getLogger("/home/allen/projekts/lsp-from-scratch/git/log.txt")
	logger.Println("LSP started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("Could not create a log file")
	}

	return log.New(logfile, "[educationlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
