package node

import (
	"github.com/D-NRG/mermaid/diagram/core"
)

type FlowchartNodeBuilder struct {
	node *NodeFlowChart
}

func NewFlowchartNode(id string) *FlowchartNodeBuilder {
	return &FlowchartNodeBuilder{
		node: &NodeFlowChart{ID: id},
	}
}

func (b *FlowchartNodeBuilder) Label(label string) *FlowchartNodeBuilder {
	b.node.Label = label
	return b
}

func (b *FlowchartNodeBuilder) Type(nodeType string) *FlowchartNodeBuilder {
	b.node.Type = nodeType
	return b
}

func (b *FlowchartNodeBuilder) Style(style string) *FlowchartNodeBuilder {
	b.node.Style = style
	return b
}

func (b *FlowchartNodeBuilder) Link(link string) *FlowchartNodeBuilder {
	b.node.Link = link
	return b
}

func (b *FlowchartNodeBuilder) Group(group string) *FlowchartNodeBuilder {
	b.node.Group = group
	return b
}

func (b *FlowchartNodeBuilder) Build() core.NodeRenderer {
	return b.node
}
