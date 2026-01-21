中文 ｜ [English](README.md)

该项目用于将 foundationdb 监控指标透明转换为 prometheus 类型指标，每次 prometheus 拉取监控数据时向 FDB 发送 `status json` 请求并转换为 prometheus 对应的格式，中间不做任何其它处理，只做透明转发。


安装 Client 的同时会安装对应的头文件，可以参考如下内容，当前使用的是 7.1.38 版本，然后通过 dpkg 命令安装即可
https://apple.github.io/foundationdb/api-c.html

# TODO

* 支持 [Transport Layer Security](https://apple.github.io/foundationdb/tls.html) 方式连接。
