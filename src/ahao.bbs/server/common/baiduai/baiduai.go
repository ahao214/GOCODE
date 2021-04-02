package baiduai

import (
	"ahao.bbs/server/config"
	"ahao.bbs/server/model/constants"
	"github.com/mlogclub/simple/date"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"sync"
)

type ai struct {
	ApiKey                string
	SecretKey             string
	accessToken           string
	accessTokenCreateTime int64
}

var once sync.Once
var instance *ai

func GetAi() *ai {
	once.Do(func() {
		instance = &ai{
			ApiKey:    config.Instance.BaiduAi.ApiKey,
			SecretKey: config.Instance.BaiduAi.Secretkey,
		}
	})
	return instance
}

// 获取baidu api token 临时用
func (a *ai) GetToken() string {
	durationMillis := date.NowTimestamp() - a.accessTokenCreateTime
	if len(a.accessToken) == 0 || durationMillis > (86400*1000) {
		c := NewClient(a.ApiKey, a.SecretKey)
		err := c.Auth()
		if err != nil {
			logrus.Error(err)
		}
		a.accessToken = c.AccessToken
		a.accessTokenCreateTime = date.NowTimestamp()
	}
	return a.accessToken
}

func (a *ai) AnalyzeText(title, text string) (*AiAnalyzeRet, error) {
	if title == "" || text == "" {
		return nil, errors.New("内容为空")
	}
	aiCategories := a.GetCategories(title, text)
	aiTags := a.GetTags(title, text)
	summary, _ := a.GetNewsSummary(title, text, constants.SummaryLen)

	set := hashset.New()
	if aiCategories != nil {
		if len(aiCategories.Item.TopCategory) > 0 {
			for _, v := range aiCategories.Item.TopCategory {
				set.Add(v.Tag)
			}
		}
		if len(aiCategories.Item.SecondCatrgory) > 0 {
			for _, v := range aiCategories.Item.SecondCatrgory {
				set.Add(v.Tag)
			}
		}
	}
	if aiTags != nil && len(aiTags.Items) > 0 {
		for _, v := range aiTags.Items {
			set.Add(v.Tag)
		}
	}

	var tags []string
	for _, v := range set.Values() {
		tags = append(tags, v.(string))
	}
	return &AiAnalyzeRet{
		Tags:    tags,
		Summary: summary,
	}, nil
}
