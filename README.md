# Go Vertex AI: Go Library for Google's Large Language Models on Vertex AI Platform

One can access Google's large generative AI models through their Vertex AI Platform but it doesn't have a Go SDK. This library provides an interface to connect to Google's Generative AI.

## Run the Samples

1. Download a Key for your Google Cloud Project and enable Vertex AI product for your project at [Google Cloud Console](https://console.cloud.google.com/)
2. Paste the absolute path of the file in (constants.go)[./examples/constants.go] file
3. Replace your Google Project ID in (constants.go)[./examples/constants.go] file
4. Run any of the following examples


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

If you'd like to propose a new feature or report a bug, please open an issue with the tag #FeatureRequest or #Bug before raising any PR to start a discussion about your proposed changes and check with the team. Then assign the issue to yourself and create a pull request. The team will review your pull request. The team will look into your requests in 3 weeks.
