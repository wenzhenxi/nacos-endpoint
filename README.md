# nacos-endpoint

阿里云将于2022年3月31号关闭对ACM的支持，需要切换到MSE中的专属Nacos，但是对于已经对接了ACM的用户无法实现平滑切换，必须要对代码进行修改。这就导致了如果项目数量较多切换成本遗漏检查成本会非常高，但是从原理上来说ACM和Nacos差异仅仅是有一个endpoint来获取service地址和直接固定service地址，和官方进行交流下来并不打算提供这样的能力。所以通过一个API返回service地址注册到endpoint就可以解决问题。这个项目通过请求域名 nacos实例名.nacos.xxxx.com来实现一个通用的endpoint来解决兼容性问题。

## 使用方式

Docker部署：
```shell
docker run -p 8080:8080 sunmi-docker-images-registry.cn-hangzhou.cr.aliyuncs.com/public/nacos-endpoint
```

请求：
- 实例名称.域名
返回：
- 实例名称.mse.aliyuncs.com
