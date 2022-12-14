# go-cq介绍

## 项目介绍
go-cq是一个基于cq-http的快速开发框架,它不需要开发人员与各种复杂的接口打交道,熟悉golang原生http开发的能够快速上手使用它

## 接口文档


## 前置条件
本框架依赖cq-http,请先开启cq-http服务
cq-http下载位置 https://github.com/Mrs4s/go-cqhttp

## 快速上手
`go get -u github.com/jjheiy/go-cq`

```golang
app := cq.New()
//监听群聊消息
app.HandleMsgGroupFunc("你好", func(ai *cq.Ai) {
    //监听到你好时,发送你好啊
    ai.Group.SendMsg("你好啊")
})
//开启监听
app.Run(":5701")    //监听cq-http端口
```
## 获取参数
参数以{n:参数名,t:参数类型,c:cqcode类型,r:正则表达式}的形式,除了参数名,其他都可以省略,当参数名为第一个参数时,n可以省略
```golang
app.HandleMsgGroupFunc("你好啊,{name}", func(ai *cq.Ai) {
    name := ai.Msg.GetParameter("name")
    ai.Group.SendMsg("我不是" + name + "啊,你认错人了吧!")
})
```
## 参数限制
```golang
app.HandleMsgGroupFunc("我电话{phone,r:\\d{11}}", func(ai *cq.Ai) {
    phone := ai.Msg.GetParameter("phone")
    ai.Group.SendMsg("你的电话是" + phone)
})
app.HandleMsgGroupFunc("我今年{age,t:uint8}岁了", func(ai *cq.Ai) {
    age := ai.Msg.GetParameter("age")
    ai.Group.SendMsg("你的年龄是" + age)
})
```
**注意: 如果参数有正则表达式限制,则r必须为最后一个参数,并且当参数不是结尾时,需要使用非贪婪匹配**

## 管理员系统
go-cq内置了管理员系统,包含超级管理员,不同权限的管理员,白名单,黑名单系统
```golang
app.HandleMsgAdmin("禁言{qq,c:id} {time,t:int}小时", func(ai *cq.Ai) {
    qq := ai.Msg.GetParameter("qq")
    id := ai.Code.GetAtQQ(qq)   //获取真实qq号
    times := ai.Msg.GetParameter("time")
    time, _ := strconv.Atoi(times)
    ai.Group.TabooUser(id, time*60) //禁言操作
    ai.Group.SendMsg(qq + "已被禁言")
}, 8) //8为该操作最小管理员权限值(大于等于8的都会触发

app.HandleMsgAdmin("发公告{content}", func(ai *cq.Ai) {
    content := ai.Msg.GetParameter("content")
    ai.Group.SendNotice(content) //发公告操作
}, 10)

app.HandleMsgAdmin("修改权限{qq,c:id} {power,t:uint8}", func(ai *cq.Ai) {
    qq := ai.Msg.GetParameter("qq")
    id := ai.Code.GetAtQQ(qq)
    powers := ai.Msg.GetParameter("power")
    power, _ := strconv.Atoi(powers)
    ai.AddAdmins(id, uint8(power))
    ai.Group.SendMsg("已将" + qq + "的权限修改为" + powers)
}, 8)
//将qq号为2357054981添加为管理员,权限为9,此时只能禁言,无法发公告
app.AddAdmins(2357054981, 9)
```

## 实现一个简单的学习功能
```golang
app.HandleMsgAdmin("学习 {key},{value}", func(ai *cq.Ai) {
    key := ai.Msg.GetParameter("key")
    value := ai.Msg.GetParameter("value")
    ai.HandleMsgGroupFunc(key, func(ai *cq.Ai) {
        ai.Group.SendMsg(value)
    })
    ai.Group.SendMsg("学习成功\r\n问:" + key + "\r\n答:" + value)
}, 5)
```
由于参数解析是存放在内存中,我们不建议动态去创建解析事件
go-cq中内置了一个shit组件,我们可以使用它实现学习功能
## shits组件实现学习功能
shit组件只会在消息没有进行任何处理时才会执行
```golang
mp := map[string]string{ //模拟静态数据 
    "你好":      "我不好",
    "为什么不好":   "就是不好",
    "为什么就是不好": "gun",
}
app.Shits(func(ai *cq.Ai) {
    value := ai.Msg.GetValue(mp)    //从数据源里匹配
    if len(value) == 0 {
        return
    }
    ai.Group.SendMsg(value)
    ai.Stop()
})
```
**注意** `ai.stop()`会直接跳出当前会话
## 前置器
前置器是在参数解析之前执行的函数,常用做参数效验,快速操作,中间件
```golang
app.Uses(func(ai *cq.Ai) {
    if strings.Contains(ai.Msg.GetRawMessage(), "sb") {
        err := ai.Msg.DelMsg()//测回消息
        if err != nil {
            panic(err)
        }
        ai.Group.SendMsg("检测到不好的词汇,已撤回")
        ai.Stop() //结束当前会话
    }
})
```
## 后置器
后置器是在参数解析之后执行的函数,常用做数据统计,信息收集,日志处理
```golang
app.Afters(func(ai *cq.Ai) {
    mp := map[string]any{
        "id":  ai.User.GetId(),
        "age": ai.User.GetAge(),
    }
    fmt.Println(mp)
})
```
## 定时器
1. 延时执行

在一段时间后执行
```golang
app.HandleMsgGroupFunc("{num}分钟后提醒我喝水", func(ai *cq.Ai) {
    num, err := strconv.Atoi(ai.Msg.GetParameter("num"))
    if err != nil {
        ai.Group.SendMsg("输入格式有误")
        return
    }
    ai.DelayEvent(num*60, func() {
        at := ai.Code.At(strconv.FormatInt(ai.User.GetId(), 10))
        ai.Group.SendMsg(at + " 多喝热水")
    })
})
```

2. 定时执行

指定某个时间执行,可设置执行间隔时间,为0表示只执行一次
```golang
app.HandleMsgGroupFunc("开启宵禁{time1},{time2}", func(ai *cq.Ai) {
    time1 := ai.Msg.GetParameter("time1") //"2022-10-08 23:00:00"
    time2 := ai.Msg.GetParameter("time2") //"2022-10-09 7:00:00"
    ai.AddTimedEvent("curfew", time1, func() {
        ai.Group.SendMsg("宵禁时间到了,大家晚安!!!")
        ai.Group.TabooGroup(true)
    }, 24*60*60)
    ai.AddTimedEvent("liftABan", time2, func() {
        ai.Group.TabooGroup(false)
        ai.Group.SendMsg("新的一天开始了,大家早上好!!!")
    }, 24*60*60)
})
```
