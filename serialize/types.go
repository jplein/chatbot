package serialize

type Role string

const (
	System    Role = "system"
	User      Role = "user"
	Assistant Role = "assistant"
)

type Record struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type ChatPayload struct {
	Model    string   `json:"model"`
	Messages []Record `json:"messages"`
}

type ChatResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatResponseChoice struct {
	Message      Record `json:"message"`
	FinishReason string `json:"finish_reason"`
	Index        int    `json:"index"`
}

type ChatResponse struct {
	ID      string               `json:"id"`
	Object  string               `json:"object"`
	Created int                  `json:"created"`
	Model   string               `json:"model"`
	Usage   ChatResponseUsage    `json:"usage"`
	Choices []ChatResponseChoice `json:"choices"`
}
