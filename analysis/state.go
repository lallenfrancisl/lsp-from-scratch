package analysis

import (
	"fmt"
	"strings"

	"github.com/lallenfrancisl/lsp-from-scratch/lsp"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri string, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri string, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf(
				"file: %s, content length: %d",
				uri,
				len(document),
			),
		},
	}
}

func (s *State) Definition(
	id int, uri string, position lsp.Position,
) lsp.DefinitionResponse {
	definitionLine := position.Line - 1
	if definitionLine < 0 {
		definitionLine = 0
	}

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      definitionLine,
					Character: 0,
				},
				End: lsp.Position{
					Line:      definitionLine,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) CodeAction(
	id int, uri string, selectedRange lsp.Range,
) lsp.TextDocumentCodeActionResponse {
	document := s.Documents[uri]
	var response lsp.TextDocumentCodeActionResponse

	lines := strings.Split(document, "\n")
	line := lines[selectedRange.Start.Line]

	changes := map[string][]lsp.TextEdit{}
	changes[uri] = []lsp.TextEdit{
		{
			Range: lsp.Range{
				Start: lsp.Position{
					Line: selectedRange.Start.Line,
					Character: 0,
				},
				End: lsp.Position{
					Line: selectedRange.Start.Line,
					Character: len(line),
				},
			},
			NewText: strings.ToUpper(line),
		},
	}

	response = lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: []lsp.CodeAction{
			{
				Title: "Change line to uppercase",
				Edit: &lsp.WorkspaceEdit{
					Changes: changes,
				},
			},
		},
	}

	return response
}
