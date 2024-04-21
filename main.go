package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/lallenfrancisl/lsp-from-scratch/lsp"
	"github.com/lallenfrancisl/lsp-from-scratch/rpc"
)

func main() {
	logger := getLogger("/home/allen/projekts/lsp-from-scratch/gitignored/log.txt")
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

	if method == "initialize" {
		var request lsp.InitialiseRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Couldn't parse the initialise request: %s", err)
		} else {
			logger.Printf(
				"Received message with method: %s",
				request.Method,
			)
			logger.Printf(
				"Connected to %s, version %s",
				request.Params.ClientInfo.Name,
				request.Params.ClientInfo.Version,
			)

			msg := lsp.NewInitialiseResponse(request.ID)
			reply := rpc.EncodeMessage(msg)

			writer := os.Stdout
			writer.Write([]byte(reply))
			logger.Println("Sent initialise reply")
		}
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("Could not create a log file")
	}

	return log.New(logfile, "[educationlsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
