FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app

# 接收架构参数
ARG TARGETARCH
COPY bin/${TARGETARCH}/main /app/main
