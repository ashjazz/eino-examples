package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-examples/ash/my_model_config"
	"github.com/cloudwego/eino-examples/ash/my_tool"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
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

	// 创建 tools 节点
	todoToolsNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: todoTools,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// 构建完整的处理链chain
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("todo_tools_node"))

	// 编译并运行agent
	agent, err := chain.Compile(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 运行示例
	response, err := agent.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "添加一个学习 Eino 的 TODO，同时搜索一下 cloudwego/eino 的仓库地址",
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, resp := range response {
		fmt.Println(resp)
	}
}
