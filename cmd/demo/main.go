package main

import (
	"fmt"

	"github.com/D-NRG/mermaid/diagram"
	"github.com/D-NRG/mermaid/diagram/core"
	"github.com/D-NRG/mermaid/diagram/flowchart"
	"github.com/D-NRG/mermaid/diagram/sequence"
)

func main() {
	s := diagram.Diagram{
		Type: "sequence",
		Nodes: []core.NodeRenderer{
			sequence.NewSequenceNode("User").Label("User").Group("Frontend").Build(),
			sequence.NewSequenceNode("WebUI").Label("Web Interface").Group("Frontend").Build(),
			sequence.NewSequenceNode("API").Label("API Server").Group("Backend").Build(),
			sequence.NewSequenceNode("DB").Label("PostgreSQL").Group("Backend").Build(),
			sequence.NewSequenceNode("Queue").Label("Order Queue").Group("Infra").Build(),
			sequence.NewSequenceNode("Worker").Label("Worker").Group("Infra").Build(),
		},
		Links: []core.LinkRenderer{
			sequence.NewSequenceLink().
				From("User", "WebUI", "API", "DB", "API").
				To("WebUI", "API", "DB", "API", "Queue").
				Labels("Click 'Create Order'", "POST /orders", "INSERT INTO orders", "OK", "Publish event").
				Styles(sequence.SolidOpenArrow, sequence.SolidArrowHead, sequence.SolidArrowHead, sequence.DottedArrowHead, sequence.SolidOpenArrow).
				Activations("+", "+", "+", "-").
				Build(),

			// Queue → Worker
			sequence.NewSequenceLink().
				From("Queue").To("Worker").
				Labels("New order").
				Styles(sequence.DottedOpenArrow).
				Build(),

			// Worker → Worker (самообработка)
			sequence.NewSequenceLink().
				From("Worker").To("Worker").
				Labels("validate + enrich").
				Styles(sequence.SolidBidirectional).
				Build(),

			// Worker → DB
			sequence.NewSequenceLink().
				From("Worker").To("DB").
				Labels("UPDATE order").
				Styles(sequence.DottedArrowHead).
				Build(),

			// API → WebUI
			sequence.NewSequenceLink().
				From("API").To("WebUI").
				Labels("return 200 OK").
				Styles(sequence.SolidArrowHead).
				Activations("-").
				Build(),

			// UI → User
			sequence.NewSequenceLink().
				From("WebUI").To("User").
				Labels("Order Created!").
				Styles(sequence.SolidOpenArrow).
				Activations("-").
				Build(),

			sequence.NewSequenceLink().
				From("User", "WebUI").
				To("WebUI", "API").
				Labels("Retry click", "Cancel").
				Conditions("1st request failed", "User clicked Cancel").
				Styles(sequence.SolidArrowHead, sequence.DottedCross).
				Build(),
		},
	}

	f := diagram.Diagram{
		Type: "flowchart",
		Nodes: []core.NodeRenderer{
			// Ingest
			flowchart.NewFlowchartNode("sensor_A").Label("Sensor A").Group("Ingest").Type(flowchart.TypeStart).Build(),
			flowchart.NewFlowchartNode("sensor_B").Label("Sensor B").Group("Ingest").Type(flowchart.TypeStart).Build(),
			flowchart.NewFlowchartNode("gateway").Label("Gateway").Group("Ingest").Type(flowchart.TypeSubprocess).Build(),

			// Pipeline
			flowchart.NewFlowchartNode("parser").Label("Parser").Group("Pipeline").Type(flowchart.TypeProcess).Build(),
			flowchart.NewFlowchartNode("validator").Label("Validator").Group("Pipeline").Type(flowchart.TypeRhombus).Build(),
			flowchart.NewFlowchartNode("normalizer").Label("Normalizer").Group("Pipeline").Type(flowchart.TypeHexagon).Build(),

			// Storage
			flowchart.NewFlowchartNode("db").Label("TimescaleDB").Group("Storage").Type(flowchart.TypeDatabase).Build(),
			flowchart.NewFlowchartNode("analytics").Label("Clickhouse").Group("Storage").Type(flowchart.TypeDatabase).Style(flowchart.StyleDotted).Build(), // background batch

			// API
			flowchart.NewFlowchartNode("cache").Label("Redis").Group("API").Type(flowchart.TypeParallelogram).Build(),
			flowchart.NewFlowchartNode("rest").Label("REST API").Group("API").Type(flowchart.TypeParallelogramAlt).Build(),
			flowchart.NewFlowchartNode("graphql").Label("GraphQL").Group("API").Type(flowchart.TypeParallelogramAlt).Style(flowchart.StyleDotted).Build(),

			// Auth & Events
			flowchart.NewFlowchartNode("licence").Label("License Check").Group("Auth").Type(flowchart.TypeCircle).Style(flowchart.StyleDotted).Build(),
			flowchart.NewFlowchartNode("nats").Label("NATS Streaming").Group("Events").Type(flowchart.TypeDoubleCircle).Build(),

			// Monitoring
			flowchart.NewFlowchartNode("logger").Label("Logger").Group("Monitoring").Type(flowchart.TypeProcess).Style("fill:#f5f5f5,stroke:#aaa").Build(),
			flowchart.NewFlowchartNode("alert").Label("Alert System").Group("Monitoring").Type(flowchart.TypeTrapezoid).Build(),

			// Meta: комментарии
			flowchart.NewFlowchartNode("comment_api").Label("🌐 Exposed to clients").Group("Meta").Type(flowchart.TypeProcess).
				Style(flowchart.StyleFillTransparent + "," + flowchart.StyleStrokeTransparent).Build(),
		},
		Links: []core.LinkRenderer{
			// Ingest
			flowchart.NewFlowchartLink().From("sensor_A", "sensor_B").To("gateway").Labels("rawA", "rawB").DefaultStyle(flowchart.StandardArrow).Build(),
			// Pipeline
			flowchart.NewFlowchartLink().From("gateway").To("parser").Labels("payload").DefaultStyle(flowchart.StandardArrow).Build(),
			flowchart.NewFlowchartLink().From("parser").To("validator").Labels("parsed").DefaultStyle(flowchart.StandardArrow).Build(),
			flowchart.NewFlowchartLink().From("validator").To("normalizer").Labels("valid").DefaultStyle(flowchart.StandardArrow).Build(),
			// Storage
			flowchart.NewFlowchartLink().From("normalizer").To("db").Labels("clean").DefaultStyle(flowchart.StandardArrow).Build(),
			flowchart.NewFlowchartLink().From("db").To("analytics").Labels("batch ETL").DefaultStyle(flowchart.Dotted).Build(),
			// API
			flowchart.NewFlowchartLink().From("db").To("cache").Labels("hot data").DefaultStyle(flowchart.StandardArrow).Build(),
			flowchart.NewFlowchartLink().From("cache").To("rest", "graphql").Labels("serve", "query").DefaultStyle(flowchart.StandardArrow).Build(),
			// Auth & Events
			flowchart.NewFlowchartLink().From("licence").To("gateway").Labels("authz").DefaultStyle(flowchart.Dotted).Build(),
			flowchart.NewFlowchartLink().From("nats").To("normalizer").Labels("config updates").DefaultStyle(flowchart.DottedTwice).Build(),
			// Monitoring
			flowchart.NewFlowchartLink().From("parser", "rest", "db").To("logger").Labels("log", "access", "metrics").DefaultStyle(flowchart.StandardArrow).Build(),
			flowchart.NewFlowchartLink().From("logger").To("alert").Labels("if anomaly").DefaultStyle(flowchart.BoldArrow).Build(),
			// Comments
			flowchart.NewFlowchartLink().From("comment_api").To("rest").DefaultStyle(flowchart.Invisible).Build(),
		},
	}

	err := f.RenderUpdateREADME("README.md")
	// t, err := f.Render()
	// fmt.Println(t)
	fmt.Println(err)
	// out := s.RenderUpdateREADME("README.md")
	out, err1 := s.Render()
	fmt.Println(err1)
	fmt.Println(out)
}
