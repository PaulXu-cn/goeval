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

func main () {
	fmt.Print(string(goevel.Eval("", "fmt.Print(\"Hello World!\")")))
}
```

```
Hello World!
```