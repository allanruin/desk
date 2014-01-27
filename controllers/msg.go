package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

//用户发过来的消息格式
type RecvTextMsg struct {
	// XMLName      xml.Name `xml:"xml"`
	ToUserName   string `xml:"ToUserName"`
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	MsgId        int64 //消息id，64位整型
}

// type RetTextMsg BaseTextMsg

type RetTextMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
}

/*
<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>12345678</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[你好]]></Content>
</xml>
// the CDATA is not so important
*/

func (r *RecvTextMsg) Return(s string) *RetTextMsg {
	rt := new(RetTextMsg)
	rt.FromUserName = r.ToUserName
	rt.ToUserName = r.FromUserName
	rt.MsgType = r.MsgType
	rt.CreateTime = fmt.Sprint(time.Now().Unix())
	rt.Content = s
	return rt
}

func (rc *RecvTextMsg) Process(state string) *RetTextMsg {
	ProcessFunc, ok := StateFuncs[state]
	rt := new(RetTextMsg)
	if ok {
		rt = ProcessFunc(rc)
	} else {
		beego.Error("state function map error")
	}

	return rt
}

/*
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[subscribe]]></Event>
*/

type EventMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Event        string
}

func (r *EventMsg) Return(s string) *RetTextMsg {
	rt := new(RetTextMsg)
	rt.FromUserName = r.ToUserName
	rt.ToUserName = r.FromUserName
	rt.MsgType = "text"
	rt.CreateTime = fmt.Sprint(time.Now().Unix())
	rt.Content = s
	return rt
}

// func (a *AccessToken) GetToken() string {
// 	if time.Now().Unix
// }

/*
<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>1357290913</CreateTime>
<MsgType><![CDATA[voice]]></MsgType>
<MediaId><![CDATA[media_id]]></MediaId>
<Format><![CDATA[Format]]></Format>
<Recognition><![CDATA[腾讯微信团队]]></Recognition>
<MsgId>1234567890123456</MsgId>
</xml>
*/
type VoiceMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	MediaId      string //语音消息媒体id，可以调用多媒体文件下载接口拉取该媒体
	Format       string //语音格式：amr
	Recognition  string //语音识别结果，UTF8编码
	MsgId        string
}

func (r *VoiceMsg) Return(s string) *RetTextMsg {
	rt := new(RetTextMsg)
	rt.FromUserName = r.ToUserName
	rt.ToUserName = r.FromUserName
	rt.MsgType = "text"
	rt.CreateTime = fmt.Sprint(time.Now().Unix())
	rt.Content = s

	return rt
}

/*
<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[CLICK]]></Event>
<EventKey><![CDATA[EVENTKEY]]></EventKey>
</xml>
*/

type MenuMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Event        string
	EventKey     string
}

func (r *MenuMsg) Return(s string) *RetTextMsg {
	rt := new(RetTextMsg)
	rt.FromUserName = r.ToUserName
	rt.ToUserName = r.FromUserName
	rt.MsgType = "text"
	rt.CreateTime = fmt.Sprint(time.Now().Unix())
	rt.Content = s

	return rt
}

// 客服消息
// https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
