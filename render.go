// This is a modified version of modgraphviz created by the Go authors.
// Original Modgraphviz resides in the experimental repository.
// https://github.com/golang/exp/tree/master/cmd/modgraphviz

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"golang.org/x/mod/module"
)

// render translates “go mod graph” output taken from
// the 'in' reader into Graphviz's DOT language, writing
// to the 'out' writer.
func render(in io.Reader, out io.Writer, short, unPicked bool) error {
	graph, err := convert(in)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "digraph gomodgraph {\n")
	fmt.Fprintf(out, "\tpad=1;\n")
	fmt.Fprintf(out, "\trankdir=TB;\n")
	fmt.Fprintf(out, "\tranksep=\"1.2 equally\";\n")
	fmt.Fprintf(out, "\tsplines=ortho;\n")
	fmt.Fprintf(out, "\tnodesep=\"0.8\";\n")
	fmt.Fprintf(out, "\tnode [shape=plaintext style=\"filled,rounded\" penwidth=2 fontsize=12 fontname=\"monospace\"];\n")

	fmt.Fprintf(out, "\t%q [color=\"#007d9c\" shape=box height=1 style=\"rounded\" fontsize=16 label=<<b>%s</b>>];\n", graph.root, rootToHTML(graph.root, "#000000"))

	for _, n := range graph.mvsPicked {
		fmt.Fprintf(out, "\t%q [fillcolor=\"#007d9c\" label=<%s>];\n", n, textToHTML(n, "#ffffff", "#e3e3e3", short))
	}

	ignoreNodes := make(map[string]bool)
	if unPicked {
		for _, n := range graph.mvsUnpicked {
			fmt.Fprintf(out, "\t%q [fillcolor=\"#bababa\" label=<%s>];\n", n, textToHTML(n, "#0e0e0e", "#3f3f3f", short))
		}
	} else {
		for _, n := range graph.mvsUnpicked {
			ignoreNodes[n] = true
		}
		graph.mvsUnpicked = nil
	}

	out.Write(edgesAsDOT(graph, ignoreNodes))

	fmt.Fprintf(out, "}\n")

	return nil
}

// edgesAsDOT returns the edges in DOT notation.
func edgesAsDOT(gr *graph, ignoreNodes map[string]bool) []byte {
	var buf bytes.Buffer
	for _, e := range gr.edges {
		if ignoreNodes[e.from] || ignoreNodes[e.to] {
			continue
		}
		fmt.Fprintf(&buf, "\t%q -> %q", e.from, e.to)
		if _, ok := find(gr.mvsUnpicked, e.to); ok {
			fmt.Fprintf(&buf, "[style=dashed]")
		}
		fmt.Fprintf(&buf, ";\n")
	}
	return buf.Bytes()
}

func textToHTML(line string, modColor, verColor string, short bool) string {
	var mod, ver string
	if i := strings.IndexByte(line, '@'); i >= 0 {
		mod, ver = line[:i], line[i+1:]
	}

	u := fmt.Sprintf(`href="https://pkg.go.dev/%s?tab=doc"`, mod)

	var sb strings.Builder
	sb.WriteString(`<table border="0" cellspacing="8" `)
	if mod != "" {
		sb.WriteString(u)
	}
	sb.WriteString(`>`)

	if short {
		if strings.Count(mod, "/") >= 2 {
			i := strings.LastIndex(mod, "/")
			m1 := mod[i:]
			mod = mod[:i]
			mod = mod[strings.LastIndex(mod, "/")+1:] + m1
		}
		if len(ver) > 0 {
			if module.IsPseudoVersion(ver) {
				ver = ver[strings.LastIndex(ver, "-")+1:]
			}
		}
		sb.WriteString(`<tr><td><font color="`)
		sb.WriteString(modColor)
		sb.WriteString(`"><b>`)
		sb.WriteString(mod + `</b></font><font color="`)
		sb.WriteString(verColor)
		sb.WriteString(`">@`)
		sb.WriteString(ver)
		sb.WriteString("</font></td></tr>")
	} else {
		if len(mod) > 0 {
			sb.WriteString(`<tr><td><font color="`)
			sb.WriteString(modColor)
			sb.WriteString(`"><b>`)
			sb.WriteString(mod)
			sb.WriteString("</b></font></td></tr>")
		}

		if len(ver) > 0 {
			sb.WriteString(`<tr><td><font color="`)
			sb.WriteString(verColor)
			sb.WriteString(`" point-size="10">`)
			sb.WriteString(ver)
			sb.WriteString("</font></td></tr>")
		}
	}

	sb.WriteString("</table>")

	return sb.String()
}

func rootToHTML(root string, color string) string {
	var sb strings.Builder
	sb.WriteString(`<table border="0" cellspacing="8" >`)
	sb.WriteString(`<tr><td><font color="`)
	sb.WriteString(color)
	sb.WriteString(`"><b>`)
	sb.WriteString(root)
	sb.WriteString("</b></font></td></tr>")
	sb.WriteString("</table>")

	return sb.String()
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
