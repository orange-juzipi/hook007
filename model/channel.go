package model

type Channel struct {
	Model
	Type         ChannelType `json:"type" gorm:"not null;type:varchar(32);comment:类型"`
	UserID       uint        `json:"userID" gorm:"comment:用户ID"`
	Configration string      `json:"configuration" gorm:"type:text;comment:配置项"`
}

type ChannelType string

const (
	QW       ChannelType = "qw"       // 企微
	DingTalk ChannelType = "dingtalk" // 钉钉
	Lark     ChannelType = "lark"     // 飞书
	Email    ChannelType = "email"    // 邮件
	Discord  ChannelType = "discord"  // Discord
	Slack    ChannelType = "slack"    // Slack
	Telegram ChannelType = "telegram" // Telegram
	Matrix   ChannelType = "matrix"   // Matrix
	Webhook  ChannelType = "webhook"  // Webhook
)
