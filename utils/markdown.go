package utils

import "regexp"

// 正则规则
var regexRules = []struct {
    pattern *regexp.Regexp
    replace string
}{
    // 在英文字符与汉字之间插入空格
    {regexp.MustCompile(`([a-zA-Z])(\p{Han})`), "$1 $2"},
    {regexp.MustCompile(`(\p{Han})([a-zA-Z])`), "$1 $2"},
    // 在英文符号与汉字之间插入空格
    {regexp.MustCompile(`([[:punct:]])(\p{Han})`), "$1 $2"},
    {regexp.MustCompile(`(\p{Han})([[:punct:]])`), "$1 $2"},
    // 在数字与汉字之间插入空格
    {regexp.MustCompile(`([0-9])(\p{Han})`), "$1 $2"},
    {regexp.MustCompile(`(\p{Han})([0-9])`), "$1 $2"},
}

func BeautifulMarkdown(content string) string {
    // 通过正则表达式规则集优化输入的 Markdown 原文
    for _, rule := range regexRules {
        content = rule.pattern.ReplaceAllString(content, rule.replace)
    }
    return content
}
