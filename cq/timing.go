package cq

import (
	"time"
)

type event func()

type timedEvent struct {
	name     string
	event    *event
	datetime string
	delayed  int64
}

type timerEvent struct {
	name    string
	timer   *time.Timer
	event   *event
	delayed int64
}

func (ch *CqH) RunTimer() {
	ch.runTimedEvent()
}

var timedEventList []*timedEvent

//延时执行事件
func (ch *CqH) DelayEvent(ds int, e event) {
	go func() {
		t := time.NewTimer(time.Second * time.Duration(ds))
		defer t.Stop()
		for {
			select {
			case <-t.C:
				e()
				return
			}
		}
	}()
}

//添加定时事件
func (ch *CqH) AddTimedEvent(name string, datetime string, e event, delayed int64) {
	timedEventList = append(timedEventList, &timedEvent{
		name:     name,
		event:    &e,
		datetime: datetime,
		delayed:  delayed,
	})
}

//删除定时事件
func (ch *CqH) DelTimedEvent(name string) {
	for i, v := range timedEventList {
		if (*v).name == name {
			timedEventList = append(timedEventList[:i-1], timedEventList[i+1:]...)
		}
	}
}

//定时事件执行
func (ch *CqH) runTimedEvent() {
	go func() {
		var ts = make([]*timerEvent, len(timedEventList))
		for i, v := range timedEventList {
			ts[i] = &timerEvent{
				name:    (*v).name,
				timer:   time.NewTimer(getNextTime((*v).datetime)),
				event:   (*v).event,
				delayed: (*v).delayed,
			}
			defer (*ts[i]).timer.Stop()
		}
		for i := 0; i < len(ts); i++ {
			select {
			case <-(*ts[i]).timer.C:
				(*(*ts[i]).event)()
				if (*ts[i]).delayed == 0 {
					ch.DelTimedEvent((*ts[i]).name)
				} else {
					(*ts[i]).timer.Reset(time.Duration((*ts[i]).delayed) * time.Second)
				}
			default:
				if i == len(ts)-1 {
					i = -1
					continue
				}
			}
			if i == len(ts)-1 {
				i = -1
			}
		}
	}()
}

//计算到期时间
func getNextTime(datetime string) time.Duration {
	loc, _ := time.LoadLocation("Asia/Chongqing")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
	if err != nil {
		return 0
	}
	res := t.Sub(time.Now())
	if res < 0 {
		return 0
	}
	return res
}
