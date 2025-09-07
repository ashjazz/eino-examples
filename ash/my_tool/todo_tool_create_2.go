package my_tool

import (
	"context"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"log"
)

type TodoUpdateParams struct {
	Id        string  `json:"id" jsonschema:"description=id of the todo"`
	Content   *string `json:"content,omitempty" jsonschema:"description=content of the todo"`
	StartedAt *int64  `json:"started_at,omitempty" jsonschema:"description=start time in unix timestamp"`
	Deadline  *int64  `json:"deadline,omitempty" jsonschema:"description=deadline of the todo in unix timestamp"`
	Done      *bool   `json:"done,omitempty" jsonschema:"description=done status"`
}

func UpdateTodoFunc(_ context.Context, params *TodoUpdateParams) (string, error) {
	// Mock 处理逻辑
	return `{"msg": "update todo success"}`, nil
}

// UpdateTodoTool
// 相比第一种方法，通过tag注解为每个参数添加描述信息，框架自动生成tool的params信息
// 只需要维护一份代码，提升代码的可维护性
func UpdateTodoTool() tool.InvokableTool {
	// 构造InvokableTool
	t, err := utils.InferTool(
		"update_todo",
		"Update a todo item, eg: content,deadline...",
		UpdateTodoFunc,
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return t
}
