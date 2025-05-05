// Package sequence provides builder utilities for constructing nodes and links
// in a sequence diagram. It supports fluent-style configuration of diagram elements.
package sequence

import (
	"github.com/D-NRG/mermaid/diagram/core"
)

type (
	SequenceNodeBuilder struct {
		node *NodeSequence
	}
	SequenceLinkBuilder struct {
		link *SequenceLink
	}
)

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

func NewSequenceLink() *SequenceLinkBuilder {
	return &SequenceLinkBuilder{
		link: &SequenceLink{},
	}
}

func (b *SequenceLinkBuilder) From(from ...string) *SequenceLinkBuilder {
	b.link.From = from

	return b
}

func (b *SequenceLinkBuilder) To(to ...string) *SequenceLinkBuilder {
	b.link.To = to

	return b
}

func (b *SequenceLinkBuilder) Labels(labels ...string) *SequenceLinkBuilder {
	b.link.Labels = labels

	return b
}

func (b *SequenceLinkBuilder) Styles(styles ...SequenceStyle) *SequenceLinkBuilder {
	b.link.Styles = styles

	return b
}

func (b *SequenceLinkBuilder) Activations(activations ...string) *SequenceLinkBuilder {
	b.link.Activations = activations

	return b
}

func (b *SequenceLinkBuilder) Conditions(conditions ...string) *SequenceLinkBuilder {
	b.link.Conditions = conditions

	return b
}

func (b *SequenceLinkBuilder) DefaultStyle(style SequenceStyle) *SequenceLinkBuilder {
	b.link.DefaultStyle = style

	return b
}

func (b *SequenceLinkBuilder) Build() core.LinkRenderer {
	return b.link
}
