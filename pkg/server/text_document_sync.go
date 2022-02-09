package server

import (
	"context"
	"fmt"

	"go.lsp.dev/protocol"

	"github.com/tilt-dev/starlark-lsp/pkg/analysis"
	"github.com/tilt-dev/starlark-lsp/pkg/document"
)

func (s *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	uri := params.TextDocument.URI
	contents := []byte(params.TextDocument.Text)
	tree, err := analysis.Parse(ctx, contents)
	if err != nil {
		return fmt.Errorf("could not parse file %q: %v", uri, err)
	}

	doc := document.NewDocument(contents, tree)
	s.docs.Write(uri, doc)
	return nil
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	if len(params.ContentChanges) == 0 {
		return nil
	}

	uri := params.TextDocument.URI
	contents := []byte(params.ContentChanges[0].Text)
	tree, err := analysis.Parse(ctx, contents)
	if err != nil {
		s.docs.Remove(uri)
		return fmt.Errorf("could not parse file %q: %v", uri, err)
	}

	doc := document.NewDocument(contents, tree)
	s.docs.Write(uri, doc)
	return nil
}

func (s *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	uri := params.TextDocument.URI
	contents := []byte(params.Text)
	tree, err := analysis.Parse(ctx, contents)
	if err != nil {
		return fmt.Errorf("could not parse file %q: %v", uri, err)
	}

	doc := document.NewDocument(contents, tree)
	s.docs.Write(uri, doc)
	return nil
}

func (s *Server) DidClose(_ context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	s.docs.Remove(params.TextDocument.URI)
	return nil
}