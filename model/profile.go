package model

type Profile struct {
	UserName  string   // 用户名
	Category  string   // 分组类别 “大胸，长腿”
	Avatar    string   // 用户头像
	PushTitle string   // 帖子标题
	PushAt    string   // 帖子发送时间
	PushText  []string // 帖子内文字
	PushImgs  []string // 帖子内图片
}
