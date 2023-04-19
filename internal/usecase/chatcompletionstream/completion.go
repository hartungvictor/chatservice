package chatcompletionstream

import (
	"github.com/hartungvictor/chatservice/internal/domain/gateway"
	openai "github.com/sashabaranov/go-openai"
)

type ChatCompletionStream struct {
	chatGateway  gateway.ChatGateWay
	OpenAiClient *openai.Client
}
