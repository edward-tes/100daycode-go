# 第一天[2022-03-06]


## 1. 安装
[官网](https://go.dev/doc)下载安装

### 初始化项目

```fish
$ go mod init [projectName]
```

验证go.sum文件，在程序中使用了的包作为依赖包，安装依赖包
```
$ go mod tidy
```

执行应用
```
$ go run ./
```

### 添加新的包

https://pkg.go.dev

```
$ go get [packageName]
```
## 2. 模块间的相互引用

a. 模块命名规范

example.com/hello

b. 包映射到本地文件夹

```fish
$ go mod edit -replace example.com/greetings=../greetings
```

c. 模块里面不需要main函数

## 3. 捕获异常

1. 引入`errors`模块

1. 函数可以返回多个值，返回`nil`

1. 新建一个异常

```go
errors.New("error message")
```

## 4. 生成随机数

1. 引入 `math/rand`
1. init() 初始化方法，在go程序启动时会自动调用，不需要手动执行

## https://github.com/edward-tes/100daycode-go