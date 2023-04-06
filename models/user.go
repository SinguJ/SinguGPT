package models

// User 用户信息
type User struct {
	// 用户 ID
	ID string
	// 用户名
	Name string `yaml:"name"`
	// 邮箱地址
	Emails []string `yaml:"emails"`
}
