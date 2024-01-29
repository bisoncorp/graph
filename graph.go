package graph

type Node interface {
	Links() []Link
}

type Link interface {
	NodeIndex() int
	Weight() int
}

type Interface interface {
	Nodes() []Node
}
