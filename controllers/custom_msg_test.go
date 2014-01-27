package controllers

import (
	"bytes"
	"encoding/json"
	"testing"
)

var TextCases = []struct {
	in  interface{}
	out string
}{
	{CustomText{
		ToUser:  "oU7acuKn1S0jLyoDrds4Idl63Pg4",
		MsgType: "text",
		Text: CustomContent{
			Content: "Hello World",
		},
	},
		`{"touser":"oU7acuKn1S0jLyoDrds4Idl63Pg4","msgtype":"text","text":{"content":"Hello World"}}`},
	{CustomPic{
		ToUser:  "oU7acuKn1S0jLyoDrds4Idl63Pg4",
		MsgType: "image",
		Image: CustomMedia{
			MediaId: "MEDIA_ID",
		},
	},
		`{"touser":"oU7acuKn1S0jLyoDrds4Idl63Pg4","msgtype":"image","image":{"media_id":"MEDIA_ID"}}`},
	{CustomVoice{
		ToUser:  "oU7acuKn1S0jLyoDrds4Idl63Pg4",
		MsgType: "voice",
		Voice: CustomMedia{
			MediaId: "MEDIA_ID",
		},
	},
		`{"touser":"oU7acuKn1S0jLyoDrds4Idl63Pg4","msgtype":"voice","voice":{"media_id":"MEDIA_ID"}}`},
	{CustomVideo{
		ToUser:  "oU7acuKn1S0jLyoDrds4Idl63Pg4",
		MsgType: "video",
		Video: CustomVieoMedia{
			MediaId:     "MEDIA_ID",
			Title:       "TITLE",
			Description: "DESCRIPTION",
		},
	},
		`{"touser":"oU7acuKn1S0jLyoDrds4Idl63Pg4","msgtype":"video","video":{"media_id":"MEDIA_ID","title":"TITLE","description":"DESCRIPTION"}}`},
	{CustomMusic{
		ToUser:  "oU7acuKn1S0jLyoDrds4Idl63Pg4",
		MsgType: "music",
		Music: CustomMusicMedia{
			Title:        "MUSIC_TITLE",
			Description:  "MUSIC_DESCRIPTION",
			MusicUrl:     "MUSIC_URL",
			HqMusicUrl:   "HQ_MUSIC_URL",
			ThumbMediaId: "THUMB_MEDIA_ID",
		},
	},
		`{"touser":"oU7acuKn1S0jLyoDrds4Idl63Pg4","msgtype":"music","music":{"title":"MUSIC_TITLE","description":"MUSIC_DESCRIPTION","musicurl":"MUSIC_URL","hqmusicurl":"HQ_MUSIC_URL","thumb_media_id":"THUMB_MEDIA_ID"}}`},
	{CustomNews{
		ToUser:  "oU7acuKn1S0jLyoDrds4Idl63Pg4",
		MsgType: "news",
		News: CustomArticles{
			Articles: []CustomNewsMedia{
				{
					Title:       "Happy Day",
					Description: "Is Really A Happy Day",
					Url:         "URL",
					PicUrl:      "PIC_URLT"},
				{
					Title:       "Happy Day",
					Description: "Is Really B Happy Day",
					Url:         "URL",
					PicUrl:      "PIC_URL",
				},
			},
		},
	},
		`{"touser":"oU7acuKn1S0jLyoDrds4Idl63Pg4","msgtype":"news","news":{"articles":[{"title":"Happy Day","description":"Is Really A Happy Day","url":"URL","picurl":"PIC_URLT"},{"title":"Happy Day","description":"Is Really B Happy Day","url":"URL","picurl":"PIC_URL"}]}}`},
}

func ToJson(v interface{}) (string, error) {
	res, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	str := bytes.NewBuffer(res).String()
	return str, err
}

func TestText(t *testing.T) {
	for i, tt := range TextCases {

		str, err := ToJson(tt.in)
		if err != nil || str != tt.out {
			t.Errorf("第%d个测试没通过！", i)
			t.Log(str)
			t.Logf(tt.out)
		} else {
			t.Logf("第%d个测试pass", i)
		}

	}
}
