package my_chat_model

import (
	"context"
	"example.com/m/v2/model_config"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

func NewChatModel(ctx context.Context) (model.BaseChatModel, error) {
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  model_config.ApiKey,
		Model:   model_config.ModelName,
		BaseURL: model_config.BaseURL,
	})
	if err != nil {
		return nil, err
	}
	return chatModel, nil
}
