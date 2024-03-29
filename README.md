# with-env
[![check](https://github.com/hellojukay/with-env/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/hellojukay/with-env/actions/workflows/go.yml)

通过 .env 文件给命令注入环境变量, 会覆盖系统自带的环境变量。

1. ~/.env
2. .env

将下面文件写入到 `.env` 或者 `~/.env` 中，然后执行 `with-env wget` ， `wget` 命令就被注入了代理。
```
https_proxy=http://127.0.0.1:7890
http_proxy=http://127.0.0.1:7890
no_proxy=localhost,127.0.0.1

HTTPS_PROXY=http://127.0.0.1:7890
HTTP_PROXY=http://127.0.0.1:7890
NO_PROXY=localhost,127.0.0.1
```
使用方式
```
with-env curl ....
with-env wget ....
with-env git clone ...
````

# Install
```
go install github.com/hellojukay/with-env@latest
```

# Download
[https://github.com/hellojukay/with-env/releases](https://github.com/hellojukay/with-env/releases)
