package link

import (
	"fmt"
	"strings"
)

type (
	SequenceLink struct {
		From         []string
		To           []string
		Labels       []string
		Styles       []SequenceStyle
		Activations  []string // "", "+", "-"
		DefaultStyle SequenceStyle
		Conditions   []string
	}
	SequenceStyle string
)

const (
	SolidNoArrow        SequenceStyle = "->"     // Solid line without arrow
	DottedNoArrow       SequenceStyle = "-->"    // Dotted line without arrow
	SolidArrowHead      SequenceStyle = "->>"    // Solid line with arrowhead
	DottedArrowHead     SequenceStyle = "-->>"   // Dotted line with arrowhead
	SolidBidirectional  SequenceStyle = "<<->>"  // Solid line with bidirectional arrowheads (v11.0.0+)
	DottedBidirectional SequenceStyle = "<<-->>" // Dotted line with bidirectional arrowheads (v11.0.0+)
	SolidCross          SequenceStyle = "-x"     // Solid line with a cross at the end
	DottedCross         SequenceStyle = "--x"    // Dotted line with a cross at the end
	SolidOpenArrow      SequenceStyle = "-)"     // Solid line with an open arrow at the end (async)
	DottedOpenArrow     SequenceStyle = "--)"    // Dotted line with a open arrow at the end (async)
)

func (l *SequenceLink) Render(sb *strings.Builder) {
	if len(l.From) == 0 || len(l.To) == 0 {
		return
	}

	condIndex := 0
	conditionActive := false

	for i := 0; i < len(l.From) && i < len(l.To); i++ {
		if condIndex < len(l.Conditions) {
			if !conditionActive {
				fmt.Fprintf(sb, "\talt %s\n", l.Conditions[condIndex])
				conditionActive = true
			} else {
				fmt.Fprintf(sb, "\telse %s\n", l.Conditions[condIndex])
			}
			condIndex++
		}

		from := l.From[i]
		to := l.To[i]

		style := l.DefaultStyle
		if i < len(l.Styles) && l.Styles[i] != "" {
			style = l.Styles[i]
		}

		activation := ""
		if i < len(l.Activations) {
			switch l.Activations[i] {
			case "+":
				activation = "+"
			case "-":
				activation = "-"
			}
		}

		label := ""
		if i < len(l.Labels) && l.Labels[i] != "" {
			label = l.Labels[i]
		}

		fmt.Fprintf(sb, "\t%s %s%s %s", from, style, activation, to)

		if label != "" {
			fmt.Fprintf(sb, " : %s", label)
		}
		sb.WriteString("\n")
	}

	if conditionActive {
		sb.WriteString("\tend\n")
	}
}
