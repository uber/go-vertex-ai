package vertexai

type clientConfig struct {
	authToken string
}

type clientConfigOption func(*clientConfig) *clientConfig

func newClientConfig(opts ...clientConfigOption) clientConfig {
	config := &clientConfig{}

	for _, opt := range opts {
		opt(config)
	}

	return *config
}

func setAuthToken(token string) clientConfigOption {
	return func(cc *clientConfig) *clientConfig {
		cc.authToken = token
		return cc
	}
}
