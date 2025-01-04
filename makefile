# 项目变量
APP_NAME := myapp
GO_FILES := $(shell find . -name '*.go' -type f)
DOCKER_IMAGE := myapp-image
DOCKER_TAG := latest

# 环境变量
ENV_FILE := .env
CONFIG_FILE := config/config.yaml

# 默认目标
.PHONY: all
all: build

# 构建二进制文件
.PHONY: build
build:
	@echo "Building the application..."
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME) cmd/server/main.go

# 清理构建文件
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf bin/

# 运行项目
.PHONY: run
run:
	@echo "Running the application..."
	go run cmd/server/main.go --config=$(CONFIG_FILE)

# 测试
.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -coverprofile=coverage.out

# 查看测试覆盖率
.PHONY: coverage
coverage: test
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Open coverage.html to view the coverage report."

# 检查代码格式
.PHONY: lint
lint:
	@echo "Running linter..."
	golangci-lint run

# 格式化代码
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 数据库迁移
.PHONY: migrate
migrate:
	@echo "Applying database migrations..."
	migrate -path migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up

# 运行 Docker 镜像
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run --rm -it -p 8080:8080 --env-file $(ENV_FILE) $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: docker-clean
docker-clean:
	@echo "Cleaning Docker images..."
	docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG)

# 部署到生产环境
.PHONY: deploy
deploy:
	@echo "Deploying to production..."
	scp bin/$(APP_NAME) user@server:/path/to/deploy
	ssh user@server "cd /path/to/deploy && ./$(APP_NAME)"

# 帮助命令
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build         Build the application"
	@echo "  clean         Clean up build files"
	@echo "  run           Run the application locally"
	@echo "  test          Run tests with coverage"
	@echo "  coverage      Generate coverage report"
	@echo "  lint          Run linter"
	@echo "  fmt           Format code"
	@echo "  migrate       Apply database migrations"
	@echo "  docker-build  Build Docker image"
	@echo "  docker-run    Run Docker container"
	@echo "  docker-clean  Remove Docker image"
	@echo "  deploy        Deploy the application to production"
