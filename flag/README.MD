#### 使用标准库flag来定义命令行参数

flag有2种定义方法

1.
```go
var name = flag.String("name", "World", "A name to say hello to.")
```
适用于一对一的命令参数，返回的是指针，所以后续要用*name来读取参数值

2.
```go
var chinese bool
flag.BoolVar(&chinese, "chinese", false, "Use Chinese language.")
flag.BoolVar(&chinese, "c", false, "Use Chinese language.")
```
适用于一对多的命令参数，参数值会直接保存定义的变量里，所以后续用chinese来读取参数值

#### flag.PrintDefaults()可以打印参数的说明
```
-c    Use Chinese language.
-chinese
      Use Chinese language.
-name string
      A name to say hello to. (default "World")
```

#### init()不需要显式调用，会在main()执行前自动执行
