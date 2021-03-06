package ast_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/graphism/dot"
	"github.com/graphism/dot/ast"
)

func TestParseFile(t *testing.T) {
	golden := []struct {
		in  string
		out string
	}{
		{in: "../internal/testdata/empty.dot"},
		{in: "../internal/testdata/graph.dot"},
		{in: "../internal/testdata/digraph.dot"},
		{in: "../internal/testdata/strict.dot"},
		{in: "../internal/testdata/multi.dot"},
		{in: "../internal/testdata/named_graph.dot"},
		{in: "../internal/testdata/node_stmt.dot"},
		{in: "../internal/testdata/edge_stmt.dot"},
		{in: "../internal/testdata/attr_stmt.dot"},
		{in: "../internal/testdata/attr.dot"},
		{
			in:  "../internal/testdata/subgraph.dot",
			out: "../internal/testdata/subgraph.golden",
		},
		{
			in:  "../internal/testdata/semi.dot",
			out: "../internal/testdata/semi.golden",
		},
		{
			in:  "../internal/testdata/empty_attr.dot",
			out: "../internal/testdata/empty_attr.golden",
		},
		{
			in:  "../internal/testdata/attr_lists.dot",
			out: "../internal/testdata/attr_lists.golden",
		},
		{
			in:  "../internal/testdata/attr_sep.dot",
			out: "../internal/testdata/attr_sep.golden",
		},
		{in: "../internal/testdata/subgraph_vertex.dot"},
		{
			in:  "../internal/testdata/port.dot",
			out: "../internal/testdata/port.golden",
		},
	}
	for _, g := range golden {
		file, err := dot.ParseFile(g.in)
		if err != nil {
			t.Errorf("%q: unable to parse file; %v", g.in, err)
			continue
		}
		// If no output path is specified, the input is already golden.
		out := g.in
		if len(g.out) > 0 {
			out = g.out
		}
		buf, err := ioutil.ReadFile(out)
		if err != nil {
			t.Errorf("%q: unable to read file; %v", g.in, err)
			continue
		}
		got := file.String()
		// Remove trailing newline.
		want := string(bytes.TrimSpace(buf))
		if got != want {
			t.Errorf("%q: graph mismatch; expected %q, got %q", g.in, want, got)
		}
	}
}

// Verify that all statements implement the Stmt interface.
var (
	_ ast.Stmt = &ast.NodeStmt{}
	_ ast.Stmt = &ast.EdgeStmt{}
	_ ast.Stmt = &ast.AttrStmt{}
	_ ast.Stmt = &ast.Attr{}
	_ ast.Stmt = &ast.Subgraph{}
)

// Verify that all vertices implement the Vertex interface.
var (
	_ ast.Vertex = &ast.Node{}
	_ ast.Vertex = &ast.Subgraph{}
)
