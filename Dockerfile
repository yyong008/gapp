# 使用官方 Go 镜像作为构建环境
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制到容器中
COPY go.mod go.sum ./ 

# 下载依赖
RUN go mod tidy

# 将源代码复制到容器中
COPY . .

# 构建 Go 应用程序
RUN GOOS=linux GOARCH=amd64 go build -o myapp ./cmd/server/main.go

RUN echo "build complete /app"
RUN ls -al /app

# 使用更小的镜像运行应用程序
FROM ubuntu:latest


WORKDIR /app


# 从构建环境中复制编译好的二进制文件
COPY --from=builder /app/myapp /app/myapp
COPY --from=builder /app/config/config.yaml /app/config/config.yaml
COPY --from=builder /app/storage /app/storage


# 确认文件是否存在并列出目录
RUN ls -al /app

# 给文件设置可执行权限

RUN chmod +x /app/myapp

# 设置环境变量
ENV APP_ENV=production

# 暴露应用监听的端口
EXPOSE 3000

# 启动应用
CMD ["/app/myapp"]
