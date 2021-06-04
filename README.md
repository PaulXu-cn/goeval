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
    goeval "github.com/PaulXu-cn/goeval"
)

func main() {
    if re, err := goeval.Eval("", "fmt.Print(\"Hello World!\")", "fmt"); nil == err {
        fmt.Print(string(re))
    } else {
        fmt.Print(err.Error())
    }
}
```

```
Hello World!
```

It's sample !