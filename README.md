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

1. 构建基础Server
2. 用户上线功能
3. 用户消息广播机制
4. 用户业务层封装
5. 在线用户查询
6. 修改用户名
7. 超时踢出功能
8. 私聊功能
9. 客户端实现（？仅作学习和测试，因为已经做了个B端的）