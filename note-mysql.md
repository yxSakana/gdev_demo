## MySQL字符集

- `GBK`: 长度2
- `UTF-8`: 3
- `latin1`: 1; 默认字符集
- `utf8mb4`: 4

四个颗粒度

- server全局: 仅影响最新的数据
- 数据库级: 仅影响最新的数据
- 表级: 会对历史数据的编码进行修改
- 列级: 会对历史数据的编码进行

## 约束条件

- 主键`primary key`
- 唯一`unique`
- 检查约束(对列数据的范围、格式限制) `cgecj`
- 默认值
- `not null`
- 外键约束`foreign key`

## 事务的四个特性

- 原子性: 通过`Redo log`和`Undo log`实现
- 一致性: 一个事务执行前后，数据从一个一致状态到另一个一致状态
- 隔离性: 通过`MVCC`实现
- 持久性: 一个事务一旦提交，其影响是永久的

## 刷新日志缓存

- 0: 每秒定时将日志缓冲区写入日志文件，并刷新日志到磁盘，事务提交时不做任何操作
- 1: 每次事务提交，将日志缓冲区写入文件，并刷新日志到磁盘(推荐)
- 2: 每次事务提交都将缓冲写入文件，但不刷新日志到磁盘，`InnoDB`按计划每秒刷新一次

## 复制步骤

- 源端将更改记录到`binlog`
- 副本将源端的`binlog`复制到自己的中继日志中
- 副本读取`binlog`并重放到副本上

## `binlog`

3个格式

- **statement：** 基于 SQL 语句的。某些语句和函数如 UUID, LOAD DATA INFILE 等在复制过程可能导致数据不一致甚至出错。
- **row：** 基于行的模式，记录的是行的变化。很安全，文件大。在一些大表中清除大量数据时在 binlog 中会生成很多条语句，可能导致从库延迟变大。
- **mixed：** 混合模式，根据语句来选用是 statement 还是 row 模式。

## 隔离级别

```sql
set session transaction isolation level xxx
```



- READ UNCOMMITTED: 存在脏读(在事务中可以查看其他事务中还没有被提交的读)
- READ COMMITTED: 可以看到其他事务在它开始之后**提交**的修改，但在该事务提交之前，其所做的修改对其他事务不可见。仍允许不可重复读
- REPEATABLE READ: 解决了不可重复读，但存在幻读，但MySQL通过MVCC解决了幻读。
- SERIALIZABLE: 不同事务之间不可能发生冲突。会在读取的每一行都加锁，所以可能会产生大量超时和锁竞争

概念

- 脏读: 当前事务中读取其他事务未提交的数据
- 幻读: 当前事务中读取某个范围内的记录，之后另一个事务在该范围内插入了新的记录，当前事物再读取时会出现幻行
- 不可重复读: 同一事务中执行相同语句，可能会有不同结果

幻读 与 不可重复读区别

不可重复度是无法处理同一行内的数据发生改变(`UPDATE`、`DELETE`)

幻读是无法处理新的数据行的改变(`INSERT`、`DELETE`)

## 事务日志

存储引擎只更改内存中的副本，然后将更改记录写入到事务日志中，由于写入事务日志时是追加写入，是顺序I/O，所以比直接更改磁盘中的数据的随机I/O要快

## MVCC原理

MVCC（多版本并发控制，Multi-Version Concurrency Control）核心概念

- 数据版本(数据快照): 每次写操作都会生成一个新的数据版本，旧版本不会立即删除，而是供并发事务读取
- 事务快照: 事务启动时的数据库快照，供读操作
- 版本号/时间戳: 每个事务都有一个递增的时间戳/版本号来标记事务顺序。每条数据都带有版本信息，表明该数据由哪个事务修改
- 隐藏字段: 有诸如`CreateVersion`、`DeleteVersion`来记录修改事务的版本号

实现机制

读操作: 利用多个数据版本和版本号来决定当前事物能够读取哪些数据

写操作: 插入时给新数据加上事务的版本信息

# 性能优化

## 索引

使用`B-tree`实现

自适应哈希索引：InnoDB中如果某些索引值被频繁访问，则会在索引的基础上再生成一个hash索引

索引对以下类型有效

- 全值匹配: 与索引中所有列匹配
- 匹配最左前缀: 只使用索引的第一列
- 匹配列前缀: 只匹配某一列的开头
- 匹配范围值: xx~yy范围之间
- 精准匹配某一列而范围匹配另一列
- 只访问索引的查询

选择性: `SELECT COUNT(field)/COUNT(*) FROM xxx`

使用`EXPLAIN`查看

```sql
EXPLAIN SELECT id FROM xxx
```

- `Extra`
  - `Using index`: 覆盖索引
- `Type`
  - index: 使用了索引扫描来排序

### 聚簇索引

一种数据存储方式，`InnoDB`中实际是在同一个结构中保存了B-tree和数据行。一个表只有一个聚簇索引。

特点：

- 数据行与B-tree在一起
  - 性能提升
  - 插入速度取决于插入顺序
  - 更新聚簇索引代价高
- 可能有页分裂
- 可能导致全表扫描变慢

### 索引维护

页分裂: 行的主键要求必须将这一行插入某个已满的页中时，存储引擎会将页分裂为两个

### 覆盖索引

一个索引包含所有需要查询的字段的值

### 重复索引

在相同列上按相同顺序创建相同的索引。(注意主键默认通过索引实现)

## 切分查询

MySQL的通信是半双工的

# 运维

## 监控数据

- 查询延迟
- 磁盘增长率
- 连接数增长
- 复制延迟
- 自增键
- 报错
- 使用峰值而不是平均值
- 使用百分数

