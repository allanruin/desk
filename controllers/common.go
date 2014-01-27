package controllers

import (
	"bytes"
	"crypto/sha1"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	// "github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
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
var ACCESS_TOKEN_URL = fmt.Sprintf("https://56.1.91.151:4444/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", APPID, APPSECRET)

var COSTOMER_SERVICE_URL = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", ACCESS_TOKEN)

//
var ACCESS_TOKEN = "NOT_GET_TOKEN_YET"

var WELCOME_MESSAGE = "欢迎试用人保广东桌面服务助手\n需求理清中...请稍后~ o(*￣▽￣*)o "

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

//{"access_token":"ACCESS_TOKEN","expires_in":7200}

type AccessFine struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// {"errcode":40013,"errmsg":"invalid appid"}
type AccessError struct {
	ErrCode string `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (a *AccessError) Error() string {
	return fmt.Sprintf("获取AccessToken出错，Errcode = %s,Errmsg = %s", a.ErrCode, a.ErrMsg)
}

type AccessToken struct {
	Token     string
	RecvTime  time.Time
	ExpiresIn int
}

func GetByProxy(url_addr, proxy_addr string) (*http.Response, error) {

	request, _ := http.NewRequest("GET", url_addr, nil)
	proxy, err := url.Parse(proxy_addr)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	return client.Do(request)

}

func GetAccessToken() (string, error) {
	// cl := http.Client{}
	// r, err := cl.Get(ACCESS_TOKEN_URL)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cl := &http.Client{Transport: tr}
	r, err := cl.Get(ACCESS_TOKEN_URL)

	// r, err := GetByProxy(ACCESS_TOKEN_URL, "http://56.1.200.45:8088")
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(r.Body)

	af := new(AccessFine)
	err = json.Unmarshal(body, af)

	if err != nil {
		return "", err
	} else {
		return af.AccessToken, nil
	}

	if af.AccessToken == "" {
		ae := new(AccessError)
		err := json.Unmarshal(body, ae)
		if err != nil {
			return "", err
		}

		return "", ae
	}

	return "", errors.New("获取TOKEN未知错误")
}

func init() {
	// var err error
	// ACCESS_TOKEN, err = GetAccessToken()

	// if err != nil {
	// 	beego.Error(err)
	// }

}
