# Redis

![redis八股文提纲](/pic/redis八股文提纲.webp)
[toc]

## 一、Redis数据类型和数据结构

![Redis数据类型和数据结构](/pic/redis数据结构.jpg)

参考：<https://xiaolincoding.com/redis/data_struct/data_struct.html>

### 1.1 ziplist

**zipist数据结构如下**：
![ziplist](/pic/ziplist.jpg)

压缩列表节点包括三部分：

- prevlen
记录了「前一个节点」的长度，目的是为了实现从后向前遍历；
- encoding
记录了当前节点实际数据的「类型和长度」，类型主要有两种：字符串和整数；
- data
记录了当前节点的实际数据，类型和长度都由 encoding 决定；

1. ziplistNode保存prevlen的目的
为了实现节点从后往前遍历，后面listpack保存len也是为了
相同目的。

### 1.2 quicklist

quicklist其实就是【双向链表+压缩列表】的组合，因为quicklist本身是个列表，而链表中的每个元素又是一个压缩列表。

1. 压缩列表的弊端
会出现【连锁更新】的风险
2. quicklist解决办法
通过控制每个链表节点中压缩列表的大小或元素个数，来规避
连锁更新带来的影响。因为压缩列表元素越少或越小，连锁更新带来的影响就越小，从而提供了更好的访问性能。

**quicklist数据结构如下**：
![quicklist](/pic/quicklist.jpg)

### 1.3 listpack

quicklist减少了【连锁更新】造成的影响，但是并没有完全解决这一问题，这源于ziplist的底层设计结构。要想彻底解决这个问题，需要设计一个新的数据结构。redis7.0已经将所有使用zipist数据结构的对象全部替换为listpack。

**listpack数据结构如下**：
![listpack](/pic/listpack.jpg)

listpack节点包括三部分：

1. encoding
记录了当前节点实际数据的「类型和长度」，类型主要有两种：字符串和整数；
2. data
记录了当前节点的实际数据，类型和长度都由 encoding 决定；
3. len
encoding+data的总长度；

<font color='red'>相比于ziplist，listpack不再记录前一个加点长度字段，只记录当前节点长度。当我们向 listpack 加入一个新元素的时候，不会影响其他节点的长度字段的变化，从而避免了压缩列表的连锁更新问题。</font>

### 1.4 zskiplist

**quicklist（跳表）数据结构如下**：
![zskiplist跳表](/pic/zskiplist.jpg)

数据量很大时，跳表的查找复杂度为O(logN)

1. 跳表节点层数设置
  跳表在创建节点的时候，随机生成每个节点的层数（层高最大为：redis5.0位64，redis7.0为32）。

## 二、Redis数据持久化

### 2.1 redis数据持久化方式

#### 2.1.1 AOF日志

每执行一条`写操作`命令，就把该命令以追加的方式写入到一个文件里。重启 Redis 的时候，先去读取这个文件里的命令，并且执行它，就会恢复缓存数据。
`AOF日志是在主线程中执行的`

AOF写回策略如下：
![AOF写回策略](/pic/AOF写回策略.jpg)

AOF日志过大会触发`AOF重写机制`，通过重写机制，删除被覆盖的旧命令，如：

```shell
set name xk
set name xkchen
```

上述两个命令在重写时就会变成`set name xkchen`。这样一来，一个键值对在重写日志中只用一条命令就行了。

#### 2.1.2 RDB快照

将某一时刻的内存数据，以`二进制`的方式写入磁盘；redis的快照是`全量快照`，每次执行快照，是将内存中所有数据都记录到磁盘中。

**RDB快照缺点：**
当服务器发生故障时，丢失的数据会比AOF更多。因为是全量快照，所以频率不能太高，否则会影响redis性能，而AOF可以秒级的方式记录数据。

Redis 提供了两个命令来生成 RDB 文件，分别是 save 和 bgsave，他们的区别就在于是否在「主线程」里执行：

- 执行了 save 命令，就会在主线程生成 RDB 文件，由于和执行操作命令在同一个线程，所以如果写入 RDB 文件的时间太长，会阻塞主线程；
- 执行了 bgsave 命令，会创建一个子进程来生成 RDB 文件，这样可以避免主线程的阻塞；

#### 2.1.3. 混合持久化

开启方式：

```shell
aof-use-rdb-preamble yes
```

集成AOF和RDB的优点；

混合持久化工作在 <font color='red'>AOF 日志重写过程</font>

使用混合持久化，AOF文件的前半部分是RDB格式的全量数据，后半部分是AOF的增量命令数据。这样做的好处：

1. 由于前半部分是 RDB 内容，这样加载的时候速度会很快
2. 加载完 RDB 的内容后，才会加载后半部分的 AOF 内容

![redis混合持久化](/pic/redis混合持久化.jpg)

### 2.2 如何避免大key

1. 大key带来的影响

   - 客户端阻塞：由于redis单线程，在操作大key会比较耗时，客户端需要等待很久
   - 服务端阻塞：使用del命令，会阻塞工作线程
   - 引发网络阻塞：如果单个Key大小1M，每秒来1000次访问，那么就会产生1000M流量，对于普通千兆网卡服务器是灾难性的
   - 内存分布不均：集群模型在 slot 分片均匀情况下，会出现数据和查询倾斜情况，部分有大 key 的 Redis 节点占用内存多，QPS 也会比较大

2. 如何避免大key
在设计阶段，就把大key拆分成小key。或者定期检查redis是否存在大key，如果可以删除，使用unlink代替del命令，防止阻塞主线程。

## 三、Redis集群

## 四、

## 五、
