package main

import (
	"fmt"
	"mermaid/diagram"
	"mermaid/diagram/core"
	"mermaid/diagram/flowchart/link"
	"mermaid/diagram/flowchart/node"
	links "mermaid/diagram/sequence/link"
	nodes "mermaid/diagram/sequence/node"
)

func main() {
	s := diagram.Diagram{
		Type: "sequence",
		Nodes: []core.NodeRenderer{
			nodes.NewSequenceNode("User").Label("User").Group("Frontend").Build(),
			nodes.NewSequenceNode("Server").Label("Backend Server").Group("Backend").Build(),
			nodes.NewSequenceNode("DB").Label("Database").Group("Backend").Build(),
			nodes.NewSequenceNode("Check").Build(),
		},
		Links: []core.LinkRenderer{
			links.NewSequenceLink().
				From("User", "Server", "DB", "Server").
				To("Server", "DB", "Server", "User").
				Labels("Send request", "Store data", "Data saved", "Response").
				Styles(links.DottedArrowHead, links.SolidArrowHead, links.DottedArrowHead, links.DottedOpenArrow).
				Activations("+", "+", "-", "-").
				DefaultStyle(links.SolidArrowHead).Build(),
			links.NewSequenceLink().
				From("User", "User").
				To("Server", "Server").
				Labels("Retry Request", "Cancel Request").
				Styles(links.SolidArrowHead, links.DottedCross).
				DefaultStyle(links.SolidArrowHead).
				Conditions("Server busy", "User cancels").Build(),
		},
	}
	f := diagram.Diagram{
		Type: "flowchart",
		Nodes: []core.NodeRenderer{
			node.NewFlowchartNode("SMS").Label("Smartmicro sensor").Type(node.TypeStart).Build(),
			node.NewFlowchartNode("tcp").Label("TCP client SMS DTP").Type(node.TypeStart).Build(),
			node.NewFlowchartNode("DB").Label("Clickhouse").Type(node.TypeDatabase).Build(),
			node.NewFlowchartNode("crud").Label("CRUD classes by vehicle length").Type(node.TypeStart).Build(),
			node.NewFlowchartNode("memory").Label("IN-MEMORY available sensors").Type(node.TypeStart).Build(),
			node.NewFlowchartNode("rest").Label("REST HTTP-API").Type(node.TypeStart).Build(),
			node.NewFlowchartNode("conf").Label("Dynamic configuration").Group("Config").Type(node.TypeStart).Build(),
			node.NewFlowchartNode("confini").Label("configs/default.yml configs/<env>.yml").Group("Config").Type(node.TypeProcess).Build(),
			node.NewFlowchartNode("confmerge").Label("Environment variables").Group("Config").Type(node.TypeProcess).Build(),
			node.NewFlowchartNode("amqp").Label("AMQP").Type(node.TypeProcess).Build(),
			node.NewFlowchartNode("licence").Label("License").Type(node.TypeStart).Style(node.StyleDotted).Build(),
			node.NewFlowchartNode("NATS").Label("NATS").Type(node.TypeProcess).Style(node.StyleDotted).Build(),
			node.NewFlowchartNode("confmergeNATS").Label("MQTT config topic").Group("Config").Type(node.TypeProcess).Style(node.StyleDotted).Build(),

			node.NewFlowchartNode("comment_supported").Label("📌 Supported sensors:UMRR-11 UMRR-12 (not tested)").Type(node.TypeProcess).Style("fill:transparent,stroke:transparent").Build(),
			node.NewFlowchartNode("comment_todo").Label("🛠️ TODO:UMRR-0A UMRR-0C").Type(node.TypeProcess).Style("fill:transparent,stroke:transparent").Build(),
			node.NewFlowchartNode("comment_insert").Label("💬 Inserts data in batches").Type(node.TypeProcess).Style("fill:transparent,stroke:transparent").Build(),
			node.NewFlowchartNode("comment_merge").Label("🛈 Used for config merge").Type(node.TypeProcess).Style("fill:transparent,stroke:transparent").Build(),
		},
		Links: []core.LinkRenderer{
			link.NewFlowchartLink().From("SMS").To("tcp").Labels("TCP pipe").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("tcp").To("DB", "amqp", "memory").Labels("batch insert", "Statistics / Per vehicle records").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("DB").To("crud").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("crud", "memory").To("rest").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("confini", "confmerge").To("conf").Labels("initial config", "config merge").DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().From("licence", "conf").To("tcp").Labels("sources", "connection licence").Styles(link.Dotted).DefaultStyle(link.StandardArrow).Build(),
			link.NewFlowchartLink().
				From("NATS", "licence", "NATS", "confmergeNATS").
				To("licence", "NATS", "confmergeNATS", "conf").
				Labels("licence update", "status update", "", "config merge").DefaultStyle(link.Dotted).Build(),

			link.NewFlowchartLink().From("comment_supported").To("SMS").DefaultStyle(link.Invisible).Build(),
			link.NewFlowchartLink().From("comment_todo").To("rest").DefaultStyle(link.Invisible).Build(),
			link.NewFlowchartLink().From("comment_insert").To("DB").DefaultStyle(link.Invisible).Build(),
			link.NewFlowchartLink().From("comment_merge").To("confmerge").DefaultStyle(link.Invisible).Build(),
		},
	}

	err := f.RenderUpdateREADME("README.md")
	//t, err := f.Render()
	//fmt.Println(t)
	fmt.Println(err)
	//out := s.RenderUpdateREADME("README.md")
	out, err1 := s.Render()
	fmt.Println(out)
	fmt.Println(err1)
}
