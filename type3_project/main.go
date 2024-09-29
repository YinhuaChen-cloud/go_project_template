package main

import (
    "plugin"
    "log"
)

func main() {
    // 加载插件
    p, err := plugin.Open("print.so")
    if err != nil {
        log.Fatal(err)
    }

    // 查找 PrintHello 函数
    symbol, err := p.Lookup("PrintHello")
    if err != nil {
        log.Fatal(err)
    }

    // 将 symbol 转换为函数类型并调用
    printHello := symbol.(func())
    printHello()
}

