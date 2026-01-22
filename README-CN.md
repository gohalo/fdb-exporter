中文 ｜ [English](README.md)

该项目用于将 foundationdb 监控指标透明转换为 prometheus 类型指标，每次 prometheus 拉取监控数据时向 FDB 发送 `status json` 请求并转换为 prometheus 对应的格式，中间不做任何其它处理，只做透明转发。

# 安装

可以直接通过 `make build` 编译生成二进制文件，当前预编译的文件只支持 `x86_64` 平台。

编译需要依赖 [C API](https://apple.github.io/foundationdb/api-c.html) 库，通常安装 Client 的同时会安装对应的头文件，不同平台的安装方式可以参考 [Installing FDB Client](https://apple.github.io/foundationdb/api-general.html#installing-client-binaries) 文档，例如在 Debain/Ubuntu 平台上安装 7.1.38 版本，下载二进制文件之后通过 `dpkg -i foundationdb-clients_7.1.38-1_amd64.deb` 命令安装即可。

# 运行

连接到 FDB 依赖两个参数：

* APIVersion 默认是 `710`，可以通过参数 `--fdb_api_version=710` 或者环境变量 `export FDB_API_VERSION=710` 修改。
* ClusterFile 默认是 `/etc/foundationdb/fdb.cluster`，可以通过参数 `--fdb_cluster_file=xxx` 或者环境变量 `export FDB_CLUSTER_FILE=xxx` 修改。

然后，直接执行 `fdb-exporter` 命令即可，默认监听 `:8080` 端口，该参数只能通过命令行参数 `--addr=":8090"` 修改，访问 `/` 或者 `/metrics` 都可以。

注意，运行依赖 `fdb` 动态库，如果不在标准路径下，可以通过 `LD_LIBRARY_PATH` 环境变量指定。

如果要通过 systemd 控制运行，可以参考 [fdb-exporter.service](./contrib/fdb-exporter.service) 示例。

# 指标采集

在 Prometheus 配置时需要添加 `cluster` 这个 `Label` 用于区分不同的集群，示例如下。

``` text
- job_name: 'FDB_CLUSTER_TEST'
  metrics_path: '/'
  static_configs:
    - targets: ['your-host-name-or-ip:8080']
      labels:
        cluster: fdb_cluster_test
```

# TODO

* 支持 [Transport Layer Security](https://apple.github.io/foundationdb/tls.html) 方式连接。

# 参考

* [Architecture](https://apple.github.io/foundationdb/architecture.html)
* [Special Keys](https://apple.github.io/foundationdb/special-keys.html)
* [Monitored Metrics](https://apple.github.io/foundationdb/monitored-metrics.html)
