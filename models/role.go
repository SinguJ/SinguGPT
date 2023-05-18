package models

// Role 角色信息
type Role struct {
    // 角色名
    Name string `yaml:"name"`
    // 可用权限
    Perms []string `yaml:"perms"`
}
