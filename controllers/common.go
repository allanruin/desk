package controllers

import (
	"crypto/sha1"
	"fmt"
	"io"
)

// func CheckSignature()
var Token string = "rb95518"
var APPSECRET = "piccgdxxjsb"

// var ACCESS_TOKEN_URL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"
var ACCESS_TOKEN_URL = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=%s", APPSECRET)

func GetSha1FromStr(str string) string {
	t := sha1.New()
	io.WriteString(t, str)
	return fmt.Sprint("%x", t.Sum(nil))
}
