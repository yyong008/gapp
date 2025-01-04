# gapp

gapp 是一个 go gin 的单体 web 项目项目。

## 开始

```sh
./cmd/server/main.go
```

## TODO

- [x]config 配置管理
- [x]app 抽象
- [x]优雅停机
- [x]中间件
  - [x]cors
  - [x]auth
- [x]数据库支持
- [x]简单迁移
- [x]jwt 封装和配置
- [x]swagger
- [x]makefile
- [x]dockerfile
- [x]数据库支持
- []zap log支持
## swagger

- `http://localhost:8080/swagger/index.html`

```sh
swag init -g ./cmd/server/main.go
```
