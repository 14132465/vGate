# 第一阶段：构建阶段（使用大而全的编译镜像）
FROM golang:1.21-alpine AS builder

# 设置代理（可选，方便国内下载依赖）
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 1. 先复制依赖文件，利用 Docker 缓存（避免每次都要重新下载）
COPY go.mod go.sum ./
RUN go mod download

# 2. 复制源码并编译
COPY . .
# 编译：-ldflags 去掉调试信息，-s -w 可减小体积
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o vgate ./cmd/gate

# 第二阶段：运行阶段（使用极简的空白镜像）
FROM alpine

# 设置时区（可选）
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#ENV TZ=Asia/Shanghai

# 将编译好的二进制文件从 builder 复制过来
COPY --from=builder /app/vgate /application/cmd/gate/vgate

# 如果项目需要配置文件，也一并复制
COPY --from=builder /app/config/config.example.yml /application/config/config.vgate.yaml

# 暴露端口（根据你的项目修改，例如 8080）
EXPOSE 8080 6789

# 运行
ENTRYPOINT ["/application/cmd/gate/vgate"]