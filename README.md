# fec-go
fec-go


### 安装

1.库包安装

1.1安装xorm

```
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm
```

2.安装gin

```
go get github.com/gin-gonic/gin
```

参考资料： [centos6 安装go框架gin的步骤，以及中间遇到的坑](http://www.fancyecommerce.com/2017/12/28/centos6-%e5%ae%89%e8%a3%85go%e6%a1%86%e6%9e%b6gin%e7%9a%84%e6%ad%a5%e9%aa%a4%ef%bc%8c%e4%bb%a5%e5%8f%8a%e4%b8%ad%e9%97%b4%e9%81%87%e5%88%b0%e7%9a%84%e5%9d%91/)

3.安装jwt-go

```
go get github.com/dgrijalva/jwt-go
```

```
// go get github.com/alecthomas/log4go


```

### 配置



1.将 config/config.ini 的内容，复制到 `/etc/fec-go/config.ini`
,上面的配置文件中配置相应的参数


2.main.go 设置绑定的ip和端口

```
listenIp := "120.24.37.249:3000"
```

将这个文件复制到main包下面，然后执行

```
go run main.go
```

启动go web























