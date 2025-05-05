package diagram

import (
	"fmt"
	"github.com/D-NRG/mermaid/diagram/core"
	nodef "github.com/D-NRG/mermaid/diagram/flowchart/node"
	nodes "github.com/D-NRG/mermaid/diagram/sequence/node"
	"os"
	"regexp"
	"strings"
)

type (
	Diagram struct {
		Type  string
		Nodes []core.NodeRenderer
		Links []core.LinkRenderer
	}
)

const (
	TypeFlowChart = "flowchart"
	TypeSequence  = "sequence"
)

func (d *Diagram) Render() (string, error) {
	var sb strings.Builder

	if err := d.typeDiagram(&sb); err != nil {
		return "", err
	}

	return sb.String(), nil
}

func (d *Diagram) RenderUpdateREADME(readmePath string) error {
	var sb strings.Builder

	if err := d.typeDiagram(&sb); err != nil {
		return err
	}

	content, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	re := regexp.MustCompile(`(?s)(<!-- BEGIN_MERMAID -->).*?(<!-- END_MERMAID -->)`)
	if !re.Match(content) {
		return fmt.Errorf("README does not contain required markers: <!-- BEGIN_MERMAID --> ... <!-- END_MERMAID -->")
	}

	newContent := re.ReplaceAllString(string(content), "$1\n"+sb.String()+"$2")

	if err = os.WriteFile(readmePath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
func (d *Diagram) typeDiagram(sb *strings.Builder) error {
	switch d.Type {
	case "sequence":
		d.renderSequenceDiagram(sb)
	case "flowchart":
		d.renderFlowchart(sb)
	default:
		return fmt.Errorf("invalid diagram type: %s", d.Type)
	}
	return nil
}

func (d *Diagram) renderFlowchart(sb *strings.Builder) {
	sb.WriteString("```mermaid\nflowchart TD\n\n")

	nodef.RenderFlowchartNodes(d.Nodes, sb)

	for _, link := range d.Links {
		link.Render(sb)
	}

	sb.WriteString("```\n")
}

func (d *Diagram) renderSequenceDiagram(sb *strings.Builder) {
	sb.WriteString("```mermaid\nsequenceDiagram\n\n")

	nodes.RenderSequenceNodes(d.Nodes, sb)

	for _, link := range d.Links {
		link.Render(sb)
	}

	sb.WriteString("```\n")
}
