package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/coolapso/megophone/internal/util"
	"net/url"
	"os"
	"strings"

	gomasto "github.com/mattn/go-mastodon"
	"github.com/spf13/viper"
)

func configMastodonServer(reader *bufio.Reader, c *config) {
	if server, isSet := os.LookupEnv("MEGOPHONE_MASTODON_SERVER"); isSet {
		c.m.SetServer(server)
	}

	fmt.Printf("Mastodon Server (%v): ", c.m.GetServer())
	GetServerInput, _ := reader.ReadString('\n')
	if cleanInput := strings.TrimSpace(GetServerInput); cleanInput != "" {
		c.m.SetServer(cleanInput)
	}
	viper.Set("mastodon_server", c.m.GetServer())
}

func registerMastodonApp(ctx context.Context, c *config) (*gomasto.Application, error) {
	appConfig := &gomasto.AppConfig{
		Server:       c.m.GetServer(),
		ClientName:   "megophone",
		Scopes:       "read write follow",
		Website:      "https://github.com/coolapso/megophone",
		RedirectURIs: redirectUri,
	}

	return gomasto.RegisterApp(ctx, appConfig)
}

func getMastodonUserAuthorizationCode(reader *bufio.Reader, app *gomasto.Application) (string, error) {
	u, err := url.Parse(app.AuthURI)
	if err != nil {
		return "", fmt.Errorf("Failed to parse url, %v\n", err)
	}

	//We don't care about the error here, if it doesn't work, user can always grab the link
	_ = util.OpenURL(u.String())
	fmt.Printf("Check your browser and copy/paste the given authorization code,\nif your browser didn't open use the url below:\n")
	fmt.Printf("\n%s\n\n", u)
	fmt.Print("Paste the code here:")
	getAccessTokenInput, _ := reader.ReadString('\n')
	authorizationCode := strings.TrimSpace(getAccessTokenInput)

	return authorizationCode, nil
}

func getAccessToken(ctx context.Context, authorizationCode string) (string, error) {
	client := gomasto.NewClient(mastodonClientConfig())

	if err := client.AuthenticateToken(ctx, authorizationCode, redirectUri); err != nil {
		return "", err
	}

	return client.Config.AccessToken, nil
}

func mastodonClientConfig() *gomasto.Config {
	return &gomasto.Config{
		Server:       viper.GetString("mastodon_server"),
		ClientID:     viper.GetString("mastodon_client_id"),
		ClientSecret: viper.GetString("mastodon_client_secret"),
		AccessToken:  viper.GetString("mastodon_access_token"),
	}
}

func configMastodon(ctx context.Context, reader *bufio.Reader, c *config) error {
	configMastodonServer(reader, c)
	app, err := registerMastodonApp(ctx, c)
	if err != nil {
		return fmt.Errorf("Failed to register mastodon application %v\n", err)
	}

	viper.Set("mastodon_client_id", app.ClientID)
	viper.Set("mastodon_client_secret", app.ClientSecret)

	code, err := getMastodonUserAuthorizationCode(reader, app)
	if err != nil {
		return fmt.Errorf("Failed to configure mastodon access token, %v\n", err)
	}

	accessToken, err := getAccessToken(ctx, code)
	if err != nil {
		return fmt.Errorf("Failed to get access token, %v\n", err)
	}
	viper.Set("mastodon_access_token", accessToken)

	return nil
}
