
Golang底层数据模型

值类型: int、float等基础类型、[3]int这样的数组、struct

引用类型: Slice、map、interface

slice: 实际可理解为一个结构体, 其包括指向底层数组的指针、切片的大小、切片的容量

map: 指向底层哈希表的指针

interface: 动态类型, 底层存储 类型信息、值信息; 所以不要使用指向`interface`的指针

```go
type myInterface interface {
	aFunc()
}

func test(v *myInterface) {} // bad
func test(v myInterface) {} // good
```

## 小知识点

### `for range`

```go
// v的内存地址在每次循环时始终不变
// 每次循环时，将值覆盖到该内存地址
for v := range arr {
    v := v
    go func() {
        // use v
    }
}
```



## slice

`slice`的结构类似于:

```go
type Slice struct {
    Data unsafe.Pointer  // 指向底层数组的指针
    Len  int             // 数组当前大小
    Cap  int			 // 数组容量
}
```

所以

- `slice`是一种引用类型，传入函数时(golang函数为值传递)，仍会携带指向底层的指针，所以会修改数据
- 从数组中取`slice`成本很低

## map

`map`是一个哈希表的**引用**

所以

- 为引用类型，赋值/传参时，注意该操作可能会修改底层数据

### 哈希表的实现与哈希冲突

#### 实现

哈希表的一般结构由:

- 数组(或桶): 存储`k-v`的固定大小的数组
- 哈希函数: 根据`key`计算出在数组中的索引位置
- 冲突解决策略: 处理多个键映射到同一索引时，如何解决

哈希函数具有

- 均匀性: 键应被均匀映射
- 确定性: 相同输入总产生相同输出
- 高效性: 计算哈希的时间复杂度应为O(1)

负载因子（Load Factor，即已存储键值对数量与数组容量之比）

#### 哈希冲突

- 开放地址法: 在发生冲突时，将冲突键放在数组的空闲位置
  - 线性探测: 冲突发生时，顺序向后查找下一个空闲位置
  - 二次探测: 使用非线性步长（如平方步长）来寻找空闲位置，减少聚集效应。
  - 双重哈希: 使用第二个哈希函数来确定冲突时的探测步长
- 链地址法: 哈希表的桶中存放的是链表，将对应的`k-v`存储在链表中，发送冲突就往该链表往后延伸
- 重哈希: 当哈希表的负载因子超过一定阈值时: 创建一个更大的数组并重新计算所有键的哈希值，将键值对重新分布到新的数组中(Rehashing)。

### key

`map`的`key`可以是任何**可以被比较的类型**(包括: `int` `float` `string` `bool` 不包含不可比较的类型的`struct`和数组)

### 初始化

```go
var (
    m1 = make(map[T1]T2) // 初始化 - 读写安全
	m2 map[T1]T2         // 声明 - 写入时会 panic
)

// 使用初始化列表
m := map[T1]T2{
  k1: v1,
  k2: v2,
  k3: v3,
}
```

### 有序遍历

```go
type OrderMap struct {
    OrderKey []string
    Data     map[string]int
}

om := OrderMap{...}
for k range om.OrderKey {
    fmt.Println(om.Data[k])
}
```

### 内存释放

删除一个`key`后，会经历:

- 将`k-v`从`map`中移除
- 因为没有被任何变量引用，被GC标记
- GC回收相关内存

### sync.map

提供并发安全的map

关键原理

- 读写分离: 有两个`map`: 只读字典、读写字典。
  - 读: 优先访问只读字典，如果没有再访问读写字典
  - 写: 延迟写入。立即更新读写字典。只读字典的更新发生在: 读操作时，发现只读字典中数据超过`misses`计数器阈值
- 延迟写入
- 原子操作: 读操作大部分情况下无锁；写操作会加锁保护读写字典
- 条目淘汰

## struct

### 初始化

```go
// 应
// 明确指定字段名
// 省略零值
k := User{
    FirstName: "John",
    LastName: "Doe",
    Admin: true,
}

// 省略所有字段
var user User

// struct 指针
val := T{Name: "foo"}
ptr := &T{Name: "bar"}
```

## 错误处理

....

## 函数

...

## 接口

接口的底层结构存储了: 动态类型、动态值

```go
type iface struct {
    tab *itab           // 动态类型
    data unsafe.Pointer // 动态值
}

type itab struct {
    inter *interfacetype // 指向 描述接口类型的方法集的结构体
    typ   *typeInfo      // 实现该接口的具体类型
    hash  uint32         // (类型和接口的)哈希值(用于快速查找)
    fun   [1]uintptr     // 方法表(存储实现该接口的方法的指针)
}
```

## 反射

### 原理

...

### 示例

...

## goroutine

...

## channel

底层主要有四部分

- 存放数据的循环链表
- 记录当前发送/接受数据的下标
- 等待队列
- 锁

### 写入时

- 无缓冲: 写操作会阻塞，直到有另一个goroutine读取channel中的数据，相当于做一次同步
- 有缓冲: 缓冲未满，不会阻塞；缓冲满时，会阻塞直到有空间
- channel被关闭: 触发`panic`
- 为nil: 永远阻塞

### 读取时

- 无缓冲: 读操作会阻塞，直到有另一个`goroutine`向`channel`中写数据，相当于做一次同步
- 有缓冲: 缓冲中有数据，直接读取，不会阻塞；没有数据，会阻塞直到有数据
- `channel`被关闭: 可读取`channel`中剩余的数据；如果`channel`中已经没有数据则返回对应类型的零值
- 为nil: 永远阻塞

## select

```go
select {
case ch <- 3:
    fmt.Println("写入成功")
case <-time.After(2 * time.Second):
    fmt.Println("写入超时，放弃写入")
}
```

- 一个：有某个`case`满足时，执行该`case`
- 多个：当多个 case 同时满足条件时，`select` 会随机选择一个 case 执行。

- 0个：如果所有 case 都不满足条件且没有 `default` 分支，`select` 会阻塞，直到有 case 满足条件。

- 0个且有`default`：如果有 `default` 分支且所有其他 case 都不满足条件，`select` 会直接执行 `default` 分支而不会阻塞。

## mutex

锁的分类

- 悲观锁: 假设数据竞争频繁发生。所有对共享资源的访问都会加锁。
- 乐观锁: 假设数据竞争很少发送。不会在操作前加锁，只会在提交时判断数据是否有被其他线程修改，没有则提交，有则失败。多基于`CAS`
- 读写锁: ...
- 自旋锁: 线程尝试获取锁时不会立即阻塞，而是在一段时间内反复尝试。避免线程切换的开销，但是有反复尝试的开销，仅用在每次持有锁的时间很短时

`sync.Mutex`的两个模式

- 正常模式
  - `mutex`只有一个`goroutine`竞争，直接获取
  - 多个`goroutine`竞争
    - 如果`mutex`已经被获取，则新的`goroutine`被加入到`waiter`队列(FIFO)且处于自旋状态
    - 如果`mutex`处于空闲状态，则新的`goroutine`参与竞争，被唤醒的`waiter`也参与竞争，但由于新`goroutine`正在CPU中运行，能更快获取`mutex`
- 饥饿模式(`waiter`队列在1ms内获取不到锁)
  - `mutex`的持有者直接将锁交给`waiter`，新的`goroutine`直接加入到`waiter`，同时不会保持自旋
  - 如果`waiter`中取出的已经是最后一个了 或 `waiter`等待时间小于1ms

## context

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    
    Done() <-chan struct{}  // context未关闭: 返回nil; 关闭: 返回被关闭的管道(仍可读)
    
    Err() error // 返回关闭的原因
    
    Value(Key interface{}) interface{}  // 传递值
}
```

类型

- `context.Background()`: 空context一般用于其他context的父节点
- `context.WithCancel()`: 关闭自己以及子context。一般用来关闭协程
- `context.WithDeadline`(): 定时关闭....
- `context.WithTimeout()`: 在超时时关闭....
- `context.WithValue()`: 一般在协程之间传递上下文中的值/变量

## GMP

- G: `goroutine`协程, 包含具体任务

- M: `thread`线程, 运行`goroutine`包含的任务

- P: `processor`处理器, 含有存放`goroutine`的本地队列

- 全局队列: 存放`goroutine`的全局队列

一个G(`goroutine`)被创建时，先优先加入到P的本地队列，如果P中的本地队列满了，则将本地队列中一半的G移动到全局队列中；

M(`thread`)要运行任务时，首先关联某一个未被关联的P，并从中获取一个G(`goroutine`)，如果P中没有可被调度的G，则尝试从全局队列中获取或从其他

P的本地队列中偷取一半放到自己的P的本地队列中；

当M(`thread`)执行G被阻塞时，M与P的关联被取消，将P让渡给其他空闲的M；

如果没有足够的M来关联P并执行G，那么就会创建新的M。

- M0: 主线程，负责初始化操作和启动第一个G，之后与其他M相同
- G0: 每次启动M时创建的第一个`goroutine`，仅负责调度G

有几个需要注意的

- 线程复用: 线程被创建后，会被重复使用，不断被调度执行任务，不会存在大量的线程创建、销毁造成的资源损耗
- 抢占: 一个`goroutine`最多占用CPU 10ms(而不是等待`goroutine`主动让出CPU)，防止其他`goroutine`被饿死

**debug 工具**: `go tool trace main.out`

## GC

