// Package sequence defines constants for sequence diagram styles,
// representing different types of arrows and message flows.
package sequence

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
