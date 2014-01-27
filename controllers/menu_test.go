package controllers

import (
	"testing"
)

var MenuCase = []struct {
	in  ButtonCreate
	out string
}{
	{
		ButtonCreate{
			Buttons: []Buttoner{
				ButtonClick{
					Type: "click",
					Name: "今日歌曲",
					Key:  "V1001_TODAY_MUSIC",
				}, // end of a button
				ButtonClick{
					Type: "click",
					Name: "歌手简介",
					Key:  "V1001_TODAY_SINGER",
				},
				ButtonSub{
					Name:      "菜单",
					SubButton: []Buttoner{ButtonView{"view", "搜索", "http://www.soso.com/"}, ButtonView{"view", "视频", "http://v.qq.com/"}, ButtonClick{"click", "赞一下我们", "V1001_GOOD"}},
				},
			},
		},
		`{"button":[{"type":"click","name":"今日歌曲","key":"V1001_TODAY_MUSIC"},{"type":"click","name":"歌手简介","key":"V1001_TODAY_SINGER"},{"name":"菜单","sub_button":[{"type":"view","name":"搜索","url":"http://www.soso.com/"},{"type":"view","name":"视频","url":"http://v.qq.com/"},{"type":"click","name":"赞一下我们","key":"V1001_GOOD"}]}]}`,
	}, // end of a case
}

// 汗，原来请求查询接口查询到的menu跟请求时用的结构不一样。。囧
func TestMenu(t *testing.T) {
	for i, tt := range MenuCase {

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
