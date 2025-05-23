// Package sequence defines renderable nodes for sequence diagrams,
// supporting participant labeling and optional grouping into boxes.
package sequence

import (
	"fmt"
	"sort"
	"strings"

	"github.com/D-NRG/mermaid/diagram/core"
)

type (
	NodeSequence struct {
		ID    string
		Label string
		Group string
	}
)

func RenderSequenceNodes(nodes []core.NodeRenderer, sb *strings.Builder) {
	grouped := make(map[string][]*NodeSequence)

	var ungrouped []*NodeSequence

	for _, node := range nodes {
		if ns, ok := node.(*NodeSequence); ok {
			if ns.Group == "" {
				ungrouped = append(ungrouped, ns)
			} else {
				grouped[ns.Group] = append(grouped[ns.Group], ns)
			}
		}
	}

	for _, node := range ungrouped {
		node.render(sb)
	}

	groupNames := make([]string, 0, len(grouped))
	for name := range grouped {
		groupNames = append(groupNames, name)
	}

	sort.Strings(groupNames)

	for _, group := range groupNames {
		_, _ = fmt.Fprintf(sb, "box %s\n", group)

		for _, node := range grouped[group] {
			node.render(sb)
		}

		sb.WriteString("end\n\n")
	}
}

func (n *NodeSequence) render(sb *strings.Builder) {
	label := n.Label
	if label == "" {
		label = n.ID
	}

	if label != n.ID {
		_, _ = fmt.Fprintf(sb, "\tparticipant %s as %s\n", n.ID, label)
	} else {
		_, _ = fmt.Fprintf(sb, "\tparticipant %s\n", n.ID)
	}
}
