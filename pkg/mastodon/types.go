package mastodon

type Secrets struct {
	server       string
	clientID     string
	clientSecret string
	accessToken  string
}

func (s *Secrets) SetServer(server string) {
	s.server = server
}

func (s *Secrets) SetClientID(key string) {
	s.clientID = key
}

func (s *Secrets) SetClientSecret(secret string) {
	s.clientSecret = secret
}

func (s *Secrets) SetAccessToken(secret string) {
	s.accessToken = secret
}

func (s *Secrets) GetServer() string {
	return s.server
}

func (s *Secrets) GetClientID() string {
	return s.clientID
}

func (s *Secrets) GetClientSecret() string {
	return s.clientSecret
}

func (s *Secrets) GetAccessToken() string {
	return s.accessToken
}
