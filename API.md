# go-cq api文档

## cq

- New()

获取一个cq启动对象,后续所有的操作都由该对象执行
```go
app := cq.New()
```
### CqH(app)

- Use(handler ...handle)

私聊前置器,在接收到私聊消息之前处理的事件

- Uses(handler ...handle)

群聊前置器,在接收到群聊消息之前处理的事件

- After(handler ...handle)

私聊后置器,在处理完私聊消息后执行的事件

- Afters(handler ...handle)

群聊后置器,在处理完群聊消息后执行的事件

- Shit(handler ...handle)

私聊shit组件,在没有处理任何私聊消息时执行的事件

- Shits(handler ...handle)

群聊shit组件,在没有处理任何群聊消息时执行的事件

- HandleMsgFunc(pattern string,handler handle)

1. pattern  匹配参数
2. handler  执行事件

所有消息都会处理的事件(优先级最低)

- HandleMsgTemporaryFunc(pattern string,handler handle)

1. pattern  匹配参数
2. handler  执行事件

添加处理临时会话的事件

- HandleMsgGroupFunc(pattern string,handler handle)

1. pattern  匹配参数
2. handler  执行事件

添加群聊消息处理的事件

- HandleMsgPrivateFunc(pattern string,handler handle)

1. pattern  匹配参数
2. handler  执行事件

添加私聊消息处理的事件

- HandleMsgAdmin(pattern string,handler handle,power uint8)

1. pattern  匹配参数
2. handler  执行事件
3. power    权限值

添加管理员消息处理的事件(最高优先级)

- HandleMsgAdminSuper(pattern string,handler handle)

1. pattern  匹配参数
2. handler  执行事件

添加超级管理员消息处理的事件

- HandleMsgWhitelist(pattern string,handler handle)

1. pattern  匹配参数
2. handler  执行事件

添加白名单消息处理的事件

- HandleMsgBlacklist(pattern string,handler handle)

1. pattern  匹配参数
2. handler  执行事件

添加黑名单消息处理的事件

- app.HandleMsgBySectionFunc(pattern string,handler handle,tag string,ids ...int64)

1. pattern  匹配参数
2. handler  执行事件
3. tag   事件标签
4. ids    对于事件执行的id

添加自定义id消息处理的事件

- AddAdmins(id int64, power uint8)

1. id 
2. power  权限值

添加管理员

- app.DelAdmins(ids ...int64)

删除管理员

- AddBlacklists(ids ...int64)

添加黑名单
- DelBlacklists(ids ...int64)

删除黑名单
- AddWhitelists(ids ...int64)

添加白名单
- DelWhitelists(ids ...int64)

删除白名单
- AddAdminSupers(ids ...int64)

添加超级管理员
- DelAdminSupers(ids ...int64)

删除超级管理员
- AddTimedEvent(name string, datetime string, e event, delayed int64)

1. name    定时器名字
2. datetime     执行的时间 格式(2006-01-02 15:04:05)
3. e     执行的事件
4. delayed      间隔时间(0表示只执行一次)

添加定时器任务
- DelTimedEvent(name string)

1. name  定时器名字

删除定时器任务
- RunTimer()

开始执行定时器任务

- DelayEvent(ds int, e event)

1. ds 延时时间
2. 事件

延时执行事件

### Ai