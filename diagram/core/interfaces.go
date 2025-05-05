package core

import "strings"

type (
	NodeRenderer interface {
	}
	LinkRenderer interface {
		Render(builder *strings.Builder)
	}
)
