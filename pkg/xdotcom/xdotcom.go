package xdotcom

type Secrets struct {
	oauthToken       string
	oauthTokenSecret string
	apiKey           string
	apiKeySecret     string
}

func (x *Secrets) SetOauthToken(s string) {
	x.oauthToken = s
}

func (x *Secrets) SetOauthTokenSecret(s string) {
	x.oauthTokenSecret = s
}

func (x *Secrets) SetApiKey(s string) {
	x.apiKey = s
}

func (x *Secrets) SetApiKeySecret(s string) {
	x.apiKeySecret = s
}

func (x *Secrets) GetOauthToken() string {
	return x.oauthToken
}

func (x *Secrets) GetOauthTokenSecret() string {
	return x.oauthTokenSecret
}

func (x *Secrets) GetApiKey() string {
	return x.apiKey
}

func (x *Secrets) GetApiKeySecret() string {
	return x.apiKeySecret
}
