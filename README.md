# ChatOyO即时聊天服务

## 编译执行

```shell
go build -o server main.go server.go user.go
./server
```

## 应用层协议

### 1. who

获取当前在线用户列表

无参数

### 2. rename

修改当前用户的名字

`` rename|[NEW_NAME] ``

### 3. private

It's a private conversation. (致敬新概念英语2第一课)

`` to|[UID/USER_NAME]|[MESSAGE]``

## 迭代

### 0.0 Genesis - 基本私聊公聊功能

- 0.0.1. 构建基础Server
- 0.0.2. 用户上线功能
- 0.0.3. 用户消息广播机制
- 0.0.4. 用户业务层封装
- 0.0.5. 在线用户查询
- 0.0.6. 修改用户名
- 0.0.7. 超时踢出功能
- 0.0.8. 私聊功能
- 0.0.8-1 用于测试的客户端(WIP) 连接
- 0.0.8-2 用于测试的客户端(WIP) 命令行解析
- 0.0.8-3 用于测试的客户端(WIP) 菜单显示
- 0.0.8-4 用于测试的客户端(WIP) 用户名更新
- 0.0.8-5 用于测试的客户端(WIP) 公聊和私聊模式

Ref. [案例-即时通信系统](https://www.yuque.com/aceld/mo95lb/ks1lr9)

### 0.1 Evolve - 更完善的玩具系统
- 0.1.1 代码优化和架构调整
- 0.1.2 使用Websocket协议