# Go Vertex AI:
#  Go Library for Google's Large Language Models on Vertex AI Platform

Google's large generative AI models can be accessed through their Vertex AI Platform that let’s you test, customize, and deploy instances of Google’s PaLM2 Large Language Models (LLMs) so that you can leverage the capabilities of PaLM 2 in your applications, but it doesn't have an official Go SDK to invoke the LLMs. This library provides a Go interface to connect your Go Applications to Google's Generative AI LLMs. This library supports Text Completion, Multi-Turn Chat, and text embeddings generation.

## Usage

1. Download the library

```
go get github.com/uber/go-vertex-ai
```

2. Import the library in your project

```
import vertexai "github.com/uber/go-vertex-ai"
```

3. Call the library functions

```
  client, err := vertexai.New(key) // key of your project from Google Cloud Console

	resp, err := client.TextGeneration(context.Background(), vertexai.TextGenerationRequest{
		ProjectID:      examples.ProjectID,
		EndpointID:     "text-bison@001",
		Content:        "Hello Google! What can you do?",
	})
```

## Running Examples

### Setup

1. Clone the repo
2. Download a Key for your Google Cloud Project and enable Vertex AI product for your project at [Google Cloud Console](https://console.cloud.google.com/)
3. Paste the absolute path of the Key file in (constants.go)[./examples/constants.go] file
4. Replace your Google Project ID in (constants.go)[./examples/constants.go] file
5. Run the Samples below from the library


### TextGeneration
  [Google Documentation](https://cloud.google.com/vertex-ai/docs/generative-ai/text/test-text-prompts)

```
go run ./examples/text-generation/text_generation.go
```

### ChatGeneration
  [Google Documentation](https://cloud.google.com/vertex-ai/docs/generative-ai/chat/test-chat-prompts)

```
go run ./examples/chat-generation/chat_generation.go
```

### TextEmbedding
  [Google Documentation](https://cloud.google.com/vertex-ai/docs/generative-ai/embeddings/get-text-embeddings)

```
go run ./examples/text-embedding/text_embedding.go
```

## Contributing

To propose a new feature or report a bug, please open an issue with the tag #FeatureRequest or #Bug.

For Pull Requests:  please raise an issue before raising a PR, to start a discussion about your proposed changes and check with the team. Then assign the issue to yourself and create a pull request. The team will review your PR in about 3 weeks.

## License

Vertex AI for GO is copyright 2023 by Uber Technologies, Inc, and licensed under the Apache 2.0 license.  See the LICENSE and NOTICE files for more information.
