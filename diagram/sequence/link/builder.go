package link

import (
	"mermaid/diagram/core"
)

type SequenceLinkBuilder struct {
	link *SequenceLink
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
