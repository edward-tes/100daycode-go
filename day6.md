# 第6天[2022-03-11]

模糊测试(fuzzing)

## 准备

1. 安装 go1.18beta
2. 写单元测试
3. 添加模糊测试

执行模糊测试

```sh
$ go test -fuzz=Fuzz
```

4. 执行测试

```sh
go test -fuzz=Fuzz -fuzztime 30s
```