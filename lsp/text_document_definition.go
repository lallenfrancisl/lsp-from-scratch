package lsp

type DefinitionRequest struct {
	Request
	Params HoverParams `json:"params"`
}

type DefinitionParams struct {
	TextDocumentPositionParams
}

type DefinitionResponse struct {
	Response
	Result Location `json:"result"`
}
