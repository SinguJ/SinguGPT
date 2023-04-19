package store

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v2"
)

// 读取 YAML
func loadYaml(path string, obj interface{}) error {
    // 判断配置文件是否存在，若不存在则抛出错误
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return fmt.Errorf("配置文件 %s 不存在", path)
    }

    // 读取 YAML 配置文件
    data, err := os.ReadFile(path)
    if err != nil {
        return fmt.Errorf("读取配置文件失败: %v", err)
    }

    if err := yaml.Unmarshal(data, obj); err != nil {
        return fmt.Errorf("解析配置文件失败: %v", err)
    }

    return nil
}
