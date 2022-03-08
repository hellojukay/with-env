# with-proxy
给命令加上网络代理
```
https_proxy=http://127.0.0.1:7890
http_proxy=http://127.0.0.1:7890
HTTPS_PROXY=http://127.0.0.1:7890
HTTP_PROXY=http://127.0.0.1:7890
```
使用方式
```
with-proxy curl ....
with-proxy wget ....
with-proxy git clone ...
````

# Install
```
go install github.com/hellojukay/with-proxy@latest
```