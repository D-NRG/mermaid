package link

import (
	"fmt"
	"strings"
)

type (
	FlowchartLink struct {
		From         []string
		To           []string
		Labels       []string
		Styles       []FlowchartStyle
		DefaultStyle FlowchartStyle
	}
	FlowchartStyle string
)

const (
	StandardArrow      FlowchartStyle = "-->"
	StandardTwiceArrow FlowchartStyle = "<-->"
	WithoutArrow       FlowchartStyle = "---"
	XArrow             FlowchartStyle = "--x"
	XTwiceArrow        FlowchartStyle = "x--x"
	OArrow             FlowchartStyle = "--o"
	OTwiceArrow        FlowchartStyle = "o--o"
	Dotted             FlowchartStyle = "-.->"
	DottedTwice        FlowchartStyle = "<-.->"
	BoldArrow          FlowchartStyle = "==>"
	BoldArrowTwice     FlowchartStyle = "<==>"
	Invisible          FlowchartStyle = "~~~"
)

func (l *FlowchartLink) Render(sb *strings.Builder) {
	if len(l.From) == 0 || len(l.To) == 0 {
		return
	}

	switch {
	case l.isChain():
		l.renderChainLink(sb)
	case len(l.From) > 1 && len(l.To) > 1:
		l.renderGroupLink(sb)
	case len(l.From) == 1 && len(l.To) > 1:
		l.renderOneToManyLink(sb)
	case len(l.From) > 1 && len(l.To) == 1:
		l.renderManyToOneLink(sb)
	case len(l.From) == 1 && len(l.To) == 1:
		l.renderSingleLink(sb)
	}
}

func (l *FlowchartLink) renderGroupLink(sb *strings.Builder) {
	fmt.Fprintf(sb, "\t%s %s %s\n",
		strings.Join(l.From, " & "),
		l.DefaultStyle,
		strings.Join(l.To, " & "))
}

func (l *FlowchartLink) renderChainLink(sb *strings.Builder) {
	var chain []string
	chain = append(chain, l.From[0])
	for i := 0; i < len(l.From); i++ {
		style := l.DefaultStyle
		if len(l.Styles) > i && l.Styles[i] != "" {
			style = l.Styles[i]
		}
		label := ""
		if len(l.Labels) > i && l.Labels[i] != "" {
			label = fmt.Sprintf(`|"%s"|`, l.Labels[i])
		}
		chain = append(chain, fmt.Sprintf("%s%s%s", style, label, l.To[i]))
	}
	fmt.Fprintf(sb, "\t%s\n", strings.Join(chain, ""))
}

func (l *FlowchartLink) renderOneToManyLink(sb *strings.Builder) {
	l.renderGroupedLink(sb, []string{l.From[0]}, l.To, false)
}

func (l *FlowchartLink) renderManyToOneLink(sb *strings.Builder) {
	l.renderGroupedLink(sb, l.From, []string{l.To[0]}, true)
}

func (l *FlowchartLink) renderSingleLink(sb *strings.Builder) {
	label := ""
	if len(l.Labels) > 0 && l.Labels[0] != "" {
		label = fmt.Sprintf(`|"%s"|`, l.Labels[0])
	}
	fmt.Fprintf(sb, "\t%s %s%s %s\n",
		l.From[0],
		l.DefaultStyle,
		label,
		l.To[0])
}

func (l *FlowchartLink) isChain() bool {
	if len(l.From) != len(l.To) || len(l.From) < 2 {
		return false
	}
	for i := 1; i < len(l.From); i++ {
		if l.From[i] != l.To[i-1] {
			return false
		}
	}
	return true
}

func (l *FlowchartLink) renderGroupedLink(
	sb *strings.Builder,
	heads []string,
	tails []string,
	isManyToOne bool,
) {
	switch {
	case len(l.Labels) > 1:
		var count int
		if isManyToOne {
			count = len(heads)
		} else {
			count = len(tails)
		}

		for i := 0; i < count; i++ {
			head := heads[0]
			tail := tails[0]
			if isManyToOne {
				head = heads[i]
			} else if i < len(tails) {
				tail = tails[i]
			}

			style := l.DefaultStyle
			if i < len(l.Styles) && l.Styles[i] != "" {
				style = l.Styles[i]
			}

			label := ""
			if i < len(l.Labels) && l.Labels[i] != "" {
				label = fmt.Sprintf(`|"%s"|`, l.Labels[i])
			}
			fmt.Fprintf(sb, "\t%s %s%s %s\n", head, style, label, tail)
		}

	case len(l.Labels) == 1:
		label := fmt.Sprintf(`|"%s"|`, l.Labels[0])
		fmt.Fprintf(sb, "\t%s %s%s %s\n",
			strings.Join(heads, " & "),
			l.DefaultStyle,
			label,
			strings.Join(tails, " & "))

	default:
		fmt.Fprintf(sb, "\t%s %s %s\n",
			strings.Join(heads, " & "),
			l.DefaultStyle,
			strings.Join(tails, " & "))
	}
}
