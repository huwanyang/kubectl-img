[![ci-release](https://github.com/huwanyang/kubectl-img/actions/workflows/ci-release.yml/badge.svg?branch=master)](https://github.com/huwanyang/kubectl-img/actions/workflows/ci-release.yml)
[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/huwanyang/kubectl-img)
[![GitHub License](https://img.shields.io/github/license/huwanyang/kubectl-img?color=brightgreen&logo=apache)](https://github.com/huwanyang/kubectl-img/blob/master/LICENSE)

kubectl-img 是基于 Cobra 脚手架创建的命令插件，可以显示指定的 k8s 资源类型的 image 镜像信息，例如：deployments | daemonsets | statefulsets | jobs | cronjobs，
并支持多种输出方式，例如：json | yaml | xml | table。

## 安装说明

#### Linux

```shell script
$ export version=1.0.0
$ curl -L -o kubectl-img.tar.gz https://github.com/huwanyang/kubectl-img/releases/download/v${version}/kubectl-img_${version}_Linux_x86_64.tar.gz
$ tar -xvf kubectl-img.tar.gz
$ cp kubectl-img /usr/local/bin/kubectl-img
# 使用 krew 管理 kubectl plugin
$ cp kubectl-img $HOME/.krew/bin/
```

#### OSX

```shell script
$ export version=1.0.0
$ curl -L -o kubectl-img.tar.gz https://github.com/huwanyang/kubectl-img/releases/download/v${version}/kubectl-img_${version}_Darwin_x86_64.tar.gz
$ tar -xvf kubectl-img.tar.gz
$ cp kubectl-img /usr/local/bin/kubectl-img
# 使用 krew 管理 kubectl plugin
$ cp kubectl-img $HOME/.krew/bin/
```

#### 源码安装

```shell script
# 下载并编译 kubectl-img 到本地 $GOPATH/bin 下，可直接使用 kubectl-img 命令
$ GO111MODULE=on go get github.com/huwanyang/kubectl-img@latest
# 使用 krew 管理 kubectl plugin
$ cp $GOPATH/bin/kubectl-img $HOME/.krew/bin/
```

## 使用说明

```shell script
$ kubectl-img image -h
image 命令可以显示指定的 k8s 资源类型镜像，例如: deployments|daemonsets|statefulsets|jobs|cronjobs. 
同时可以指定输出格式，例如: json|yaml|xml|table

Usage:
  kubectl-img image [flags]

Flags:
  -c, --cronjobs        Show resource cronjobs images
  -a, --daemonsets      Show resource daemonsets images
  -d, --deployments     Show resource deployment images
  -h, --help            help for image
  -j, --jobs            Show resource jobs images
  -o, --output string   Output format. One of: json|yaml|xml|table  (default "table")
  -f, --statefulsets    Show resource statefulsets images

Global Flags:
      --as string                      Username to impersonate for the operation
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --cache-dir string               Default cache directory (default "/Users/huwanyang/.kube/cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
  -n, --namespace string               If present, the namespace scope for this CLI request
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                  The address and port of the Kubernetes API server
      --tls-server-name string         Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use
```
```shell script
# 显示指定 Namespace 的 Deployment images
$ kubectl-img image -d -n hwy-demo
+-----------+------------+----------------+----------------+--------------+
| NAMESPACE |    TYPE    | RESOURCE_NAME  | CONTAINER_NAME |    IMAGE     |
+-----------+------------+----------------+----------------+--------------+
| hwy-demo  | Deployment |   hwy-nginx    |    nginx-1     | nginx:latest |
| hwy-demo  | Deployment | hwy-nginx-test |   nginx-test   | nginx:latest |
+-----------+------------+----------------+----------------+--------------+

# 显示指定 Namespace 的所有资源 images
$ kubectl-img image -dafcj -n hwy-demo
+-----------+------------+-------------------------+----------------+--------------+
| NAMESPACE |    TYPE    |      RESOURCE_NAME      | CONTAINER_NAME |    IMAGE     |
+-----------+------------+-------------------------+----------------+--------------+
| hwy-demo  | Deployment |        hwy-nginx        |    nginx-1     | nginx:latest |
| hwy-demo  | Deployment |     hwy-nginx-test      |   nginx-test   | nginx:latest |
| hwy-demo  |    Job     | cronjob-demo-1619797860 |     hello      |   busybox    |
| hwy-demo  |    Job     | cronjob-demo-1651911360 |     hello      |   busybox    |
| hwy-demo  |    Job     | cronjob-demo-1651911480 |     hello      |   busybox    |
| hwy-demo  |    Job     | cronjob-demo-1651911600 |     hello      |   busybox    |
| hwy-demo  |  CronJob   |      cronjob-demo       |     hello      |   busybox    |
+-----------+------------+-------------------------+----------------+--------------+

# 指定输出格式
$ kubectl-img image -d -n hwy-demo -o json
[
   {
      "CONTAINER_NAME": "nginx-1",
      "IMAGE": "nginx:latest",
      "NAMESPACE": "hwy-demo",
      "RESOURCE_NAME": "hwy-nginx",
      "TYPE": "Deployment"
   },
   {
      "CONTAINER_NAME": "nginx-test",
      "IMAGE": "nginx:latest",
      "NAMESPACE": "hwy-demo",
      "RESOURCE_NAME": "hwy-nginx-test",
      "TYPE": "Deployment"
   }
]
```