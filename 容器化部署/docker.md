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

### 1.7 查看latest tag对应镜像的具体版本号

```bash
# 以redis为例
docker image inspect redis:latest | grep -i version
```

### 1.8 下载特定版本的镜像

```bash
# 拉取redis历史版本
docker pull redis:6.2.5
```

### 1.9 删除容器、镜像

```bash
# 删除容器
docker rm $(docker ps -a -q) # 删除所有容器
docker rm -f '{container_id}' # 删除特定容器

# 删除镜像
docker rmi -f redis:latest
```

## 二、docker启动一个redis服务

```bash
# 启动一个简单的redis服务（这里需要严格注意后面-d -p的顺序）
docker run --name myredis -d -p 16379:6379 redis:6.2.5

# 启动一个redis服务，开启aof（rdb默认开启）
docker run \
--restart=always \
--log-opt max-size=100m \
--log-opt max-file=2 \
-p 16379:6379 \
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

## 三、docker面试题

### 3.1 探探对docker的理解

docker是一个轻量级的虚拟机，也是去实现容器化技术的一个应用工具。它是基于linux内核的cgroup（资源限制）、namespace（资源隔离）、unionFS等技术对进程进行封装和隔离，属于操作系统层面的虚拟化技术。独立的进程隔离与宿主和其他隔离的进程，所以也称为`容器`。（虚拟机是硬件层面上的虚拟化技术）
容器不用单独启动操作系统，而虚拟机需要单独启动操作系统。

### 3.1 docker compose的作用

Docker Compose 是一个用于定义和运行多容器 Docker 应用程序的工具。它使用 YAML 文件来配置应用程序服务、网络和卷

### 3.2 Docker中卷的使用

用作持久化容器数据。卷可以挂载到容器系统中的一个目录，容器被删除后，卷及上面的数据依然存在

```shell
# 1. 创建卷，名字为myvol
docker volumn create myvol

# 2. 列出本地 Docker 主机上的全部卷 
docker volumn ls 

# 创建了一个新的容器，并将容器内的 /vol 目录挂载到了名为 bizvol 的卷。假如容器的文件系统中没有 /vol 这个目录，那么会创建；假如已有这个目录，那么则会使用这个目录（该目录的内容到时候会变成卷里面的内容）。同理，系统中没有叫 bizvol 的卷，那么该命令也会创建一个这样的卷；如果已经存在这个卷了，那么则使用这个卷。 
docker container run -it --name voltainer --mount source=bizvol,target=/vol alpine
```

### 3.3 Dockerfile 中ENTRYPOINT指令的用途

ENTRYPOINT指定容器启动时将执行的命令。常用于设置容器的主进程。例如：
ENTRYPOINT ["java", "-jar", "myapp.jar"]

### 3.4 Dockerfile 中的CMD指令的用途是什么

CMD指令设置容器运行时默认执行的命令。它可以在容器启动期间被覆盖。例如：
CMD ["java", "-jar", "myapp.jar"]

### 3.5 如何从 Dockerfile 构建 Docker 镜像

使用docker build命令。例如：docker build -t my_image

### 3.6 解释Docker层的概念

 "层"的概念是通过联合文件系统AuFS来实现的。使用层可以让我们在扩展镜像时会更方便。
举个例子：目录1包含文件a，文件c；目录2包含文件b，文件c。通过联合文件方式，将目录1和目录2挂在在目录3上，此时，目录3拥有文件a、b、c三个文件。如果在目录3中对文件a、b、c进行修改，对应的目录1和目录2也会生效。
