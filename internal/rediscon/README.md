

## 旁路缓存

![旁路缓存](../../docs/cache-旁路缓存.svg)

伪代码

```
func Write(data) {
    mysql.update(data)
    cache.del(data)
}

func Read(key) {
    var result
    
    if cache.get(key) {
        return result
    }
    
    result = mysql.get(key)
    cache.set(key, result)
    return result
}
```

读写穿透

![读写穿透](../../docs/cache-读写穿透.svg)

异步缓存写入
