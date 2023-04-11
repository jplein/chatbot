package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/jplein/chatbot/serialize"
	"github.com/jplein/chatbot/storage"
	"github.com/jplein/chatbot/transcript"
)

const ChatEndpoint = "https://api.openai.com/v1/chat/completions"

// const ChatEndpoint = "http://127.0.0.1:8080/"
const DefaultChatModel = "gpt-3.5-turbo"

func Send(dir *storage.Dir, apiKey string, msg string) (string, int, error) {
	var err error
	var parentPID int = os.Getppid()

	question := serialize.Record{Role: serialize.User, Content: msg}
	if err = transcript.Write(dir, parentPID, question); err != nil {
		return "", 0, err
	}

	var req http.Request
	req.Header = make(http.Header)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	req.Method = http.MethodPost

	var uri *url.URL
	if uri, err = url.Parse(ChatEndpoint); err != nil {
		return "", 0, err
	}

	var context []serialize.Record
	if context, err = transcript.Read(dir, parentPID); err != nil {
		return "", 0, err
	}

	req.URL = uri

	payload := serialize.ChatPayload{
		Model: DefaultChatModel,
	}

	for _, record := range context {
		payload.Messages = append(payload.Messages, record)
	}

	payload.Messages = append(
		payload.Messages,
		serialize.Record{
			Role:    serialize.User,
			Content: msg,
		},
	)

	var payloadBuf []byte
	if payloadBuf, err = json.Marshal(payload); err != nil {
		return "", 0, err
	}

	req.Body = io.NopCloser(strings.NewReader(string(payloadBuf)))

	client := &http.Client{}
	var res *http.Response
	if res, err = client.Do(&req); err != nil {
		return "", 0, err
	}

	defer res.Body.Close()

	var body []byte
	if body, err = io.ReadAll(res.Body); err != nil {
		return "", 0, err
	}

	if res.StatusCode != 200 {
		return "", 0, fmt.Errorf("unexpected status code %d: %s", res.StatusCode, body)
	}

	var chatResponse serialize.ChatResponse
	if err = json.Unmarshal(body, &chatResponse); err != nil {
		return "", 0, err
	}

	if chatResponse.Choices == nil || len(chatResponse.Choices) == 0 {
		return "", 0, fmt.Errorf("no choices in response from chat API")
	}

	responseChoice := chatResponse.Choices[0]

	if err = transcript.Write(dir, parentPID, responseChoice.Message); err != nil {
		return "", 0, err
	}

	return responseChoice.Message.Content, chatResponse.Usage.TotalTokens, nil
}
