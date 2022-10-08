package cq

import (
	"fmt"
	request "github.com/jjheiy/go-cq/cq/request"
	"regexp"
	"strconv"
	"strings"
)

type cqCode struct{}

//获取qq号
func (cc *cqCode) GetAtQQ(cqStr string) int64 {
	cqStr = strings.TrimSpace(cqStr)
	if len(cqStr) == 0 {
		return 0
	}
	if cqStr[:1] != "[" {
		id, _ := strconv.ParseInt(cqStr, 10, 64)
		return id
	}
	cqs := strings.Split(cqStr, ",")
	if strings.Split(cqs[0], ":")[1] == "!at" {
		return 0
	}
	for _, cq := range cqs {
		if strings.Contains(cq, "qq=") {
			id, _ := strconv.ParseInt(cq[3:len(cq)-1], 10, 64)
			return id
		}
	}
	return 0
}

//获取cqcode字典
func (cc *cqCode) GetCqCodeMap(cqStr string) map[string]any {
	var res = make(map[string]any)
	cqStr = strings.TrimSpace(cqStr)
	if cqStr[:1] != "[" || cqStr[len(cqStr)-1:] != "]" {
		return nil
	}
	cqStr = cqStr[1:] + cqStr[:len(cqStr)-2]
	cqs := strings.Split(cqStr, ",")
	if len(cqs) == 0 {
		return nil
	}
	i := 0
	for {
		s := "="
		if i == 0 {
			s = ":"
		}
		temp := strings.Split(cqs[i], s)
		if len(temp) == 2 {
			res[temp[0]] = temp[1]
		} else if len(temp) == 1 {
			res[temp[0]] = ""
		} else {
			return nil
		}
		if i == len(cqs)-1 {
			break
		}
		i++
	}
	return res
}

//获取字符串中所有的cqcode
func (cc *cqCode) GetCqCodeStr(msg string) []string {
	cqCodeRegexp := regexp.MustCompile(`\[CQ:\w+(?:,\w*=[^\]]*)*]`)
	res := cqCodeRegexp.FindAllString(msg, -1)
	return res
}

//at
func (cc *cqCode) At(id string) string {
	return "[CQ:at,qq=" + id + "]"
}

//分享链接
func (cc *cqCode) Share(url string, title string) string {
	return "[CQ:share,url=" + url + ",title=" + title + "]"
}

//音乐
func (cc *cqCode) Music(name string) string {
	dataGet, err := request.Get("https://music.cyrilstudio.top/search?keywords=" + name + "&limit=1")
	if err != nil {
		return ""
	}
	data := dataGet["result"].(map[string]interface{})
	if int(data["songCount"].(float64)) == 0 {
		return "未找到该歌曲"
	}
	id := fmt.Sprintf("%d", int((data["songs"].([]interface{})[0].(map[string]interface{})["id"].(float64))))
	return "[CQ:music,type=163,id=" + id + "]"
}

//文字以语音发送
func (cc *cqCode) Tts(msg string) string {
	return "[CQ:tts,text=" + msg + "]"
}

//表情
func (cc *cqCode) Face(id string) string {
	return "[CQ:face,id=" + id + "]"
}

//猜拳
func (cc *cqCode) Rps() string {
	return "[CQ:rps]"
}

//骰子
func (cc *cqCode) Dice() string {
	return "[CQ:dice]"
}

//图片
func (cc *cqCode) Img(url string) string {
	return "[CQ:image,file=" + url + "]"
}

//回复
func (cc *cqCode) Reply(id int32) string {
	return "[CQ:reply,id=" + strconv.Itoa(int(id)) + "]"
}

//红包
func (cc *cqCode) RedBag(title string) string {
	return "[CQ:redbag,title=" + title + "]"
}

//戳一戳
func (cc *cqCode) Poke(qq int64) string {
	return "[CQ:poke,qq=" + strconv.Itoa(int(qq)) + "]"
}

//礼物
func (cc *cqCode) Gift(qq int64, id string) string {
	return "[CQ:gift,qq=" + strconv.Itoa(int(qq)) + ",id=" + id + "]"
}
