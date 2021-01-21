package dingtalk

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/xinzf/gokit/logger"
	"github.com/xinzf/gokit/storage"
	"time"
)

const (
	ACCESS_TOKEN_EXPIRES = int64(7000)
	ACCESS_URL           = "https://oapi.dingtalk.com"
)

const ACCESS_TOKEN_KEY = "dd:token"

var Option option

type option struct {
	AppKey, AppSecret string
	Cache             AccessTokenCache
}

var AccessToken *_accessToken

func init() {
	Option.Cache = new(defaultCache)
	if AccessToken == nil {
		AccessToken = &_accessToken{
			//cache: new(defaultCache),
		}
	}
}

type AccessTokenCache interface {
	Get(key string) (string, error)
	Set(key, token string, expiration time.Duration) error
}

type defaultCache struct {
}

func (this *defaultCache) Get(key string) (string, error) {
	key = fmt.Sprintf("%s:%s", ACCESS_TOKEN_KEY, key)
	return storage.Redis.Client().Get(key).Result()
}

func (this *defaultCache) Set(key, token string, expiration time.Duration) error {
	key = fmt.Sprintf("%s:%s", ACCESS_TOKEN_KEY, key)
	return storage.Redis.Client().Set(key, token, expiration).Err()
}

type _accessToken struct {
}

func (this *_accessToken) GetToken() (string, error) {
	if Option.AppKey == "" {
		return "", errors.New("没有设置 appKey")
	}
	if Option.AppSecret == "" {
		return "", errors.New("没有设置 appSecret")
	}

	token, err := Option.Cache.Get(Option.AppKey)
	if token != "" {
		logger.DefaultLogger.Debug("ding.Token", "命中缓存")
		return token, nil
	}
	logger.DefaultLogger.Debug("ding.Token", "没有命中缓存")

	var rsp struct {
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
	}
	_url := fmt.Sprintf("%s/gettoken?appkey=%s&appsecret=%s", ACCESS_URL, Option.AppKey, Option.AppSecret)
	_rsp, _, errs := gorequest.New().Get(_url).EndStruct(&rsp)
	if _rsp.StatusCode != 200 {
		return "", fmt.Errorf("钉钉服务器异常,httpStatus: %d", _rsp.StatusCode)
	}

	if len(errs) > 0 {
		return "", errs[0]
	}

	if rsp.Errcode != 0 {
		return "", errors.New(rsp.Errmsg)
	}

	err = Option.Cache.Set(
		Option.AppKey,
		rsp.AccessToken,
		time.Duration(ACCESS_TOKEN_EXPIRES)*time.Second,
	)
	if err != nil {
		return "", err
	}
	return rsp.AccessToken, nil
}
