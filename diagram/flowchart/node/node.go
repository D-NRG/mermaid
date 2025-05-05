package node

import (
	"fmt"
	"github.com/D-NRG/mermaid/diagram/core"
	"sort"
	"strings"
)

type NodeFlowChart struct {
	ID    string
	Label string
	Type  string
	Style string
	Link  string
	Group string
}

const (
	TypeStart            = "start"
	TypeEnd              = "end"
	TypeProcess          = "process"
	TypeSubprocess       = "subprocess"
	TypeDatabase         = "database"
	TypeRhombus          = "rhombus"
	TypeCircle           = "circle"
	TypeHexagon          = "hexagon"
	TypeParallelogram    = "parallelogram"
	TypeParallelogramAlt = "parallelogram alt"
	TypeDoubleCircle     = "double circle"
	TypeTrapezoid        = "trapezoid"
	TypeTrapezoidAlt     = "trapezoid alt"

	//Styles
	StyleDotted            = "stroke-dasharray: 5 5"
	StyleFillTransparent   = "fill: transparent"
	StyleStrokeTransparent = "stroke: transparent"
)

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
		fmt.Fprintf(sb, "\tsubgraph %s\n", group)

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
	label = fmt.Sprintf("\"%s\"", label)

	var nodeShape string
	switch n.Type {
	case "start", "end":
		nodeShape = fmt.Sprintf("(%s)", label)
	case "process":
		nodeShape = fmt.Sprintf("[%s]", label)
	case "subprocess":
		nodeShape = fmt.Sprintf("[[%s]]", label)
	case "database":
		nodeShape = fmt.Sprintf("[(%s)]", label)
	case "rhombus":
		nodeShape = fmt.Sprintf("{%s}", label)
	case "circle":
		nodeShape = fmt.Sprintf("((%s))", label)
	case "hexagon":
		nodeShape = fmt.Sprintf("{{%s}}", label)
	case "parallelogram alt":
		nodeShape = fmt.Sprintf("[\\%s\\]", label)
	case "parallelogram":
		nodeShape = fmt.Sprintf("[/%s/]", label)
	case "double circle":
		nodeShape = fmt.Sprintf("(((%s)))", label)
	case "trapezoid alt":
		nodeShape = fmt.Sprintf("[\\%s/]", label)
	case "trapezoid":
		nodeShape = fmt.Sprintf("[/%s\\]", label)
	default:
		nodeShape = fmt.Sprintf("[%s]", label)
	}

	fmt.Fprintf(sb, "\n\t%s%s\n", n.ID, nodeShape)

	if n.Style != "" {
		fmt.Fprintf(sb, "\tstyle %s %s\n", n.ID, n.Style)
	}

	if n.Link != "" {
		fmt.Fprintf(sb, "\tclick %s \"%s\" _blank\n", n.ID, n.Link)
	}
}
