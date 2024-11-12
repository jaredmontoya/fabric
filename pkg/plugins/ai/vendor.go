package ai

import (
	"context"

	goopenai "github.com/sashabaranov/go-openai"

	"github.com/danielmiessler/fabric/pkg/common"
	"github.com/danielmiessler/fabric/pkg/plugins"
)

type Vendor interface {
	plugins.Plugin
	ListModels() ([]string, error)
	SendStream([]*goopenai.ChatCompletionMessage, *common.ChatOptions, chan string) error
	Send(context.Context, []*goopenai.ChatCompletionMessage, *common.ChatOptions) (string, error)
}
