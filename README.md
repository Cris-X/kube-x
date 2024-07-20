# kube-x

Kube-x is a simple tool to manage your kubernetes cluster. 

Dev-Env:
- go-version: 1.22.5
- compile-env: MAC
- dev-tool: goland

## Installation

### web-framework
```bash
go get -u github.com/gin-gonic/gin
```
> docs: https://gin-gonic.com/docs

### config-management
```bash
go get -u github.com/spf13/viper
```
> docs: https://github.com/spf13/viper

### k8s-client
```bash
go get -u k8s.io/client-go@v0.30.2
```
> docs: https://pkg.go.dev/k8s.io/client-go