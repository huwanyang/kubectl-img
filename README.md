[![ci-release](https://github.com/huwanyang/kubectl-img/actions/workflows/ci-release.yml/badge.svg?branch=master)](https://github.com/huwanyang/kubectl-img/actions/workflows/ci-release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/huwanyang/kubectl-img)](https://goreportcard.com/report/github.com/huwanyang/kubectl-img)

kubectl-img 是基于 Cobra 脚手架创建的命令插件，可以显示指定的 k8s 资源类型（deployments|daemonsets|statefulsets|jobs|cronjobs）
的 image 信息，并支持多种输出方式（json|yaml|xml|table）。

## Usage

```$xslt
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