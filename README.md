# Gin Layout V2

这是一个开箱即用的 Go Web 服务器框架，基于 Gin 搭建，并整合了主流的开源组件。

## 特性

- **Web 框架**: Gin
- **配置管理**: Viper (YAML 格式)
- **依赖注入**: Wire
- **ORM & 代码生成**: Gorm + Gorm-Gen
- **日志**: 标准库 `slog` + Lumberjack (按大小自动切割，Error 日志独立)
- **API 文档**: Swagger
- **模块拆分**: 采用按层平铺拆分方式 (如 `handler/user.go`, `service/user.go`)

## 快速开始

### 1. 环境准备

确保已安装 Go 1.21 或更高版本。
安装必要的工具：

```bash
make init
```

### 2. 配置修改

修改 `configs/config.yaml`，配置 MySQL 连接字符串及其他参数。

### 3. 生成代码与运行

生成 Swagger 文档并注入 Wire 依赖，然后启动服务：

```bash
make run
```

访问 Swagger 文档：[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### 4. Gorm Gen 代码生成

```bash
make gen
```
