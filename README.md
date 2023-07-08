# HTTP CONNECT tunneling Proxy

[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## A simple http proxy service that supports authentication
## Support for using other proxies to access the final destination
## This project relies on https://github.com/mwitkow/go-http-dialer

You need to write the proxy list into redis in advance, and maintain the proxy life cycle by yourself

Redis proxy data format is zset

The service fetches a proxy from redis each time and accesses the final destination address

If you don't need to use the proxy pool to access the final target, please set the USE_PROXY_POOL parameter of .env to a value other than 1

## Instructions

Copy example.env to .env and modify the .env parameters
````
git clone git@github.com:anjingjingan/httpProxyTunnel.git
cp example.env.env
````

## Construct
````
cd httpProxyTunnel
go build .
````

# HTTP CONNECT 隧道代理

[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## 一个简单的 http 代理服务，支持验证
## 支持使用其它代理访问最终目标
## 本项目依赖了 https://github.com/mwitkow/go-http-dialer

需要事先把代理列表写入 redis ，并自行维护代理生命周期 

redis 代理数据格式为 zset 

服务每次从 redis 取一个代理，访问最终目标地址

如果不需要使用代理池访问最终目标，请把 .env 的 USE_PROXY_POOL 参数设置为不是 1 的其它值

## 使用方法

复制 example.env 为 .env ，并修改 .env 参数
````
git clone git@github.com:anjingjingan/httpProxyTunnel.git
cp example.env .env
````

## 构建
````
cd httpProxyTunnel
go build .
````