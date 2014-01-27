package controllers

import (
	// "fmt"
	_ "github.com/go-sql-driver/mysql" // what the fuck we need this in go test.
)

// 腾讯文档地址：
// http://mp.weixin.qq.com/wiki/index.php?title=%E5%8F%91%E9%80%81%E5%AE%A2%E6%9C%8D%E6%B6%88%E6%81%AF

/*
发送文本消息

{
    "touser":"OPENID",
    "msgtype":"text",
    "text":
    {
         "content":"Hello World"
    }
}
*/
type CustomContent struct {
	Content string `json:"content"`
}

type CustomText struct {
	ToUser  string        `json:"touser"`
	MsgType string        `json:"msgtype"`
	Text    CustomContent `json:"text"`
}

/*
发送图片消息

{
    "touser":"OPENID",
    "msgtype":"image",
    "image":
    {
      "media_id":"MEDIA_ID"
    }
}
*/
type CustomMedia struct {
	MediaId string `json:"media_id"`
}

type CustomPic struct {
	ToUser  string      `json:"touser"`
	MsgType string      `json:"msgtype"`
	Image   CustomMedia `json:"image"`
}

/*发送语音消息
{
    "touser":"OPENID",
    "msgtype":"voice",
    "voice":
    {
      "media_id":"MEDIA_ID"
    }
}
*/
type CustomVoice struct {
	ToUser  string      `json:"touser"`
	MsgType string      `json:"msgtype"`
	Voice   CustomMedia `json:"voice"`
}

/*
发送视频消息

{
    "touser":"OPENID",
    "msgtype":"video",
    "video":
    {
      "media_id":"MEDIA_ID",
      "title":"TITLE",
      "description":"DESCRIPTION"
    }
}
*/
type CustomVieoMedia struct {
	MediaId     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CustomVideo struct {
	ToUser  string          `json:"touser"`
	MsgType string          `json:"msgtype"`
	Video   CustomVieoMedia `json:"video"`
}

/*
发送音乐消息

{
    "touser":"OPENID",
    "msgtype":"music",
    "music":
    {
      "title":"MUSIC_TITLE",
      "description":"MUSIC_DESCRIPTION",
      "musicurl":"MUSIC_URL",
      "hqmusicurl":"HQ_MUSIC_URL",
      "thumb_media_id":"THUMB_MEDIA_ID"
    }
}
*/
type CustomMusicMedia struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"musicurl"`
	HqMusicUrl   string `json:"hqmusicurl"`     //高品质音乐链接，wifi环境优先使用该链接播放音乐
	ThumbMediaId string `json:"thumb_media_id"` // 缩略图的媒体ID
}
type CustomMusic struct {
	ToUser  string           `json:"touser"`
	MsgType string           `json:"msgtype"`
	Music   CustomMusicMedia `json:"music"`
}

/*
发送图文消息

图文消息条数限制在10条以内，注意，如果图文数超过10，则将会无响应。

{
    "touser":"OPENID",
    "msgtype":"news",
    "news":{
        "articles": [
         {
             "title":"Happy Day",
             "description":"Is Really A Happy Day",
             "url":"URL",
             "picurl":"PIC_URL"
         },
         {
             "title":"Happy Day",
             "description":"Is Really A Happy Day",
             "url":"URL",
             "picurl":"PIC_URL"
         }
         ]
    }
}
*/

type CustomArticles struct {
	Articles []CustomNewsMedia `json:"articles"`
}

type CustomNewsMedia struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}
type CustomNews struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	News    CustomArticles `json:"news"`
}
