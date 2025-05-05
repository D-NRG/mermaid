// Package sequence defines types and logic for rendering sequence diagram links.
// It includes customizable styles, labels, and conditional flows between entities.
package sequence

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

func (l *SequenceLink) Render(sb *strings.Builder) {
	if len(l.From) == 0 || len(l.To) == 0 {
		return
	}

	condIndex := 0
	conditionActive := false

	for i := 0; i < len(l.From) && i < len(l.To); i++ {
		if condIndex < len(l.Conditions) {
			prefix := "alt"
			if conditionActive {
				prefix = "else"
			}

			_, _ = fmt.Fprintf(sb, "\t%s %s\n", prefix, l.Conditions[condIndex])

			condIndex++
			conditionActive = true
		}

		l.renderStep(sb, i, l.From[i], l.To[i])
	}

	if conditionActive {
		sb.WriteString("\tend\n")
	}
}

func (l *SequenceLink) renderStep(sb *strings.Builder, i int, from, to string) {
	style := l.DefaultStyle
	if i < len(l.Styles) && l.Styles[i] != "" {
		style = l.Styles[i]
	}

	activation := ""

	if i < len(l.Activations) {
		switch l.Activations[i] {
		case "+", "-":
			activation = l.Activations[i]
		}
	}

	label := ""
	if i < len(l.Labels) && l.Labels[i] != "" {
		label = l.Labels[i]
	}

	_, _ = fmt.Fprintf(sb, "\t%s %s%s %s", from, style, activation, to)

	if label != "" {
		_, _ = fmt.Fprintf(sb, " : %s", label)
	}

	sb.WriteString("\n")
}
