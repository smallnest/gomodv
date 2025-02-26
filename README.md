# gomodv

[![Go Report Card](https://goreportcard.com/badge/github.com/smallnest/gomodv)](https://goreportcard.com/report/github.com/smallnest/gomodv) &nbsp;&nbsp;&nbsp; [![Go Coverage](https://gocover.io/_badge/github.com/smallnest/gomodv?nocache=modgv)](https://gocover.io/_badge/github.com/smallnest/gomodv?nocache=modgv) &nbsp;&nbsp;&nbsp; [![Go API Reference](https://img.shields.io/badge/go-docs-blue.svg?style=flat)](https://pkg.go.dev/github.com/smallnest/gomodv?tab=doc)

This is a modified version of [modgraphviz](https://github.com/golang/exp/tree/master/cmd/modgraphviz) and forked from [lucasepe/modgv](https://github.com/lucasepe/modgv).

Converts 'go mod graph' output into [GraphViz](https://graphviz.gitlab.io/download/)'s DOT language.

- takes no options or arguments
- it reads the output generated by “go mod graph” on stdin
- generates a DOT language and writes to stdout

## Usage:

```bash
go mod graph | gomodv | dot -Tpng -o graph.png
```

For each module:
- the node representing the greatest version (i.e., the version chosen by Go's MVS algorithm) is colored blue.
- other nodes, which aren't in the final build list, are colored grey

## Installation

```bash
go get github.com/smallnest/gomodv
```

Here 👉 https://graphviz.gitlab.io/download/ how to install [GraphViz](https://graphviz.gitlab.io/download/) for your OS.

## Sample output (PNG)

```bash
go mod graph | gomodv | dot -Tpng -o graph2.png
```

![](./graph2.png)

In short mode and not render unpicked:

```bash
go mod graph | gomodv -s -unpicked=false| dot -Tpng -o graph.png
```

![](./graph.png)
---

## Sample output (PDF with clickable links to module docs)

```bash
go mod graph | gomodv | dot -Tps2 -o graph.ps
ps2pdf graph.ps graph.pdf
```

![View generated PDF](./graph.pdf)

