package models

import (
    "testing"
)

var testMarkdown string

func init() {
    testMarkdown = "# SinguGPT\n" +
        "\n" +
        "这是一个通过电子邮件对接 ChatGPT 的 Go 语言程序\n" +
        "\n" +
        "## 安装\n" +
        "\n" +
        "以下命令将以 Linux 操作系统为例，其他操作系统请在理解下述命令后，手动操作。\n" +
        "\n" +
        "1. 下载源代码\n" +
        "\n" +
        "```bash\n" +
        "git clone git@github.com:singu-tech/SinguGPT.git\n" +
        "cd SinguGPT\n" +
        "```\n" +
        "\n" +
        "2. 编译程序\n" +
        "\n" +
        "```powershell\n" +
        ".\\GoPackage.ps1 -Publish -TargetSystem \"<操作系统>\" -TargetArch \"<CPU 架构>\" -Version \"v1.0.0-beta\"\n" +
        "```\n" +
        "\n" +
        "> * `GoPackage.ps1` 是一个 Powershell 脚本，您需要使用 Powershell 7+ 运行它。\n" +
        "> * 请修改上述命令中的 `<操作系统>`，将其改为您期望运行本程序的服务器的系统类型；Windows 系统请填写 `Win`，`Linux`\n" +
        "    系统请填写 `Linux`，`MacOS` 系统请填写 `MacOS`。\n" +
        "> * 请修改上述命令中的 `<CPU 架构>`，将其改为您期望运行本程序的服务器的 CPU 架构类型；英特尔 64 位架构请填写 `x64`，ARM 64\n" +
        "    位架构请填写 `arm64`\n" +
        "> * 建议您在执行编译前，先安装 UPX 压缩程序：[UPX](https://upx.github.io)\n" +
        "\n" +
        "完成该步骤后，将在 dist 目录下生成一个文件名类似 `SinguGPT_v1.0.0-beta_<操作系统>-<CPU 架构>.<压缩格式>` 的文件。\n" +
        "\n" +
        "3. 部署程序\n" +
        "\n" +
        "```bash\n" +
        "# 本地计算机\n" +
        "scp dist/SinguGPT_v1.0.0-beta_<操作系统>-<CPU 架构>.<压缩格式> root@<远程计算机>:~/\n" +
        "\n" +
        "# 远程计算机\n" +
        "cd ~/\n" +
        "tar zxvf SinguGPT_v1.0.0-beta_<操作系统>-<CPU 架构>.<压缩格式>\n" +
        "cp -a SinguGPT /opt/SinguGPT\n" +
        "cd /opt/SinguGPT\n" +
        "cp -a config.yaml.template config.yaml\n" +
        "cp -a user.yaml.template user.yaml\n" +
        "chmod u+x SinguGPT\n" +
        "```\n" +
        "\n" +
        "4. 修改配置文件\n" +
        "\n" +
        "您需要修改 `config.yaml` 与 `user.yaml` 配置文件：\n" +
        "\n" +
        "1. 在 `config.yaml` 中，添加 OpenAI 的 API Key，以及电子邮箱的 IMAP、SMTP 登录信息\n" +
        "2. 在 `user.yaml` 中，根据示例录入您的用户\n" +
        "\n" +
        "5. 启动服务\n" +
        "\n" +
        "```bash\n" +
        "./SinguGPT >> server.log\n" +
        "```\n"
}

func TestMarkdownToHTML(t *testing.T) {
    markdown := NewMarkdownContent(TagBody, testMarkdown)
    _ = MarkdownToHTML(markdown)
}
