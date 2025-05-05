package link

import (
	"github.com/D-NRG/mermaid/diagram/core"
)

type FlowchartLinkBuilder struct {
	link *FlowchartLink
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
