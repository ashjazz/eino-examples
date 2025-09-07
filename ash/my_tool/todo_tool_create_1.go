package my_tool

import (
	"context"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

func AddTodoFunc(_ context.Context, params map[string]any) (string, error) {
	// Mock 处理逻辑
	return `{"msg": "add todo success"}`, nil
}

// AddTodoTool
// 这种方式虽然直观，但存在一个明显的缺点：需要在 ToolInfo 中手动定义参数信息（ParamsOneOf），
// 和实际的参数结构（TodoAddParams）是分开定义的。这样不仅造成了代码的冗余，
// 而且在参数发生变化时需要同时修改两处地方，容易导致不一致，维护起来也比较麻烦。
func AddTodoTool() tool.InvokableTool {
	// 工具信息
	toolInfo := &schema.ToolInfo{
		Name: "add_todo",
		Desc: "Add a todo item",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"content": {
				Desc:     "The content of the todo item",
				Type:     schema.String,
				Required: true,
			},
			"start_at": {
				Desc: "The started time of the todo item, in unix timestamp",
				Type: schema.Integer,
			},
			"deadline": {
				Desc: "The deadline of the todo item, in unix timestamp",
				Type: schema.Integer,
			},
		}),
	}
	return utils.NewTool(toolInfo, AddTodoFunc)
}
