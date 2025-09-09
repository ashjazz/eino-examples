package my_embedding

import (
	"context"
	"example.com/m/v2/model_config"
	"github.com/cloudwego/eino-ext/components/embedding/openai"
)

func NewEmbeddingModel(ctx context.Context) (*openai.Embedder, error) {
	var (
		defaultDim = 3072
	)

	encodingFmt := openai.EmbeddingEncodingFormatFloat
	embedder, err := openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
		APIKey:         model_config.ApiKey,
		Model:          model_config.EmbeddingModelName,
		BaseURL:        model_config.BaseURL,
		ByAzure:        false,
		Dimensions:     &defaultDim,
		Timeout:        0,
		EncodingFormat: &encodingFmt,
	})
	if err != nil {
		return nil, err
	}
	return embedder, nil
}
