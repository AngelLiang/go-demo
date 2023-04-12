# gin响应json

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


参考资料：https://golang.halfiisland.com/%E8%AF%AD%E8%A8%80%E8%BF%9B%E9%98%B6/%E5%AE%9E%E7%94%A8%E6%A1%86%E6%9E%B6/gin.html#%E6%95%B0%E6%8D%AE%E8%A7%A3%E6%9E%90
