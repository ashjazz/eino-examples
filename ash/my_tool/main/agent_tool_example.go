package main

import (
	"context"
	"github.com/cloudwego/eino-examples/ash/my_model_config"
	"github.com/cloudwego/eino-examples/ash/my_tool"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
	"log"
)

func main() {
	// 创建可用的工具列表
	todoTools := []tool.BaseTool{
		my_tool.AddTodoTool(),
		my_tool.UpdateTodoTool(),
		&my_tool.ListTodoTool{},
		my_tool.SearchTool,
	}

	ctx := context.Background()
	// 创建对话模型chatModel
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  my_model_config.MyApiKey,
		Model:   my_model_config.MyModelName,
		BaseURL: my_model_config.MyBaseURL,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// 获取每个tool工具的toolInfo
	toolsInfo := make([]*schema.ToolInfo, 0, len(todoTools))
	for _, eachT := range todoTools {
		tInfo, err := eachT.Info(ctx)
		if err != nil {
			log.Fatal(err)
		}
		toolsInfo = append(toolsInfo, tInfo)
	}

	// 将每个tool的toolInfo绑定到chatModel中
	err = chatModel.BindTools(toolsInfo)
	if err != nil {
		log.Fatal(err)
	}

	// TODO 创建 tools 节点
}
