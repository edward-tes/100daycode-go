# 第二天[2022-03-07]


## 返回多个输出值

1. 声明map

```go
 messages := make(map[string]string)
```

2. 遍历map

```go
for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
```

## 添加单元测试

以`_test.go`规则结尾的文件

引入模块

```go
import (
    "testing"
    "regexp"
)
// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
```

## 编译安装文件

构建

```shell
$ go build
```

设置安装目录到shell path
```
$ export PATH=$PATH:/path/to/your/install/directory
```


或者设置GOBIN变量

```shell
 go env -w GOBIN=/path/to/your/bin
```


```shell
$ go install
```

之后便可以在任意地方执行命令(全局命令)

```
hello
```
