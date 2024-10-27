<p align="center">
  <img src="https://raw.githubusercontent.com/coolapso/megophone/refs/heads/dev/media/megophone.png" width="200" >
</p>

# megophone

A single tool for multiple social networks.

Megaphone allows you to post to multiple social networks simultaneously from your CLI.
<p align="center">
  <img src="https://raw.githubusercontent.com/coolapso/megophone/refs/heads/main/media/usage.gif">
</p>

## Features

* Configuration utility: use `megaphone configure` to set up the tool
* Multiple configuration profiles: use the `--profile` flag to setup and use different accounts
* Post to all supported social networks
* Post to all supported social networkds with images and videos
* Post only to X: use `megaphone -x "text"` to post only to X
* Post to Mastodon: use `megaphone -m "text"` to post only to Mastodon


## Supported social netrowks
* X
* Mastodon

### Planed features

* Threads
* Facebook (Still not very sure about this one)
* Threading, split longer texts and post them as threads
* Polls for Mastodon and X

## Installation 

### AUR

On Arch linux you can use the AUR `megophone-bin`

### Go Install

#### Latest version 

`go install github.com/coolapso/megophone`

#### Specific version

`go install github.com/coolapso/megophone@v1.0.0`

### Linux Script

It is also impossible to install on any linux distro with the installation script

#### Latest version

```
curl -L https://megophone.coolapso.sh/install.sh | bash
```

#### Specific version

```
curl -L https://megophone.coolapso.sh/install.sh | VERSION="v1.1.0" bash
```

### Manual install

* Grab the binary from the [releases page](https://github.com/coolapso/megophone/releases).
* Extract the binary
* Execute it

## Setup

Megophone needs access to your API Keys and Access tokens. For that, Megophone provides a configuration utility, which you can start with `megophone configure`. However, there are some steps you may need to do first. Once Megophone is configured, a configuration file with the necessary tokens and secrets is saved in `$XDG_CONFIG_HOME/megophone/config.yaml`.

### X.com

* Create an X developer account at: https://developer.x.com/en
* Create a new app "megophone"
* Generate tokens; Megophone needs read and write permissions. The necessary tokens are: 
    * API Key
    * API Key Secret
    * OAuth Token
    * OAuth Token Secret
* Provide the tokens when requested by the `megophone configure` command

These tokens can also be provided with the following environment variables:
`MEGOPHONE_X_API_KEY`, `MEGOPHONE_X_API_KEY_SECRET`, `MEGOPHONE_X_OAUTH_TOKEN`, `MEGOPHONE_X_OAUTH_TOKEN_SECRET`

> [!NOTE]  
> You are subject to Twitter API pricing and limits. Please make sure to check the X developer portal information: https://developer.x.com/en

### Mastodon 

* Mastodon configuration is all done through `megaphone configure`. During the process, it will open your browser and request you to paste the authorization code.

## Usage 

```
Post to multiple social networks from your CLI

Usage:
  megophone [flags]
  megophone [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  configure   Configures megophone
  help        Help about any command

Flags:
      --config string       config file (default is $XDG_HOME_CONFIG/megophone/config.yaml)
  -h, --help                help for megophone
  -m, --m-only              Post to Mastodon Only
  -p, --media-path string   Path of media to be uploaded
      --profile string      The configuration profile to use (default "default")
  -x, --x-only              Post to X only

Use "megophone [command] --help" for more information about a command.
```

### Profiles

Profiles are distinct sets of accounts designated for posting. For instance, you can create a profile for your personal accounts and another for your work accounts. By using the `--profile` flag, you can effortlessly switch between these profiles when posting. If the flag is not specified, the default profile will be used.

## Build 

### With makefile

`make build`

### Manually

`go build -o megophone`

# Contributions

Improvements and suggestions are always welcome, feel free to check for any open issues, open a new Issue or Pull Request

If you like this project and want to support / contribute in a different way you can always: 

<a href="https://www.buymeacoffee.com/coolapso" target="_blank">
  <img src="https://cdn.buymeacoffee.com/buttons/default-yellow.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important;" />
</a>
