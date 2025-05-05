// Package diagram defines the main entry point for rendering Mermaid diagrams.
// It supports flowchart and sequence diagram types and includes helpers for
// integrating diagram output into markdown files like README.md.
package diagram

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/D-NRG/mermaid/diagram/core"
	"github.com/D-NRG/mermaid/diagram/flowchart"
	"github.com/D-NRG/mermaid/diagram/sequence"
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
	fileMode      = 0o600
)

func (d *Diagram) Render() (string, error) {
	var sb strings.Builder

	if err := d.typeDiagram(&sb); err != nil {
		return "", err
	}

	return sb.String(), nil
}

func (d *Diagram) RenderUpdateREADME(readmePath string) error {
	//nolint:err113 // single use
	errInvalidReadMe := errors.New("invalid file: only README.md is allowed")
	//nolint:err113 // single use
	errMissingMermaidMarkers := errors.New("README does not contain required markers:" +
		" <!-- BEGIN_MERMAID --> ... <!-- END_MERMAID -->")

	absPath, err := filepath.Abs(readmePath)
	if err != nil {
		return fmt.Errorf("resolve path: %w", err)
	}

	if filepath.Base(absPath) != "README.md" {
		return fmt.Errorf("%w, got %s", errInvalidReadMe, filepath.Base(absPath))
	}

	var sb strings.Builder
	if errType := d.typeDiagram(&sb); errType != nil {
		return errType
	}

	//nolint:gosec //readmePath is validated to point only to a README.md file,
	// avoiding path traversal or arbitrary file access
	content, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	re := regexp.MustCompile(`(?s)(<!-- BEGIN_MERMAID -->).*?(<!-- END_MERMAID -->)`)
	if !re.Match(content) {
		return errMissingMermaidMarkers
	}

	newContent := re.ReplaceAllString(string(content), "$1\n"+sb.String()+"$2")

	if err = os.WriteFile(absPath, []byte(newContent), fileMode); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func (d *Diagram) typeDiagram(sb *strings.Builder) error {
	//nolint:err113 //single use
	errInvalidDiagramType := errors.New("invalid diagram type")

	switch d.Type {
	case TypeSequence:
		d.renderSequenceDiagram(sb)
	case TypeFlowChart:
		d.renderFlowchart(sb)
	default:
		return fmt.Errorf("%w: %s", errInvalidDiagramType, d.Type)
	}

	return nil
}

func (d *Diagram) renderFlowchart(sb *strings.Builder) {
	sb.WriteString("```mermaid\nflowchart TD\n\n")

	flowchart.RenderFlowchartNodes(d.Nodes, sb)

	for _, link := range d.Links {
		link.Render(sb)
	}

	sb.WriteString("```\n")
}

func (d *Diagram) renderSequenceDiagram(sb *strings.Builder) {
	sb.WriteString("```mermaid\nsequenceDiagram\n\n")

	sequence.RenderSequenceNodes(d.Nodes, sb)

	for _, link := range d.Links {
		link.Render(sb)
	}

	sb.WriteString("```\n")
}
