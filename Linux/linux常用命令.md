# linux常用命令

[toc]

## 一、systemctl

用途：管理系统服务

### 1.1 查看所有启动的命服务

```bash
systemctl list-units --type=service --state=running
```

## 二、netstat

用途：管理系统服务

### 2.1 查看所有开启的端口

```bash
netstat -nltp
```
