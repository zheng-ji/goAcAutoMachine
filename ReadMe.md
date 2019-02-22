## goAcAutoMachine

[![Build Status](https://travis-ci.org/zheng-ji/goAcAutoMachine.svg)](https://travis-ci.org/zheng-ji/goAcAutoMachine)
[![GoDoc](https://godoc.org/github.com/zheng-ji/goAcAutoMachine?status.svg)](https://godoc.org/github.com/zheng-ji/goAcAutoMachine)


Go 实现多模式字符串匹配的 AC 自动机

### Install

```
go get "github.com/zheng-ji/goAcAutoMachine"
```

### Example

```Go
package main

import (
    "fmt"
    "github.com/zheng-ji/goAcAutoMachine"
)

func main() {
    ac := goAcAutoMachine.NewAcAutoMachine()
    ac.AddPattern("红领巾")
    ac.AddPattern("祖国")
    ac.AddPattern("花朵")
    ac.Build()

    content := "我是红领巾，祖国未来的花朵"
    results := ac.QUery(content)
    for _, result := range results {
        fmt.Println(result)
    }
}
```

### License

Copyright (c) 2019 by [zheng-ji](http://zheng-ji.info) released under MIT License.
