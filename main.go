package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/lallenfrancisl/lsp-from-scratch/analysis"
	"github.com/lallenfrancisl/lsp-from-scratch/lsp"
	"github.com/lallenfrancisl/lsp-from-scratch/rpc"
)

func main() {
	logger := getLogger("/home/allen/projekts/lsp-from-scratch/gitignored/log.txt")
	logger.Println("LSP started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()

	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)

		if err != nil {
			logger.Printf("Error happened while decoding request: %s", err)

			continue
		}

		handleMessage(logger, &state, writer, method, contents)
	}
}

func handleMessage(
	logger *log.Logger,
	state *analysis.State,
	writer io.Writer,
	method string, contents []byte,
) {
	logger.Printf("Received msg with method: %s\n", method)

	if method == "initialize" {
		var request lsp.InitialiseRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Couldn't parse the initialise request: %s", err)
		} else {
			logger.Printf(
				"Connected to %s, version %s",
				request.Params.ClientInfo.Name,
				request.Params.ClientInfo.Version,
			)

			msg := lsp.NewInitialiseResponse(request.ID)
			writeMessage(writer, msg)

			logger.Println("Sent initialise reply")
		}
	} else if method == "textDocument/didOpen" {
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didOpen: %s", err)
		} else {
			logger.Printf(
				"Opened: %s",
				request.Params.TextDocument.URI,
			)
			state.OpenDocument(
				request.Params.TextDocument.URI,
				request.Params.TextDocument.Text,
			)
		}
	} else if method == "textDocument/didChange" {
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
		} else {
			logger.Printf(
				"Changed: %s",
				request.Params.TextDocument.URI,
			)

			for _, change := range request.Params.ContentChanges {
				state.UpdateDocument(
					request.Params.TextDocument.URI,
					change.Text,
				)
			}
		}
	} else if method == "textDocument/hover" {
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
		} else {
			logger.Printf(
				"Hovered: %s",
				request.Params.TextDocument.URI,
			)

			response := state.Hover(
				request.ID,
				request.Params.TextDocument.URI,
				request.Params.Position,
			)
			writeMessage(writer, response)
		}
	} else if method == "textDocument/definition" {
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/definition: %s", err)
		} else {
			logger.Printf(
				"Definition: %s",
				request.Params.TextDocument.URI,
			)

			response := state.Definition(
				request.ID,
				request.Params.TextDocument.URI,
				request.Params.Position,
			)
			writeMessage(writer, response)
		}
	} else if method == "textDocument/codeAction" {
		var request lsp.TextDocumentCodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/codeAction: %s", err)
		} else {
			logger.Printf(
				"CodeAction: %s",
				request.Params.TextDocument.URI,
			)

			response := state.CodeAction(
				request.ID,
				request.Params.TextDocument.URI,
				request.Params.Range,
			)
			writeMessage(writer, response)
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

func writeMessage(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}
