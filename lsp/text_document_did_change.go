package lsp

type TextDocumentDidChangeNotification struct {
	Notification
	Params DidChangeNotificationParams `json:"params"`
}

type DidChangeNotificationParams struct {
	TextDocument   VersionTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}
