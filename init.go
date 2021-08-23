// author: wsfuyibing <websearch@163.com>
// date: 2021-08-15

// Elastic search manager.
package es

import (
	"sync"
)

var (
	once = new(sync.Once)
)

// 内存逃逸检测.
//
// 总结起来就是:
// 1. 如果在函数外部引用 -> 堆
// 2. 如果没有在外部引用 -> 栈(优先)
// 3. 如果函数返回的是一个局部变量的地址 -> 逃逸 (从栈 逃逸到 堆内存/heap)
//
// 避免逃逸的好处:
// 1. 减少gc的压力，不逃逸的对象分配在栈上，当函数返回时就回收了资源，不需要gc标记清除
// 2. 逃逸分析完后可以确定哪些变量可以分配在栈上，栈的分配比堆快，性能好(系统开销少)
// 3. 减少动态分配所造成的内存碎片
//
// go build -gcflags '-m'
// go build -gcflags '-m -l'
func init() {
	once.Do(func() {
		Config.init()
		Conn.init()
		Queue.init()
	})
}
