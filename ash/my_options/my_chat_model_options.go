package my_options

import (
	"time"

	"github.com/cloudwego/eino/components/model"
)

type MyChatModelOptions struct {
	Options    *model.Options
	Timeout    time.Duration
	RetryCount int
}

func WithRetryCount(count int) model.Option {
	return model.WrapImplSpecificOptFn(func(option *MyChatModelOptions) {
		option.RetryCount = count
	})
}

func WithTimeout(timeout time.Duration) model.Option {
	return model.WrapImplSpecificOptFn(func(option *MyChatModelOptions) {
		option.Timeout = timeout
	})
}
