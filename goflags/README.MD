#### go-flags用于替代标准库flag的第三方库

安装
```
go get github.com/jessevdk/go-flags
```

import
```go
import "github.com/jessevdk/go-flags"
```

定义flag
```go
var opts struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Chinese bool   `short:"c" long:"chinese" description:"Use Chinese language"`
}
```
