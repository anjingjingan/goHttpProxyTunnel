## 一个简单的 http 代理服务，支持验证
## 支持在服务里面，使用其它代理访问最终目标

需要事先把代理列表写入 redis ，并自行维护代理生命周期 

redis 代理数据格式为 zset 

服务每次从 redis 取一个代理，访问最终目标地址

如果不需要使用代理池访问最终目标，请把 .env 的 USE_PROXY_POOL 参数设置为不是 1 的其它值

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