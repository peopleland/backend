package data

import (
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterRepo struct {
	client *twitter.Client
}

func NewTwitterRepo(conf *conf.Config) biz.TwitterRepo {
	config := oauth1.NewConfig(conf.TwitterConsumerKey, conf.TwitterConsumerSecret)
	token := oauth1.NewToken(conf.TwitterToken, conf.TwitterTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	return &TwitterRepo{
		client: client,
	}
}

func (repo *TwitterRepo) GetTwitterUserTimeline(userScreenName string) []string {
	var result []string
	tweets, _, err := repo.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: userScreenName,
	})
	if err != nil {
		panic(err)
	}
	for _, tweet := range tweets {
		result = append(result, tweet.Text)
	}
	return result
}
