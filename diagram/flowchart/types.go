// Package flowchart defines constants for link styles and node types
// used in Mermaid-compatible flowchart diagrams.
package flowchart

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

	// Styles.
	StyleDotted            = "stroke-dasharray: 5 5"
	StyleFillTransparent   = "fill: transparent"
	StyleStrokeTransparent = "stroke: transparent"
)
