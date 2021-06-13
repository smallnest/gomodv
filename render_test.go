package main

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	out := &bytes.Buffer{}
	in := bytes.NewBuffer([]byte(`
test.com/A@v1.0.0 test.com/B@v1.2.3
test.com/B@v1.0.0 test.com/C@v4.5.6
`))
	if err := render(in, out); err != nil {
		t.Fatal(err)
	}

	gotGraph := string(out.Bytes())
	wantGraph := `digraph gomodgraph {
	pad=1;
	rankdir=TB;
	ranksep="1.2 equally";
	splines=ortho;
	nodesep="0.8";
	node [shape=plaintext style="filled,rounded" penwidth=2 fontsize=12 fontname="monospace"];
	"" [shape=underline style="" fontsize=14 label=<<b></b>>];
	"test.com/A@v1.0.0" [fillcolor="#0c5525" label=<<table border="0" cellspacing="8" href="https://pkg.go.dev/test.com/A?tab=doc"><tr><td><font color="#fafafa"><b>test.com/A</b></font></td></tr><tr><td><font color="#fafafa" point-size="10">v1.0.0</font></td></tr></table>>];
	"test.com/B@v1.2.3" [fillcolor="#0c5525" label=<<table border="0" cellspacing="8" href="https://pkg.go.dev/test.com/B?tab=doc"><tr><td><font color="#fafafa"><b>test.com/B</b></font></td></tr><tr><td><font color="#fafafa" point-size="10">v1.2.3</font></td></tr></table>>];
	"test.com/C@v4.5.6" [fillcolor="#0c5525" label=<<table border="0" cellspacing="8" href="https://pkg.go.dev/test.com/C?tab=doc"><tr><td><font color="#fafafa"><b>test.com/C</b></font></td></tr><tr><td><font color="#fafafa" point-size="10">v4.5.6</font></td></tr></table>>];
	"test.com/B@v1.0.0" [fillcolor="#a3a3a3" label=<<table border="0" cellspacing="8" href="https://pkg.go.dev/test.com/B?tab=doc"><tr><td><font color="#0e0e0e"><b>test.com/B</b></font></td></tr><tr><td><font color="#0e0e0e" point-size="10">v1.0.0</font></td></tr></table>>];
	"test.com/A@v1.0.0" -> "test.com/B@v1.2.3";
	"test.com/B@v1.0.0" -> "test.com/C@v4.5.6";
}
`
	if gotGraph != wantGraph {
		t.Fatalf("\ngot: %s\nwant: %s", wantGraph, gotGraph)
	}
}

func Testfind(t *testing.T) {
	tests := []struct {
		set  []string
		el   string
		want bool
	}{
		{
			[]string{"filled", "rounded", "striped"},
			"filled",
			true,
		},
		{
			[]string{"filled", "rounded", "striped"},
			"slashed",
			false,
		},
		{
			[]string{"mela", "banana", "caffè"},
			"caffè",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.el, func(t *testing.T) {
			if _, got := find(tt.set, tt.el); got != tt.want {
				t.Errorf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}
