// Package flowchart provides fluent builders for constructing nodes and links
// in a flowchart diagram, with support for labels, styles, and grouping.
package flowchart

import (
	"github.com/D-NRG/mermaid/diagram/core"
)

type (
	FlowchartNodeBuilder struct {
		node *NodeFlowChart
	}
	FlowchartLinkBuilder struct {
		link *FlowchartLink
	}
)

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

func NewFlowchartLink() *FlowchartLinkBuilder {
	return &FlowchartLinkBuilder{
		link: &FlowchartLink{},
	}
}

func (b *FlowchartLinkBuilder) From(from ...string) *FlowchartLinkBuilder {
	b.link.From = from

	return b
}

func (b *FlowchartLinkBuilder) To(to ...string) *FlowchartLinkBuilder {
	b.link.To = to

	return b
}

func (b *FlowchartLinkBuilder) Labels(labels ...string) *FlowchartLinkBuilder {
	b.link.Labels = labels

	return b
}

func (b *FlowchartLinkBuilder) Styles(styles ...FlowchartStyle) *FlowchartLinkBuilder {
	b.link.Styles = styles

	return b
}

func (b *FlowchartLinkBuilder) DefaultStyle(style FlowchartStyle) *FlowchartLinkBuilder {
	b.link.DefaultStyle = style

	return b
}

func (b *FlowchartLinkBuilder) Build() core.LinkRenderer {
	return b.link
}
