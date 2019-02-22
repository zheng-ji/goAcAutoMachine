
## goAcAutoMachine

Go 实现多模式字符串配置的 AC 自动机

### Install

```
go get "github.com/zheng-ji/goAcAutomachine"
```

### Example

```Go
package main

import (
    "fmt"
    "github.com/zheng-ji/goAcAutomachine"
)

func main() {
    content := "我是红领巾，祖国未来的花朵"
    ac := goAcAutoMachine.NewAcAutoMachine()
    ac.AddPattern("红领巾")
    ac.AddPattern("祖国")
    ac.AddPattern("花朵")
    ac.Build()
    results := ac.Scan(content)
    fmt.Println("内容: " + content)
    for _, result := range results {
		fmt.Println(result)
    }
}
```

### License

Copyright (c) 2019 by [zheng-ji](http://zheng-ji.info) released under MIT License.
