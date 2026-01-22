English | [中文](README-CN.md)

This project is designed to transparently convert FoundationDB monitoring metrics into Prometheus-compatible metrics. It sends `status json` request to FDB each time Prometheus pulls monitoring data, converts the response into the corresponding Prometheus format, and performs only transparent forwarding without any additional processing in between.

# Installation

You can compile the binary file directly with the `make build` command. Currently, precompiled binaries only support the `x86_64` platform.

Compilation requires the [C API](https://apple.github.io/foundationdb/api-c.html) library. The corresponding header files are usually installed alongside the FDB Client. For installation methods on different platforms, refer to the [Installing FDB Client](https://apple.github.io/foundationdb/api-general.html#installing-client-binaries) documentation. For example, to install version 7.1.38 on Debian/Ubuntu, download the binary package and run the command `dpkg -i foundationdb-clients_7.1.38-1_amd64.deb`.

# Running

Connecting to FDB relies on two parameters:

* APIVersion: default `710`. It can be modified via the command-line parameter `--fdb_api_version=710` or the environment variable `export FDB_API_VERSION=710`.
* ClusterFile: default `/etc/foundationdb/fdb.cluster`. It can be modified via the command-line parameter `--fdb_cluster_file=xxx` or the environment variable `export FDB_CLUSTER_FILE=xxx`.

And then execute the `fdb-exporter` command to start the service, which listens on port `:8080` by default. This port can only be modified via the command-line parameter `--addr=":8090"`. Metrics can be accessed via either the `/` or `/metrics` endpoint.

Note: The service depends on the fdb dynamic library. If the library is not in the standard system path, specify its location using the `LD_LIBRARY_PATH` environment variable.

To manage the service with systemd, refer to the example configuration file [fdb-exporter.service](./contrib/fdb-exporter.service).

# Metric Collection

When configuring Prometheus, you need to add the `cluster` label to distinguish between different FDB clusters. See the example configuration below:

``` text
- job_name: 'FDB_CLUSTER_TEST'
  metrics_path: '/'
  static_configs:
    - targets: ['your-host-name-or-ip:8080']
      labels:
        cluster: fdb_cluster_test
```

# TODO

* Add support for connections via [Transport Layer Security](https://apple.github.io/foundationdb/tls.html).

# References

* [Architecture](https://apple.github.io/foundationdb/architecture.html)
* [Special Keys](https://apple.github.io/foundationdb/special-keys.html)
* [Monitored Metrics](https://apple.github.io/foundationdb/monitored-metrics.html)
