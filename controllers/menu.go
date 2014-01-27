package controllers

/* https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN */

type ButtonCreate struct {
	Buttons []Buttoner `json:"button"`
}

type ButtonClick struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ButtonView struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ButtonSub struct {
	Name      string     `json:"name"`
	SubButton []Buttoner `json:"sub_button"`
}

func (b ButtonClick) IsButton() bool {
	return true
}
func (b ButtonView) IsButton() bool {
	return true
}
func (b ButtonSub) IsButton() bool {
	return true
}

type Buttoner interface {
	IsButton() bool
}
