package models

import (
    "fmt"

    markdownLib "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/html"
)

// 启用的 HTML 特性
const htmlFlags = 0 |
    html.HrefTargetBlank |
    html.FootnoteReturnLinks |
    html.FootnoteNoHRTag |
    html.Smartypants |
    html.SmartypantsFractions |
    html.SmartypantsDashes |
    html.SmartypantsLatexDashes |
    html.LazyLoadImages

const htmlFormat = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html;charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/5.2.0/github-markdown.min.css">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/5.2.0/github-markdown-light.min.css">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/default.min.css">
<script type="text/javascript" src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/highlight.min.js"></script>
<script type="text/javascript">hljs.highlightAll();</script>
<style>
	.markdown-body {
		box-sizing: border-box;
		min-width: 200px;
		max-width: 980px;
		margin: 0 auto;
		padding: 45px;
	}

	@media (max-width: 767px) {
		.markdown-body {
			padding: 15px;
		}
	}
</style>
</head>
<body>
<article class="markdown-body">
%s
</article>
</body>
`

// MarkdownToHTML Markdown 类型的内容转换为 HTML 类型的内容
func MarkdownToHTML(markdown *MarkdownContent) *HTMLContent {
    opts := html.RendererOptions{Flags: htmlFlags}
    renderer := html.NewRenderer(opts)
    return NewHTMLContent(markdown.tag, fmt.Sprintf(htmlFormat, markdownLib.Render(markdown.ast, renderer)))
}
