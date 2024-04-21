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
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)

		if err != nil {
			logger.Printf("Error happened while decoding request: %s", err)

			continue
		}

		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with method: %s\n", method)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("Could not create a log file")
	}

	return log.New(logfile, "[educationlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
