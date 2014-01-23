package controllers

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

// func CheckSignature()
var Token string = "rb95518"

// APPID是要服务号才有的。。
// 所以要支持客服接口的话，要获取高级接口都要有ACCESS_TOKEN,
// 而获取ACCESS_TOKEN需要是APPID进行验证，所以没有APPID是没法用高级接口的
//
var APPSECRET = "2fd6de79072bf9c824de052efb5b2dd8"
var APPID = "wx25015e91334beab2"

// var ACCESS_TOKEN_URL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"
var ACCESS_TOKEN_URL = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", APPID, APPSECRET)

func GetSha1FromStr(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (this *APIController) CheckSignature() bool {
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	tmpArr := []string{Token, timestamp, nonce}
	sort.Strings(tmpArr)
	joinstr := strings.Join(tmpArr, "")
	sha1str := GetSha1FromStr(joinstr)

	if sha1str == signature {
		return true
	}

	// // debugging info
	// echostr := this.GetString("echostr")
	// beego.Info("signature: " + signature)
	// beego.Info("timestamp: " + timestamp)
	// beego.Info("nonce: " + nonce)
	// beego.Info("echostr: " + echostr)
	// beego.Trace("trace before sort: ", tmpArr)
	// beego.Info("joinstr: " + joinstr)
	// beego.Info("sha1str: " + sha1str)

	return false
}

// 将[]byte 直接数组转换成字符串
func BtyeToString(bs []byte) string {
	buf := bytes.NewBuffer(bs)
	return buf.String()
}

func Filter(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}
