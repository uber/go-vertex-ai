package vertexai

// SafetyAttributes is the safety attributes for a prediction
type SafetyAttributes struct {
	Blocked    bool      `json:"blocked"`
	Scores     []float64 `json:"scores"`
	Categories []string  `json:"categories"`
}

type inputInstances struct {
	Content  string        `json:"content"`
	Context  string        `json:"context"`
	Examples []Example     `json:"examples"`
	Messages []ChatMessage `json:"messages"`
}

// Example represents a example passed for few shot learning in multi turn chat conversations
type Example struct {
	Input  ChatMessage `json:"input"`
	Output ChatMessage `json:"output"`
}

// ChatMessage is a chat prompt
type ChatMessage struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

type parameters struct {
	Temperature    *float64 `json:"temperature,omitempty"`
	MaxDecodeSteps *int     `json:"maxDecodeSteps,omitempty"`
	TopP           *float64 `json:"topP,omitempty"`
	TopK           *int     `json:"topK,omitempty"`
}

// payload is the payload for the Vertex AI Prediction API.
type payload struct {
	Instances  []inputInstances `json:"instances"`
	Parameters parameters       `json:"parameters"`
}
