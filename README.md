# goeval

Evaluate Golang Code by the Eval Function

# Install

```shell script
$ go get github.com/PaulXu-cn/goeval
```

# Usage

```go
package main

import (
    "fmt"
    "github.com/PaulXu-cn/goeval"
)

func main() {
    if re, err := goeval.Eval("", "fmt.Print(\"Hello World!\")", "fmt"); nil == err {
        fmt.Println(string(re))
    } else {
        fmt.Println(err.Error())
    }
}
```

```
Hello World!
```

It's simple! Try it.