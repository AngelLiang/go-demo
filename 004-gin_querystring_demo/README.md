# 使用gin和swagger搭建认证接口示例

安装依赖包

```
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go mod download github.com/swaggo/files
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```


生成swag文档

```
swag init
```

启动服务

```
go run main.go
```

访问下面地址，即可访问接口文档

http://127.0.0.1:8080/swagger/index.html
