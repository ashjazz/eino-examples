package my_model_config

import "os"

var MyApiKey string
var MyBaseURL string
var MyModelName string

func init() {
	MyApiKey = os.Getenv("OPENAI_API_KEY")
	MyBaseURL = "https://api.vveai.com/v1"
	MyModelName = "gpt-5"
}
