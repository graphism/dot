package parser_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/graphism/dot"
)

func TestParseFile(t *testing.T) {
	golden := []struct {
		in  string
		out string
	}{
		{in: "testdata/empty.dot"},
		{in: "testdata/graph.dot"},
		{in: "testdata/digraph.dot"},
		{in: "testdata/strict.dot"},
		{in: "testdata/multi.dot"},
		{in: "testdata/named_graph.dot"},
		{in: "testdata/node_stmt.dot"},
		{in: "testdata/edge_stmt.dot"},
		{in: "testdata/attr_stmt.dot"},
		{in: "testdata/attr.dot"},
		{
			in:  "testdata/subgraph.dot",
			out: "testdata/subgraph.golden",
		},
		{
			in:  "testdata/semi.dot",
			out: "testdata/semi.golden",
		},
		{
			in:  "testdata/empty_attr.dot",
			out: "testdata/empty_attr.golden",
		},
		{
			in:  "testdata/attr_lists.dot",
			out: "testdata/attr_lists.golden",
		},
		{
			in:  "testdata/attr_sep.dot",
			out: "testdata/attr_sep.golden",
		},
		{in: "testdata/subgraph_vertex.dot"},
		{
			in:  "testdata/port.dot",
			out: "testdata/port.golden",
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