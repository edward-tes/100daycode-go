# 第三天[2022-03-08] 

知识点： 多模块工作空间 
目的：在一个多模块工作空间，创建两个模块。

## 需求

+ GO1.18以上，只能安装rc版本

初始化工作空间

```shell
$ go work init ./hello
```

```shell
$ go work use ./example
```go work use ./example