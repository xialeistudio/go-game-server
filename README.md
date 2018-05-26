# golang游戏服务器

## 服务器架构

```
玩家 -> LoginGate -> LoginServer -> 登录流程(返回Token)
玩家(携带Token) -> GameGate -> GameServer
```

## 服务器功能

+ LoginGate 完成登录注册的数据加解密
+ LoginServer 完成用户登录或注册
+ GameGate 数据加解密以及维护GameServer的连接