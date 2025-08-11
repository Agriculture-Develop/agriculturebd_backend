# Build stage
FROM golang:1.24.3-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# 复制依赖文件，利用缓存加速依赖下载
COPY go.mod go.sum ./
RUN go mod download

# 复制所有代码
COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main  ./api/main.go

# 运行阶段
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

# 复制编译好的二进制
COPY --from=builder /app/main /app/main

# 暴露端口（根据你的服务实际端口改）
EXPOSE 28080

# 不指定默认命令，docker-compose 里用 command 启动
