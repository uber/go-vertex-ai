# Go Vertex AI

A Go Library for Google's Large Language Models on Vertex AI Platform

Google launched its latest Large Language Model(LLM) - [PaLM 2](https://ai.google/discover/palm2/), at Google I/O 2023. PaLM 2 powers Google's Bard chat tool, its competitor to OpenAI's ChatGPT. PaLM 2 is available to developers through Google's [Vertex AI Platform](https://cloud.google.com/vertex-ai). Their platform lets developers test, customize, and deploy instances of Google's PaLM2 LLMs.

Uber started using Google's PaLM2 LLM and realized that Google doesn't provide a Go library to access the LLM. After searching for a Go implementation, we found no good choices. So we built our own and are releasing it for everyone else to use.

We built this library to connect our Go Applications with Google's LLMs and supports only the features that we needed -
[Text Generation](https://cloud.google.com/vertex-ai/docs/generative-ai/text/test-text-prompts)
[Multi-Turn Chat Generation](https://cloud.google.com/vertex-ai/docs/generative-ai/chat/test-chat-prompts)
[Text Embedding](https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings)



## Usage

1. Download the library

```
go get github.com/uber/go-vertex-ai
```

2. Import the library into your project

```
import vertexai "github.com/uber/go-vertex-ai"
```

3. Call the library functions

```
  client, err := vertexai.New(key) // key of your project from Google Cloud Console

	resp, err := client.TextGeneration(context.Background(), vertexai.TextGenerationRequest{
		ProjectID:      examples.ProjectID,
		EndpointID:     "text-bison@001",
		Content:        "Hello, Google! What can you do?",
	})
```


## Methods we implemented

### [TextGeneration](https://console.cloud.google.com/vertex-ai/publishers/google/model-garden/text-bison)

[Documentation](https://cloud.google.com/vertex-ai/docs/generative-ai/model-reference/text) | [Code](./text_generation.go)

text-bison is the name of the PaLM 2 for text large language model that understands and generates language. It is fine-tuned to follow natural language instructions and is suitable for various language tasks, such as classification, extraction, summarization, and content generation.


### [ChatGeneration](https://console.cloud.google.com/vertex-ai/generative/language/create/chat)

[Documentation](https://cloud.google.com/vertex-ai/docs/generative-ai/model-reference/text-chat) | [Code](./chat_generation.go)

chat-bison is a large language model that excels at language understanding, generation, and conversations. This chat model is fine-tuned to conduct natural multi-turn conversations.
The PaLM 2 for chat is ideal for text tasks that require back-and-forth interactions. For text tasks that can be completed with one API response (without the need for continuous conversation), use the PaLM 2 for text.


### [TextEmbedding](https://console.cloud.google.com/vertex-ai/publishers/google/model-garden/textembedding-gecko)

[Documentation](https://cloud.google.com/vertex-ai/docs/generative-ai/model-reference/text-embeddings) | [Code](./text_embeddings.go)

Text embedding is an NLP technique that converts textual data into numerical vectors that can be processed by machine learning algorithms, especially large models. These vector representations are designed to capture the semantic meaning and context of the words they represent.


## Running Examples from Library

1. Download a Key for your Google Cloud Project and enable Vertex AI product for your project at [Google Cloud Console](https://console.cloud.google.com/)
2. Paste the absolute path of the key file (downloaded above) in (constants.go)[./examples/constants.go] file
3. Replace your Google Project ID in (constants.go)[./examples/constants.go] file



### TextGeneration

```
go run ./examples/text-generation/text_generation.go
```

### ChatGeneration

```
go run ./examples/chat-generation/chat_generation.go
```

### TextEmbedding

```
go run ./examples/text-embedding/text_embedding.go
```

## Contributing

To propose a new feature or report a bug, please open an issue with the tag #FeatureRequest or #Bug.

For Pull Requests:  please raise an issue before submitting a PR to discuss your proposed changes and check with the team. Then assign the issue to yourself and create a pull request. The team reviews PRs every three weeks, but holidays or work schedules may delay us occasionally.

## License

Vertex AI for GO is copyright 2023 by Uber Technologies, Inc, licensed under the Apache 2.0 license.  See the LICENSE and NOTICE files for more information.
