# golang游戏服务器

## 服务器架构

```
玩家 -> LoginGate -> LoginServer -> 登录流程(返回Token)
玩家(携带Token) -> GameGate -> WorldServer -> 选择大区
```