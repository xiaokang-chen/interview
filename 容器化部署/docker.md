# docker

[toc]

## 一、基础命令

### 1.1 查看镜像列表

```bash
docker images
```

### 1.2 查看所有运行中的container

```bash
# 列出运行中的container
docker ps
# 列出所有container
docker ps -a
```

### 1.3 停止container

```bash
docker stop [container_id]
```

### 1.4 删除停止的container

```bash
doskcer rm [container_id]
```

### 1.5 启动一个container

```bash
# 启动一个交互式服务
docker run -it ubuntu /bin/bash
# 启动一个后台运行服务
docker run -d redis:latest redis-server 
```

### 1.6 列出所有images

```bash
docker images
```

## 二、docker启动一个redis服务

```bash
# 启动一个redis服务，开启aof（rdb默认开启）
docker run \
--restart=always \
--log-opt max-size=100m \
--log-opt max-file=2 \
-p 6379:6379 \
--name myredis \
-v /home/redis/myredis/myredis.conf:/etc/redis/redis.conf \
-v /home/redis/myredis/data:/data \
-d redis:latest redis-server /etc/redis/redis.conf \
--appendonly yes \
--requirepass 123456

# 进入redis容器
docker exec -it myredis /bin/bash

# 通过redis-check-rdb查看rdb文件（二进制文件，无法直接查看）
redis-check-rdb /data/appendonlydir/appendonly.aof.1.base.rdb

# 启动客户端连接
redis-cli
```
