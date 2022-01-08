package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"backend/app/user/internal/data/model"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	API_ENDPOINT = "https://discord.com/api/v8"
)

type discordRepo struct {
	ClientID     string
	ClientSecret string

	BotToken string

	request *gorequest.SuperAgent
}

func NewDiscordRepo(conf *conf.Config) biz.DiscordRepo {
	return &discordRepo{
		ClientID:     conf.DiscordBotClientID,
		ClientSecret: conf.DiscordBotClientSecret,
		BotToken:     conf.DiscordBotToken,
		request:      gorequest.New(),
	}
}

type DiscordTokenRequest struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	GrantType    string `json:"grant_type,omitempty"`
	Code         string `json:"code,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

type DiscordTokenResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

func (d *discordRepo) getToken(code, redirectURI string) (*DiscordTokenResponse, error) {
	reqData := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=authorization_code&code=%s&redirect_uri=%s",
		d.ClientID, d.ClientSecret, code, redirectURI)
	var resp DiscordTokenResponse
	_, _, errors := d.request.Post(API_ENDPOINT + "/oauth2/token").Type("form").
		Send(reqData).EndStruct(&resp)
	if errors != nil {
		return nil, fmt.Errorf("request discord token, err=%+v", errors[0])
	}
	return &resp, nil
}

func (d *discordRepo) GetDiscordInfo(code, redirectURI string) (*model.DiscordUser, error) {
	token, err := d.getToken(code, redirectURI)
	if err != nil {
		return nil, err
	}
	var user model.DiscordUser
	_, _, errors := d.request.Get(API_ENDPOINT+"/users/@me").
		Set("Authorization", token.TokenType+" "+token.AccessToken).
		EndStruct(&user)
	if errors != nil {
		return nil, fmt.Errorf("request discord user, err=%+v", errors[0])
	}
	return &user, nil
}

type DiscordSendMessageRequest struct {
	Content string `json:"content"`
	Tts     bool   `json:"tts"`
	Embeds  []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"embeds"`
}

type DiscordMessageResponse struct {
	Id              string    `json:"id"`
	Type            int       `json:"type"`
	Content         string    `json:"content"`
	ChannelId       string    `json:"channel_id"`
	MentionEveryone bool      `json:"mention_everyone"`
	Timestamp       time.Time `json:"timestamp"`
}

func (d *discordRepo) SendDiscordMessage(channelId string, request *DiscordSendMessageRequest) (*DiscordMessageResponse, error) {
	var resp DiscordMessageResponse
	_, _, errors := d.request.Post(API_ENDPOINT+fmt.Sprintf("/channels/%s/messages", channelId)).
		Set("Authorization", fmt.Sprintf("Bot %s", d.BotToken)).Send(*request).
		EndStruct(&resp)
	if errors != nil {
		return nil, fmt.Errorf("request discord send message, err=%+v", errors[0])
	}
	return &resp, nil
}
