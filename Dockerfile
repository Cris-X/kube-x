# 使用最新的 Go 版本镜像，例如 1.22-alpine
FROM golang:1.22-alpine as builder

# 设置工作目录
WORKDIR /go/src/cris-x.com/server

# 复制 go.mod 和 go.sum 文件以便先行下载依赖
COPY go.mod go.sum ./

# 设置 Go 环境变量，使用国内代理
RUN go env -w GO111MODULE=on \
   && go env -w GOPROXY=https://goproxy.cn,direct \
   && go mod download

# 复制项目所有文件
COPY . .

# 设置构建相关的环境变量，并编译项目
RUN go env -w CGO_ENABLED=0 \
   && go mod tidy \
   && go build -o server .

# 使用轻量级的 Alpine 镜像作为运行环境
FROM alpine:latest

# 设置镜像的维护者信息
LABEL MAINTAINER="cris-x@github.com"

# 设置运行时的工作目录
WORKDIR /go/src/cris-x.com/server

# 从构建阶段复制必要的文件到运行阶段
COPY --from=builder /go/src/cris-x.com/server/config.yaml ./config.yaml
COPY --from=builder /go/src/cris-x.com/server/.kube/config ./.kube/config
COPY --from=builder /go/src/cris-x.com/server/server ./

# 暴露服务端口
EXPOSE 8081

# 设置容器启动时的入口命令
ENTRYPOINT ["./server"]