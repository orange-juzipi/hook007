package model

type User struct {
	Model
	NickName  string `json:"nickname"`
	Email     string `json:"email" gorm:"not null;type:varchar(64);"`
	Password  string `json:"password" gorm:"not null;type:varchar(255);"`
	SecretKey string `json:"secretKey"`
}
