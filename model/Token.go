package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/util"
	"time"
)

type Token struct {
	UserId     int         `json:"userId"`
	Payload    interface{} `json:"payload"`
	ExpireTime int64       `json:"expireTime"`
	Sign       string      `json:"sign"`
}

func NewToken(userId int, payload interface{}) *Token {
	token := &Token{
		UserId:     userId,
		Payload:    payload,
		ExpireTime: time.Now().Unix(),
	}
	token.sign()
	return token
}

func (t *Token) Encode() string {
	data, _ := json.Marshal(t)
	return base64.RawURLEncoding.EncodeToString(data)
}

func DecodeToken(result string) *Token {
	token := &Token{}
	data, err := base64.RawURLEncoding.DecodeString(result)
	if err != nil {
		return token
	}
	err = json.Unmarshal(data, token)
	if err != nil {
		return token
	}
	return token
}

func (t *Token) sign() {
	data := fmt.Sprintf("%v,%v,%v", t.UserId, t.Payload, t.ExpireTime)
	h := util.Hash(data)
	if !config.Key.Set {
		config.Key.Pk, config.Key.Sk = util.GenRsaKey(2048)
		config.Key.Set = true
	}
	sign, _ := util.RsaSign(h, config.Key.Sk)
	t.Sign = sign
}

func (t *Token) Verify() bool {
	data := fmt.Sprintf("%v,%v,%v", t.UserId, t.Payload, t.ExpireTime)
	h := util.Hash(data)
	if err := util.RsaVerifySign(h, t.Sign, config.Key.Pk); err != nil {
		return false
	}
	return true
}
