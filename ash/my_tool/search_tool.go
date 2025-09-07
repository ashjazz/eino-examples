package my_tool

import (
	"context"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo/v2"
	"github.com/cloudwego/eino/components/tool"
)

var SearchTool tool.InvokableTool

func init() {
	s, err := duckduckgo.NewTextSearchTool(context.Background(), &duckduckgo.Config{})
	if err != nil {
		panic(err)
	}
	SearchTool = s
}
