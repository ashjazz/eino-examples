package my_tool

import (
	"context"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// ListTodoTool 可以自定义Tool的实现，自定义Invoke和Info方法
// Info中描述了Tool具体的参数，而Invoke负责实现具体逻辑，两个部分的参数达到解耦效果
// 但是两部分的参数约束性变弱，只有运行时才能发现具体的bug，增加了测试工作量
type ListTodoTool struct{}

func (l *ListTodoTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "list_todo",
		Desc: "List all todo items",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"finished": {
				Desc:     "filter todo items if finished",
				Type:     schema.Boolean,
				Required: true,
			},
		}),
	}, nil
}

func (l *ListTodoTool) InvokableRun(ctx context.Context, argumentsInJson string, opts ...tool.Option) (string, error) {
	// Mock调用逻辑
	return `{"todos": [{"id": "1", "content": "在2024年12月10日之前完成Eino项目演示文稿的准备工作", "started_at": 1717401600, "deadline": 1717488000, "done": false}]}`, nil
}
