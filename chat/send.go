package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/jplein/chatbot/storage"
	"github.com/jplein/chatbot/transcript"
)

const ChatEndpoint = "https://api.openai.com/v1/chat/completions"

// const ChatEndpoint = "http://127.0.0.1:8080/"
const DefaultChatModel = "gpt-3.5-turbo"

type ChatPayload struct {
	Model    string              `json:"model"`
	Messages []transcript.Record `json:"messages"`
}

type ChatResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatResponseChoice struct {
	Message      transcript.Record `json:"message"`
	FinishReason string            `json:"finish_reason"`
	Index        int               `json:"index"`
}

type ChatResponse struct {
	ID      string               `json:"id"`
	Object  string               `json:"object"`
	Created int                  `json:"created"`
	Model   string               `json:"model"`
	Usage   ChatResponseUsage    `json:"usage"`
	Choices []ChatResponseChoice `json:"choices"`
}

func Send(dir *storage.Dir, apiKey string, msg string) (string, error) {
	var err error
	var parentPID int = os.Getppid()

	question := transcript.Record{Role: transcript.User, Content: msg}
	if err = transcript.Write(dir, parentPID, question); err != nil {
		return "", err
	}

	var req http.Request
	req.Header = make(http.Header)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	req.Method = http.MethodPost

	var uri *url.URL
	if uri, err = url.Parse(ChatEndpoint); err != nil {
		return "", err
	}

	req.URL = uri

	payload := ChatPayload{
		Model: DefaultChatModel,
	}
	payload.Messages = append(
		payload.Messages,
		transcript.Record{
			Role:    transcript.User,
			Content: msg,
		},
	)

	var payloadBuf []byte
	if payloadBuf, err = json.Marshal(payload); err != nil {
		return "", err
	}

	req.Body = io.NopCloser(strings.NewReader(string(payloadBuf)))

	client := &http.Client{}
	var res *http.Response
	if res, err = client.Do(&req); err != nil {
		return "", err
	}

	defer res.Body.Close()

	var body []byte
	if body, err = io.ReadAll(res.Body); err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("unexpected status code %d: %s", res.StatusCode, body)
	}

	var chatResponse ChatResponse
	if err = json.Unmarshal(body, &chatResponse); err != nil {
		return "", err
	}

	if chatResponse.Choices == nil || len(chatResponse.Choices) == 0 {
		return "", fmt.Errorf("no choices in response from chat API")
	}

	responseChoice := chatResponse.Choices[0]

	if err = transcript.Write(dir, parentPID, responseChoice.Message); err != nil {
		return "", err
	}

	return responseChoice.Message.Content, nil
}
