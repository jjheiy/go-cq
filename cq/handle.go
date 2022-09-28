package cq

import "fmt"

type CqH struct {
}

func (ch *CqH) Use(handler ...handle) {
	roximitor = append(roximitor, handler...)
}
func (ch *CqH) After(handler ...handle) {
	postprocessor = append(postprocessor, handler...)
}

func (ch *CqH) Uses(handler ...handle) {
	roximitors = append(roximitors, handler...)
}
func (ch *CqH) Afters(handler ...handle) {
	postprocessors = append(postprocessors, handler...)
}

func (ch *CqH) Shit(handler ...handle) {
	shit = append(shit, handler...)
}
func (ch *CqH) Shits(handler ...handle) {
	shits = append(shits, handler...)
}

//处理所有消息
func (ch *CqH) HandleMsgFunc(pattern string, handler func(*Ai)) {
	handlers = append(handlers, newHandleMsg(pattern, handler))
}

//处理所有群聊消息
func (ch *CqH) HandleMsgGroupFunc(pattern string, handler func(*Ai)) {
	handlerGroups = append(handlerGroups, newHandleMsg(pattern, handler))
}

//处理所有私聊消息
func (ch *CqH) HandleMsgPrivateFunc(pattern string, handler func(*Ai)) {
	handlerPrivates = append(handlerPrivates, newHandleMsg(pattern, handler))
}

//处理所有临时会话消息
func (ch *CqH) HandleMsgTemporaryFunc(pattern string, handler func(*Ai)) {
	handlerTemporarys = append(handlerTemporarys, newHandleMsg(pattern, handler))
}

//处理超级管理员事件
func (ch *CqH) HandleMsgAdminSuper(pattern string, handler func(*Ai)) {
	if len(adminSupers) == 0 {
		return
	}
	ch.HandleMsgBySectionFunc(pattern, handler, "admin6", adminSupers...)
}

//处理管理员事件
func (ch *CqH) HandleMsgAdmin(pattern string, handler func(*Ai), power uint8) {
	ads := make([]int64, 0)
	for k, v := range admins {
		if v >= power {
			ads = append(ads, k)
		}
	}
	mp := make(powerHandle)
	mp[newHandleMsg(pattern, handler)] = power
	temp := &mp
	handleByPower[temp] = ads
}

//处理白名单事件
func (ch *CqH) HandleMsgWhitelist(pattern string, handler func(*Ai)) {
	if len(whitelists) == 0 {
		return
	}
	ch.HandleMsgBySectionFunc(pattern, handler, "whitelist", whitelists...)
}

//处理黑名单事件
func (ch *CqH) HandleMsgBlacklist(pattern string, handler func(*Ai)) {
	if len(blacklists) == 0 {
		return
	}
	ch.HandleMsgBySectionFunc(pattern, handler, "blacklist", blacklists...)
}

//处理指定id事件
func (ch *CqH) HandleMsgBySectionFunc(pattern string, handler func(*Ai), tag string, ids ...int64) {
	mp := make(map[string][]int64)
	mp[tag] = ids
	handlerBySection[newHandleMsg(pattern, handler)] = mp
}

//处理所有消息
func (ch *Ai) HandleMsgFunc(pattern string, handler func(*Ai)) {
	hm := newHandleMsg(pattern, handler)
	if hm.Parsing() {
		fmt.Println("新增all消息", hm.pattern)
	}
	handlers = append(handlers, hm)
}

//处理所有群聊消息
func (ch *Ai) HandleMsgGroupFunc(pattern string, handler func(*Ai)) {
	hm := newHandleMsg(pattern, handler)
	if hm.Parsing() {
		fmt.Println("新增群聊消息", hm.pattern)
	}
	handlerGroups = append(handlerGroups, hm)
}

//处理所有私聊消息
func (ch *Ai) HandleMsgPrivateFunc(pattern string, handler func(*Ai)) {
	hm := newHandleMsg(pattern, handler)
	if hm.Parsing() {
		fmt.Println("新增私聊消息", hm.pattern)
	}
	handlerPrivates = append(handlerPrivates, hm)
}

//处理所有临时会话消息
func (ch *Ai) HandleMsgTemporaryFunc(pattern string, handler func(*Ai)) {
	hm := newHandleMsg(pattern, handler)
	if hm.Parsing() {
		fmt.Println("新增临时会话消息", hm.pattern)
	}
	handlerTemporarys = append(handlerTemporarys, hm)
}

//处理超级管理员事件
func (ch *Ai) HandleMsgAdminSuper(pattern string, handler func(*Ai)) {
	if len(adminSupers) == 0 {
		return
	}
	ch.HandleMsgBySectionFunc(pattern, handler, "admin6", adminSupers...)
}

//处理管理员事件
func (ch *Ai) HandleMsgAdmin(pattern string, handler func(*Ai), power uint8) {
	ads := make([]int64, 0)
	for k, v := range admins {
		if v >= power {
			ads = append(ads, k)
		}
	}
	hm := newHandleMsg(pattern, handler)
	if hm.Parsing() {
		fmt.Println("新增管理员消息", hm.pattern)
	}
	mp := make(powerHandle)
	mp[hm] = power
	temp := &mp
	handleByPower[temp] = ads
}

//处理白名单事件
func (ch *Ai) HandleMsgWhitelist(pattern string, handler func(*Ai)) {
	if len(whitelists) == 0 {
		return
	}
	ch.HandleMsgBySectionFunc(pattern, handler, "whitelist", whitelists...)
}

//处理黑名单事件
func (ch *Ai) HandleMsgBlacklist(pattern string, handler func(*Ai)) {
	if len(blacklists) == 0 {
		return
	}
	ch.HandleMsgBySectionFunc(pattern, handler, "blacklist", blacklists...)
}

//处理指定id事件
func (ch *Ai) HandleMsgBySectionFunc(pattern string, handler func(*Ai), tag string, ids ...int64) {
	mp := make(map[string][]int64)
	mp[tag] = ids
	hm := newHandleMsg(pattern, handler)
	if hm.Parsing() {
		fmt.Println("新增特殊消息", hm.pattern)
	}
	handlerBySection[hm] = mp
}
