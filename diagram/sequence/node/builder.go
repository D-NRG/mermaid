package node

import (
	"mermaid/diagram/core"
)

type SequenceNodeBuilder struct {
	node *NodeSequence
}

func NewSequenceNode(id string) *SequenceNodeBuilder {
	return &SequenceNodeBuilder{
		node: &NodeSequence{ID: id},
	}
}

func (b *SequenceNodeBuilder) Label(label string) *SequenceNodeBuilder {
	b.node.Label = label
	return b
}

func (b *SequenceNodeBuilder) Group(group string) *SequenceNodeBuilder {
	b.node.Group = group
	return b
}

func (b *SequenceNodeBuilder) Build() core.NodeRenderer {
	return b.node
}
