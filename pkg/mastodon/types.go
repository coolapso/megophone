package mastodon


type Secrets struct {
	apiKey string
	apiKeySecret string
}

func(s *Secrets) SetApiKey(key string) { 
	s.apiKey = key
}

func(s *Secrets) SetApiKeySecret (secret string) {
		s.apiKeySecret = secret
}

func(s *Secrets) GetApiKey() string {
	return s.apiKey
}

func(s *Secrets) GetApiKeySecret() string {
	return s.apiKeySecret
}
