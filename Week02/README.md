学习笔记

## 异常处理

> **核心问题**是如何把抛出的异常栈信息打印出来。

错误处理代码：

```go
package test

import (
	"golang.org/x/xerrors"
)

// ErrCustomMessage var.
var ErrCustomMessage = xerrors.New("err: Error custom messagae")

// Foo func.
func Foo() error {
	return ErrCustomMessage
}

// Bar func.
func Bar() error {
	return xerrors.Errorf("bar: %w", Foo())
}

// Try func.
func Try() error {
	return xerrors.Errorf("try : %w", Bar())
}
```

错误测试代码

```go
package test

import (
	"testing"
)

func TestError(t *testing.T) {
	err := Try()
	t.Errorf("%v", err)
	t.Errorf("%+v", err)
}
```

## 参考

* [Golang Error Handling — Best Practice in 2020](https://itnext.io/golang-error-handling-best-practice-a36f47b0b94c)
* [Golang 错误处理最佳实践](https://medium.com/@dche423/golang-error-handling-best-practice-cn-42982bd72672)
* [Go语言(golang)新发布的1.13中的Error Wrapping深度分析](https://www.flysnow.org/2019/09/06/go1.13-error-wrapping.html)
