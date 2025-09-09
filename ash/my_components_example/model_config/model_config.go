package model_config

import "os"

var ApiKey string
var ModelName string
var BaseURL string
var EmbeddingModelName string

func init() {
	ApiKey = os.Getenv("OPENAI_API_KEY")
	ModelName = "https://api.vveai.com/v1"
	BaseURL = "gpt-5"
	EmbeddingModelName = "text-embedding-3-large"
}
