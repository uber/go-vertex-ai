package vertexai

// standard models supported by TextGeneration endpoint
const (
	TextBison001   string = "text-bison@001"
	TextBison      string = "text-bison"
	TextUnicorn001 string = "text-unicorn@001"
	CodeBison001   string = "code-bison@001" // Code Generation
	CodeGecko001   string = "code-gecko@001" // Code Completion
)

// standard models supported by ChatGeneration endpoint
const (
	ChatBison001     string = "chat-bison@001"
	CodeChatBison001 string = "codechat-bison@001" // Code Generation
)

// standard models supported by TextEmbedding endpoint
const (
	TextEmbeddingGecko001 string = "textembedding-gecko@001"
)
