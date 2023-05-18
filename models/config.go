package models

type Config struct {
    // 应用相关配置
    App struct {
        // 应用名称
        Name string `yaml:"name"`
        // 用户配置文件
        UserDataFile string `yaml:"user-data-file"`
        // 接入方式
        AccessMethods []AccessMethod `yaml:"access-methods"`
    } `yaml:"app"`
    // OpenAI 相关配置
    OpenAI []struct {
        // ID
        ID string `yaml:"id"`
        // OpenAI Org ID
        OrgId string `yaml:"org-id"`
        // OpenAI API Key
        ApiKey string `yaml:"apikey"`
        // 支持的 OpenAI 模型
        Models []OpenAiModel `yaml:"models"`
        // 激活命令集
        Commands []string `yaml:"commands"`
    } `yaml:"openai"`
    // 邮箱配置
    Email struct {
        // SMTP 配置
        SMTP struct {
            // 服务器地址
            Host string `yaml:"host"`
            // 端口号
            Port int `yaml:"port"`
            // 用户名
            UserName string `yaml:"username"`
            // 密码
            Password string `yaml:"password"`
        } `yaml:"smtp"`
        // IMAP 配置
        IMAP struct {
            // 服务器地址
            Host string `yaml:"host"`
            // 端口号
            Port int `yaml:"port"`
            // 用户名
            UserName string `yaml:"username"`
            // 密码
            Password string `yaml:"password"`
        } `yaml:"imap"`
    } `yaml:"email"`
}

type UserConfig struct {
    Roles []*Role          `yaml:"roles"`
    Users map[string]*User `yaml:"users"`
}
