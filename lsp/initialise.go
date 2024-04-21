package lsp

type InitialiseRequest struct {
	Request
	Params InitialiseRequestParams `json:"params"`
}

type InitialiseRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitiliseResponse struct {
	Response
	Result InitialiseResult `json:"result"`
}

type InitialiseResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync int `json:"textDocumentSync"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitialiseResponse(id int) InitiliseResponse {
	return InitiliseResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitialiseResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
			},
			ServerInfo: ServerInfo{
				Name:    "educational_lsp",
				Version: "0.0.1",
			},
		},
	}
}
