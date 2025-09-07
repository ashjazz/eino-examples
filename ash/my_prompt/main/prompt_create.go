package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-examples/ash/my_model_config"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"log"
)

var chatModel model.BaseChatModel

func init() {
	c, err := openai.NewChatModel(context.Background(), &openai.ChatModelConfig{
		APIKey:  my_model_config.MyApiKey,
		BaseURL: my_model_config.MyBaseURL,
		Model:   my_model_config.MyModelName,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	chatModel = c
}

func PromptWithFString(ctx context.Context) {

	t := prompt.FromMessages(schema.FString,
		schema.SystemMessage("我的名字是{name}"),
		schema.MessagesPlaceholder("history", true),
		schema.UserMessage("{question}"),
	)

	messages, err := t.Format(ctx, map[string]any{
		"name": "Ash",
		"history": []*schema.Message{
			schema.UserMessage("我的爱好是打电脑游戏"),
		},
		"question": "你知道我的名字吗？我的爱好是什么？",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	responseMsg, err := chatModel.Generate(ctx, messages)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(responseMsg.Content)
}

func PromptWithJinja2(ctx context.Context) {
	t := prompt.FromMessages(schema.Jinja2,
		schema.SystemMessage("请用{% if language == 'english' %}english{% else %}中文{% endif %}解释什么是{{topic}}"),
	)

	messages, err := t.Format(ctx, map[string]any{
		"language": "中文",
		"topic":    "苹果",
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	response, err := chatModel.Generate(ctx, messages)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(response.Content)
	return
}

func PromptWithGoTemplate(ctx context.Context) {
	t := prompt.FromMessages(schema.GoTemplate,
		schema.SystemMessage("请使用{{if .english}}english{{else}}中文{{end}}来解释{{.topic}}"),
	)

	messages, err := t.Format(ctx, map[string]any{
		"english": false,
		"topic":   "苹果手机",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	response, err := chatModel.Generate(ctx, messages)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(response.Content)
}

func PromptWithMessage(ctx context.Context) {
	messages := []*schema.Message{
		schema.SystemMessage("你叫做Bob，我叫做Ash"),
		{
			Role:    schema.User,
			Content: "请问你知道我叫什么嘛？",
		},
	}

	response, err := chatModel.Generate(ctx, messages)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(response.Content)
}

func main() {
	// 1、用 FString方式构造prompt模板
	fmt.Println("============ 1、FString ==============")
	PromptWithFString(context.Background())
	fmt.Println("/============ 1、FString ==============/")
	fmt.Println("")

	// 2、用 Jinja2方式构造prompt模板
	fmt.Println("============ 2、Jinja2 ==============")
	PromptWithJinja2(context.Background())
	fmt.Println("/============ 2、Jinja2 ==============/")
	fmt.Println("")

	// 3、用 GoTemplate方式构造prompt模板
	fmt.Println("============ 3、GoTemplate ==============")
	PromptWithGoTemplate(context.Background())
	fmt.Println("/============ 3、GoTemplate ==============/")
	fmt.Println("")

	// 4、直接构造prompt模板
	fmt.Println("============ 4、Messages ==============")
	PromptWithMessage(context.Background())
	fmt.Println("/============ 4、Messages ==============/")
	fmt.Println("")
}
