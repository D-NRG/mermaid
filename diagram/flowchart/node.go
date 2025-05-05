// Package flowchart provides node definitions and rendering logic
// for flowchart diagrams using Mermaid-style syntax.
package flowchart

import (
	"fmt"
	"sort"
	"strings"

	"github.com/D-NRG/mermaid/diagram/core"
)

type NodeFlowChart struct {
	ID    string
	Label string
	Type  string
	Style string
	Link  string
	Group string
}

func RenderFlowchartNodes(nodes []core.NodeRenderer, sb *strings.Builder) {
	grouped := make(map[string][]*NodeFlowChart)

	var ungrouped []*NodeFlowChart

	for _, node := range nodes {
		if nf, ok := node.(*NodeFlowChart); ok {
			if nf.Group == "" {
				ungrouped = append(ungrouped, nf)
			} else {
				grouped[nf.Group] = append(grouped[nf.Group], nf)
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
		_, _ = fmt.Fprintf(sb, "\tsubgraph %s\n", group)

		for _, node := range grouped[group] {
			node.render(sb)
		}

		sb.WriteString("\tend\n\n")
	}
}

func (n *NodeFlowChart) render(sb *strings.Builder) {
	label := n.Label
	if label == "" {
		label = n.ID
	}
	label = fmt.Sprintf("%q", label)

	nodeShape := getNodeShape(n.Type, label)

	_, _ = fmt.Fprintf(sb, "\n\t%s%s\n", n.ID, nodeShape)

	if n.Style != "" {
		_, _ = fmt.Fprintf(sb, "\tstyle %s %s\n", n.ID, n.Style)
	}

	if n.Link != "" {
		_, _ = fmt.Fprintf(sb, "\tclick %s \"%s\" _blank\n", n.ID, n.Link)
	}
}

func getNodeShape(nodeType, label string) string {
	switch nodeType {
	case "start", "end":
		return fmt.Sprintf("(%s)", label)
	case "process":
		return fmt.Sprintf("[%s]", label)
	case "subprocess":
		return fmt.Sprintf("[[%s]]", label)
	case "database":
		return fmt.Sprintf("[(%s)]", label)
	case "rhombus":
		return fmt.Sprintf("{%s}", label)
	case "circle":
		return fmt.Sprintf("((%s))", label)
	case "hexagon":
		return fmt.Sprintf("{{%s}}", label)
	case "parallelogram alt":
		return fmt.Sprintf("[\\%s\\]", label)
	case "parallelogram":
		return fmt.Sprintf("[/%s/]", label)
	case "double circle":
		return fmt.Sprintf("(((%s)))", label)
	case "trapezoid alt":
		return fmt.Sprintf("[\\%s/]", label)
	case "trapezoid":
		return fmt.Sprintf("[/%s\\]", label)
	default:
		return fmt.Sprintf("[%s]", label)
	}
}
