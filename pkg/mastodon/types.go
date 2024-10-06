package mastodon


type Secrets struct {
	clientKey string
	clientSecret string
	accessToken string
}

func(s *Secrets) SetClientKey(key string) { 
	s.clientKey = key
}

func(s *Secrets) SetClientSecret (secret string) {
		s.clientSecret = secret
}

func(s *Secrets) SetAccessToken (secret string) {
		s.accessToken = secret
}

func(s *Secrets) GetClientKey() string {
	return s.clientKey
}

func(s *Secrets) GetClientSecret() string {
	return s.clientSecret
}

func(s *Secrets) GetAccessToken() string {
	return s.clientSecret
}
