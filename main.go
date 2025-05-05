package main

import (
	"fmt"
	"github.com/D-NRG/mermaid/diagram"
	"github.com/D-NRG/mermaid/diagram/core"
	"github.com/D-NRG/mermaid/diagram/flowchart/link"
	"github.com/D-NRG/mermaid/diagram/flowchart/node"
	links "github.com/D-NRG/mermaid/diagram/sequence/link"
	nodes "github.com/D-NRG/mermaid/diagram/sequence/node"
)

func main() {
	s := diagram.Diagram{
		Type: "sequence",
		Nodes: []core.NodeRenderer{
			nodes.NewSequenceNode("User").Label("User").Group("Frontend").Build(),
			nodes.NewSequenceNode("WebUI").Label("Web Interface").Group("Frontend").Build(),
			nodes.NewSequenceNode("API").Label("API Server").Group("Backend").Build(),
			nodes.NewSequenceNode("DB").Label("PostgreSQL").Group("Backend").Build(),
			nodes.NewSequenceNode("Queue").Label("Order Queue").Group("Infra").Build(),
			nodes.NewSequenceNode("Worker").Label("Worker").Group("Infra").Build(),
		},
		Links: []core.LinkRenderer{
			links.NewSequenceLink().
				From("User", "WebUI", "API", "DB", "API").
				To("WebUI", "API", "DB", "API", "Queue").
				Labels("Click 'Create Order'", "POST /orders", "INSERT INTO orders", "OK", "Publish event").
				Styles(links.SolidOpenArrow, links.SolidArrowHead, links.SolidArrowHead, links.DottedArrowHead, links.SolidOpenArrow).
				Activations("+", "+", "+", "-").
				Build(),

			// Queue → Worker
			links.NewSequenceLink().
				From("Queue").To("Worker").
				Labels("New order").
				Styles(links.DottedOpenArrow).
				Build(),

			// Worker → Worker (самообработка)
			links.NewSequenceLink().
				From("Worker").To("Worker").
				Labels("validate + enrich").
				Styles(links.SolidBidirectional).
				Build(),

			// Worker → DB
			links.NewSequenceLink().
				From("Worker").To("DB").
				Labels("UPDATE order").
				Styles(links.DottedArrowHead).
				Build(),

			// API → WebUI
			links.NewSequenceLink().
				From("API").To("WebUI").
				Labels("return 200 OK").
				Styles(links.SolidArrowHead).
				Activations("-").
				Build(),

			// UI → User
			links.NewSequenceLink().
				From("WebUI").To("User").
				Labels("Order Created!").
				Styles(links.SolidOpenArrow).
				Activations("-").
				Build(),

			links.NewSequenceLink().
				From("User", "WebUI").
				To("WebUI", "API").
				Labels("Retry click", "Cancel").
				Conditions("1st request failed", "User clicked Cancel").
				Styles(links.SolidArrowHead, links.DottedCross).
				Build(),
		},
	}

	f := diagram.Diagram{
		Type: "flowchart",
		Nodes: []core.NodeRenderer{
			// Ingest
			node.NewFlowchartNode("sensor_A").Label("Sensor A").Group("Ingest").Type(node.TypeStart).Build(),
			node.NewFlowchartNode("sensor_B").Label("Sensor B").Group("Ingest").Type(node.TypeStart).Build(),
			node.NewFlowchartNode("gateway").Label("Gateway").Group("Ingest").Type(node.TypeSubprocess).Build(),

			// Pipeline
			node.NewFlowchartNode("parser").Label("Parser").Group("Pipeline").Type(node.TypeProcess).Build(),
			node.NewFlowchartNode("validator").Label("Validator").Group("Pipeline").Type(node.TypeRhombus).Build(),
			node.NewFlowchartNode("normalizer").Label("Normalizer").Group("Pipeline").Type(node.TypeHexagon).Build(),

			// Storage
			node.NewFlowchartNode("db").Label("TimescaleDB").Group("Storage").Type(node.TypeDatabase).Build(),
			node.NewFlowchartNode("analytics").Label("Clickhouse").Group("Storage").Type(node.TypeDatabase).Style(node.StyleDotted).Build(), // background batch

			// API
			node.NewFlowchartNode("cache").Label("Redis").Group("API").Type(node.TypeParallelogram).Build(),
			node.NewFlowchartNode("rest").Label("REST API").Group("API").Type(node.TypeParallelogramAlt).Build(),
			node.NewFlowchartNode("graphql").Label("GraphQL").Group("API").Type(node.TypeParallelogramAlt).Style(node.StyleDotted).Build(),

			// Auth & Events
			node.NewFlowchartNode("licence").Label("License Check").Group("Auth").Type(node.TypeCircle).Style(node.StyleDotted).Build(),
			node.NewFlowchartNode("nats").Label("NATS Streaming").Group("Events").Type(node.TypeDoubleCircle).Build(),

			// Monitoring
			node.NewFlowchartNode("logger").Label("Logger").Group("Monitoring").Type(node.TypeProcess).Style("fill:#f5f5f5,stroke:#aaa").Build(),
			node.NewFlowchartNode("alert").Label("Alert System").Group("Monitoring").Type(node.TypeTrapezoid).Build(),

			// Meta: комментарии
			node.NewFlowchartNode("comment_api").Label("🌐 Exposed to clients").Group("Meta").Type(node.TypeProcess).
				Style(node.StyleFillTransparent + "," + node.StyleStrokeTransparent).Build(),
		},
		Links: []core.LinkRenderer{
			// Ingest
			link.NewFlowchartLink().From("sensor_A", "sensor_B").To("gateway").Labels("rawA", "rawB").DefaultStyle(link.StandardArrow).Build(),
			// Pipeline
			link.NewFlowchartLink().From("gateway").To("parser").Labels("payload").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("parser").To("validator").Labels("parsed").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("validator").To("normalizer").Labels("valid").DefaultStyle(link.StandardArrow).Build(),
			// Storage
			link.NewFlowchartLink().From("normalizer").To("db").Labels("clean").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("db").To("analytics").Labels("batch ETL").DefaultStyle(link.Dotted).Build(),
			// API
			link.NewFlowchartLink().From("db").To("cache").Labels("hot data").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("cache").To("rest", "graphql").Labels("serve", "query").DefaultStyle(link.StandardArrow).Build(),
			// Auth & Events
			link.NewFlowchartLink().From("licence").To("gateway").Labels("authz").DefaultStyle(link.Dotted).Build(),
			link.NewFlowchartLink().From("nats").To("normalizer").Labels("config updates").DefaultStyle(link.DottedTwice).Build(),
			// Monitoring
			link.NewFlowchartLink().From("parser", "rest", "db").To("logger").Labels("log", "access", "metrics").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("logger").To("alert").Labels("if anomaly").DefaultStyle(link.BoldArrow).Build(),
			// Comments
			link.NewFlowchartLink().From("comment_api").To("rest").DefaultStyle(link.Invisible).Build(),
		},
	}

	err := f.RenderUpdateREADME("README.md")
	//t, err := f.Render()
	//fmt.Println(t)
	fmt.Println(err)
	//out := s.RenderUpdateREADME("README.md")
	out, err1 := s.Render()
	fmt.Println(err1)
	fmt.Println(out)
}
