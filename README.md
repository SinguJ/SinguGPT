# SinguGPT

这是一个通过电子邮件对接 ChatGPT 的 Go 语言程序

## 安装

以下命令将以 Linux 操作系统为例，其他操作系统请在理解下述命令后，手动操作。

1. 下载源代码

```shell
git clone git@github.com:singu-tech/SinguGPT.git
cd SinguGPT
```

2. 编译程序

```shell
.\GoPackage.ps1 -Publish -TargetSystem "<操作系统>" -TargetArch "<CPU 架构>" -Version "v1.0.0-beta"
```

> * `GoPackage.ps1` 是一个 Powershell 脚本，您需要使用 Powershell 7+ 运行它。
> * 请修改上述命令中的 `<操作系统>`，将其改为您期望运行本程序的服务器的系统类型；Windows 系统请填写 `Win`，`Linux` 系统请填写 `Linux`，`MacOS` 系统请填写 `MacOS`。
> * 请修改上述命令中的 `<CPU 架构>`，将其改为您期望运行本程序的服务器的 CPU 架构类型；英特尔 64 位架构请填写 `x64`，ARM 64 位架构请填写 `arm64`
> * 建议您在执行编译前，先安装 UPX 压缩程序：[UPX](https://upx.github.io)

完成该步骤后，将在 dist 目录下生成一个文件名类似 `SinguGPT_v1.0.0-beta_<操作系统>-<CPU 架构>.<压缩格式>` 的文件。

3. 部署程序

```shell
# 本地计算机
scp dist/SinguGPT_v1.0.0-beta_<操作系统>-<CPU 架构>.<压缩格式> root@<远程计算机>:~/

# 远程计算机
cd ~/
tar zxvf SinguGPT_v1.0.0-beta_<操作系统>-<CPU 架构>.<压缩格式>
cp -a SinguGPT /opt/SinguGPT
cd /opt/SinguGPT
cp -a config.yaml.template config.yaml
cp -a user.yaml.template user.yaml
chmod u+x SinguGPT
```

4. 修改配置文件

您需要修改 `config.yaml` 与 `user.yaml` 配置文件：

1. 在 `config.yaml` 中，添加 OpenAI 的 API Key，以及电子邮箱的 IMAP、SMTP 登录信息
2. 在 `user.yaml` 中，根据示例录入您的用户

5. 启动服务

```shell
./SinguGPT >> server.log
```
