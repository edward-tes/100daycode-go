# 第一天[2022-03-06]


## 1. Start
[官网](https://go.dev/doc)

### 创建一个文件夹

```fish
$ go mod init [projectName]
```

```
$ go mod tidy
```

```
$ go run ./
```

### 安装包

https://pkg.go.dev

```
$ go get [packageName]
```
## 2. 模块间的相互引用

###  给模块命名

example.com/hello

### 包映射到本地文件夹

```fish
$ go mod edit -replace example.com/greetings=../greetings
```

### 模块里面不需要main函数


## 3. 捕获异常
a. 引入`errors`
b. 函数可以多返回值，返回`nil`
c. 新建一个异常
```go
errors.New("error message")
```

## 4. 生成随机数

a. 数组的生成
b. init() 初始化方法，在go程序启动时会自动调用，不需要手动执行