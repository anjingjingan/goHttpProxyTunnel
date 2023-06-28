## 一个简单的 http 代理服务，支持验证
## 支持在服务里面，使用其它代理访问最终目标

## 使用方法

复制 example.env 为 .env ，并修改 .env 参数
````
git clone git@github.com:anjingjingan/goHttpProxyTunnel.git
cp example.env .env
````

## 构建
````
cd goHttpProxyTunnel
go build .
````