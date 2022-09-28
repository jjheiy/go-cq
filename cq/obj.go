package cq

import (
	"fmt"
	"hyai/cq/request"
	"regexp"
	"strings"
)

type Ai struct {
	CqH
	Code  *cqCode
	User  *User
	Group *Group
	Msg   *Msg
	state byte
}

func newAi(u *User, g *Group, m *Msg) *Ai {
	return &Ai{
		Code:  &cqCode{},
		User:  u,
		Group: g,
		Msg:   m,
		state: 0,
	}
}

func (ai *Ai) Stop() {
	ai.state = 4
}

//用户信息
type User struct {
	id       int64
	age      int32
	nickname string
	sex      string
	card     string //群名片
	area     string //地区
	level    string //等级
	role     string //角色 owner,admin,member
	title    string //专属头衔
}

func newUser(id int64, name string, sex string, age int32) *User {
	return &User{
		id:       id,
		nickname: name,
		age:      age,
		sex:      sex,
	}
}
func newUserByGroup(id int64, name string, sex string, age int32, card string, area string, level string, role string, title string) *User {
	return &User{
		id:       id,
		nickname: name,
		age:      age,
		sex:      sex,
		card:     card,
		area:     area,
		level:    level,
		role:     role,
		title:    title,
	}
}
func (u *User) GetId() int64 {
	return (*u).id
}
func (u *User) GetNickname() string {
	return (*u).nickname
}
func (u *User) GetAge() int32 {
	return (*u).age
}
func (u *User) GetSex() string {
	return (*u).sex
}
func (u *User) GetCard() string {
	return (*u).card
}
func (u *User) GetArea() string {
	return (*u).area
}
func (u *User) GetlLevel() string {
	return (*u).level
}
func (u *User) GetRole() string {
	return (*u).role
}
func (u *User) GetTitle() string {
	return (*u).title
}

//群聊信息
type Group struct {
	id int64
}

func newGroup(id int64) *Group {
	return &Group{
		id: id,
	}
}

func (g *Group) GetId() int64 {
	return (*g).id
}

//消息信息
type Msg struct {
	id         int32
	parameter  map[string]string
	msgType    string                      //消息类型 private/group
	subType    string                      //消息子类型 friend,group,group_self,other/normal,anonymous,notice
	message    string                      //消息内容
	rawMessage string                      //原始消息内容
	font       int32                       //字体
	cqCodes    map[string][]map[string]any //cqCode集
	time       int64                       //发送时间
	tempSource int                         //临时会话来源
}

func newMsg(id int32, msgType string, subType string,
	message string, rawMessage string, font int32, time int64) *Msg {
	return &Msg{
		id:         id,
		msgType:    msgType,
		subType:    subType,
		message:    message,
		rawMessage: rawMessage,
		font:       font,
		time:       time,
	}
}
func (m *Msg) setParameter(parameter map[string]string) {
	(*m).parameter = parameter
}
func (m *Msg) GetId() int32 {
	return (*m).id
}
func (m *Msg) GetParameter(name string) string {
	return (*m).parameter[name]
}

func (m *Msg) GetMsgType() string {
	return (*m).msgType
}
func (m *Msg) GetSubType() string {
	return (*m).subType
}
func (m *Msg) GetMessage() string {
	return (*m).message
}
func (m *Msg) GetRawMessage() string {
	return (*m).rawMessage
}
func (m *Msg) GetFont() int32 {
	return (*m).font
}
func (m *Msg) GetTime() int64 {
	return (*m).time
}
func (m *Msg) GetTempSource() int {
	return (*m).tempSource
}

func (m *Msg) GetValue(mp map[string]string) string {
	return mp[m.rawMessage]
}
func (m *Msg) GetIndex(strs []string) int {
	for i, s := range strs {
		if s == m.rawMessage {
			return i
		}
	}
	return -1
}

//发送私聊消息
func (u *User) SendMsg(msg string) int32 {
	url := getUrl() + "/send_private_msg"

	data, err := request.Post(url, newUserMsg(u.GetId(), msg, false).jsonToBytes())
	if err != nil {
		return 0
	}
	data = data["data"].(map[string]any)
	return int32(data["message_id"].(float64))
}

//发送临时会话消息
func (u *User) SendMsgByGroup(msg string, groupId int64) int32 {
	url := getUrl() + "/send_private_msg"

	data, err := request.Post(url, newUserMsgByGroup(u.GetId(), msg, groupId, false).jsonToBytes())
	if err != nil {
		return 0
	}
	data = data["data"].(map[string]any)
	return int32(data["message_id"].(float64))
}

//上传私聊文件
func (u *User) UpFile(file string, name string) error {
	url := getUrl() + "/send_private_msg"

	_, err := request.Post(url, newUpUserFile(u.GetId(), file, name).jsonToBytes())
	return err
}

//撤回消息
func (m *Msg) DelMsg() error {
	url := getUrl() + "/delete_msg"
	_, err := request.Post(url, newIdMsg(m.GetId()).jsonToBytes())
	return err
}

//获取消息
func (m *Msg) GetMsg() map[string]any {
	url := getUrl() + "/get_msg"
	data, err := request.Post(url, newIdMsg(m.GetId()).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//获取消息图片
func (m *Msg) GetMsgImage() map[string]any {
	url := getUrl() + "/get_image"
	data, err := request.Post(url, newImgMsg((*m).cqCodes["img"][0]["url"].(string)).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//发送群聊消息
func (g *Group) SendMsg(msg string) int32 {
	url := getUrl() + "/send_group_msg"
	data, err := request.Post(url, newGloupMsg(g.GetId(), msg, false).jsonToBytes())
	if err != nil {
		return 0
	}
	data = data["data"].(map[string]any)
	return int32(data["message_id"].(float64))
}

//群聊踢人
func (g *Group) DelGroupUser(id int64, rejectAdd ...bool) error {
	url := getUrl() + "/set_group_kick"
	var reject bool
	if len(rejectAdd) == 0 {
		reject = false
	} else {
		reject = rejectAdd[0]
	}
	_, err := request.Post(url, newGroupKick(g.GetId(), id, reject).jsonToBytes())
	return err
}

//群聊禁言
func (g *Group) TabooUser(id int64, duration int) error {
	url := getUrl() + "/set_group_ban"
	_, err := request.Post(url, newGroupBan(g.GetId(), id, duration*60).jsonToBytes())
	return err
}

//群聊全员禁言
func (g *Group) TabooGroup(enable bool) error {
	url := getUrl() + "/set_group_whole_ban"
	_, err := request.Post(url, newGroupAllBan(g.GetId(), enable).jsonToBytes())
	return err
}

//群聊设置管理员
func (g *Group) SetGroupAdmin(userId int64, enable bool) error {
	url := getUrl() + "/set_group_admin"
	_, err := request.Post(url, newGroupSetAdmin(g.GetId(), userId, enable).jsonToBytes())
	return err
}

//群聊设置名片
func (g *Group) SetGroupCard(userId int64, card string) error {
	url := getUrl() + "/set_group_card"
	_, err := request.Post(url, newGroupSetCard(g.GetId(), userId, card).jsonToBytes())
	return err
}

//群聊设置群名
func (g *Group) SetGroupName(name string) error {
	url := getUrl() + "/set_group_name"
	_, err := request.Post(url, newGroupSetName(g.GetId(), name).jsonToBytes())
	return err
}

//退出群聊
func (g *Group) LeaveGroup() error {
	url := getUrl() + "/set_group_leave"
	_, err := request.Post(url, newGroupLeave(g.GetId(), false).jsonToBytes())
	return err
}

//解散群聊
func (g *Group) DissolutionGroup() error {
	url := getUrl() + "/set_group_leave"
	_, err := request.Post(url, newGroupLeave(g.GetId(), true).jsonToBytes())
	return err
}

//群聊设置头衔
func (g *Group) SetGroupTitle(userId int64, title string) error {
	url := getUrl() + "/set_group_special_title"
	_, err := request.Post(url, newGroupTitle(g.GetId(), userId, title, -1).jsonToBytes())
	return err
}

//获取群成员信息
func (g *Group) GetGroupUserInfo(userId int64, noCache bool) request.Json {
	url := getUrl() + "/get_group_member_info"
	data, err := request.Post(url, newGroupUserInfo(g.GetId(), userId, noCache).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//获取群成员信息列表
func (g *Group) GetGroupUserInfoList() request.List {
	url := getUrl() + "/get_group_member_list"
	data, err := request.Post(url, newGroupIdp(g.GetId()).jsonToBytes())
	if err != nil {
		return nil
	}
	return data["list"].(request.List)
}

//获取群荣誉信息
func (g *Group) GetGrouphonorInfo(hType string) request.Json {
	url := getUrl() + "/get_group_honor_info"
	data, err := request.Post(url, newGroupHonorInfo(g.GetId(), hType).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//设置群头像
func (g *Group) SetGroupImg(file string, cache int) request.Json {
	url := getUrl() + "/set_group_portrait"
	data, err := request.Post(url, newSetGroupImg(g.GetId(), file, cache).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//上传群聊文件
func (g *Group) UpFile(file string, name string, folder ...string) error {
	url := getUrl() + "/send_private_msg"
	var filep = newUpGroupFile(g.GetId(), file, name)
	if len(folder) != 0 {
		(*filep).Folder = folder[0]
	}
	_, err := request.Post(url, filep.jsonToBytes())
	return err
}

//获取群聊文件信息
func (g *Group) GetFileInfo() request.Json {
	url := getUrl() + "/get_group_file_system_info"
	data, err := request.Post(url, newGroupIdp(g.GetId()).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//获取群聊根目录文件列表
func (g *Group) GetFileListRoot() request.Json {
	url := getUrl() + "/get_group_root_files"
	data, err := request.Post(url, newGroupIdp(g.GetId()).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//获取群聊子目录文件列表
func (g *Group) GetFileList(folder string) request.Json {
	url := getUrl() + "/get_group_files_by_folder"
	data, err := request.Post(url, newGetGroupFileList(g.GetId(), folder).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//创建群聊文件目录
func (g *Group) CreateFolder(name string) error {
	url := getUrl() + "/create_group_file_folder"
	_, err := request.Post(url, newCreateFolder(g.GetId(), name).jsonToBytes())
	return err
}

//删除群聊文件目录
func (g *Group) DelFolder(folderId string) error {
	url := getUrl() + "/delete_group_folder"
	_, err := request.Post(url, newGetGroupFileList(g.GetId(), folderId).jsonToBytes())
	return err
}

//删除群文件
func (g *Group) DelFile(fileId string, busid int32) error {
	url := getUrl() + "/delete_group_file"
	_, err := request.Post(url, newGroupFile(g.GetId(), fileId, busid).jsonToBytes())
	return err
}

//获取群文件
func (g *Group) GetFile(fileId string, busid int32) request.Json {
	url := getUrl() + "/get_group_file_url"
	data, err := request.Post(url, newGroupFile(g.GetId(), fileId, busid).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//获取群文件
func (g *Group) GetAtCount() request.Json {
	url := getUrl() + "/get_group_at_all_remain"
	data, err := request.Post(url, newGroupIdp(g.GetId()).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//发送群公告
func (g *Group) SendNotice(content string, img ...string) error {
	url := getUrl() + "/_send_group_notice"
	sgnp := newSendGroupNotice(g.GetId(), content)
	if len(img) != 0 {
		(*sgnp).Image = img[0]
	}
	_, err := request.Post(url, sgnp.jsonToBytes())
	return err
}

//获取群公告
func (g *Group) GetNotice() request.List {
	url := getUrl() + "/_get_group_notice"
	data, err := request.Post(url, newGroupIdp(g.GetId()).jsonToBytes())
	if err != nil {
		return nil
	}
	return data["list"].(request.List)
}

//获取群消息历史记录
func (g *Group) GetMsgHistory(messageSeq int64) request.Json {
	url := getUrl() + "/get_group_msg_history"
	data, err := request.Post(url, newGetGroupMsgHistory(g.GetId(), messageSeq).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//获取群精华消息列表
func (g *Group) GetEssenceList() request.List {
	url := getUrl() + "/get_essence_msg_list"
	data, err := request.Post(url, newGroupIdp(g.GetId()).jsonToBytes())
	if err != nil {
		return nil
	}
	return data["list"].(request.List)
}

//检查链接安全性
func (ai *Ai) GetUrlSafely(urls string) int {
	url := getUrl() + "/check_url_safely"
	data, err := request.Post(url, newUrlStr(urls).jsonToBytes())
	if err != nil {
		return 0
	}
	return int(data["level"].(float64))
}

//设置精华消息
func (ai *Ai) SetEssence(id int32) error {
	url := getUrl() + "/set_essence_msg"
	_, err := request.Post(url, newMsgIdp(id).jsonToBytes())
	return err
}

//移出精华消息
func (ai *Ai) DelEssence(id int32) error {
	url := getUrl() + "/delete_essence_msg"
	_, err := request.Post(url, newMsgIdp(id).jsonToBytes())
	return err
}

//设置登陆号资料
func (ai *Ai) SetAiInfo(nickname, company, email, college, personalNote string) error {
	url := getUrl() + "/set_qq_profile"
	_, err := request.Post(url, newSetAiInfo(nickname, company, email, college, personalNote).jsonToBytes())
	return err
}

//查询陌生人信息
func (ai *Ai) GetUserInfo(id int64, noCache bool) request.Json {
	url := getUrl() + "/get_stranger_info"
	data, err := request.Post(url, newStrangerInfo(id, noCache).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//获取群信息
func (ai *Ai) GetGroupInfo(id int64, noCache bool) request.Json {
	url := getUrl() + "/get_group_info"
	data, err := request.Post(url, newGroupInfo(id, noCache).jsonToBytes())
	if err != nil {
		return nil
	}
	return data
}

//获取好友列表
func (ai *Ai) GetFriendList() request.List {
	url := getUrl() + "/get_friend_list"
	data, err := request.Post(url, []byte(""))
	if err != nil {
		return nil
	}
	return data["list"].(request.List)
}

//获取群列表
func (ai *Ai) GetGroupList() request.List {
	url := getUrl() + "/get_group_list"
	data, err := request.Post(url, []byte(""))
	if err != nil {
		return nil
	}
	return data["list"].(request.List)
}

//获取群系统消息
func (ai *Ai) GetGroupSystemMsg() request.Json {
	url := getUrl() + "/get_group_system_msg"
	data, err := request.Post(url, []byte(""))
	if err != nil {
		return nil
	}
	return data
}

//获取单项好友列表
func (ai *Ai) GetUnidirectionalFriendList() request.List {
	url := getUrl() + "/get_unidirectional_friend_list"
	data, err := request.Post(url, []byte(""))
	if err != nil {
		return nil
	}
	return data["list"].(request.List)
}

//删除好友
func (ai *Ai) DelUser(id int64) error {
	url := getUrl() + "/delete_friend"
	_, err := request.Post(url, newUserIdp(id).jsonToBytes())
	return err
}

//群聊打卡
func (ai *Ai) GroupSign(id int64) error {
	url := getUrl() + "/send_group_sign"
	_, err := request.Post(url, newGroupIdp(id).jsonToBytes())
	return err
}

//获取状态
func (ai *Ai) GetStatus() request.Json {
	url := getUrl() + "/get_status"
	data, err := request.Post(url, []byte(""))
	if err != nil {
		return nil
	}
	return data
}

//Handle function type
type handle func(*Ai)

type handleMsg struct {
	pattern string
	handler handle
	p_names []string
	p_count int
}

func (hm *handleMsg) Match(msg string) bool {
	reg := regexp.MustCompile(hm.pattern)
	return reg.MatchString(msg)
}

func (hm *handleMsg) GetParamet(msg string) map[string]string {
	mp := make(map[string]string)
	reg := regexp.MustCompile(hm.pattern)
	ps := reg.FindStringSubmatch(msg)
	if len(ps) < 2 {
		return mp
	}
	ps = ps[1:]
	if len(ps) != hm.p_count {
		return nil
	}
	for i := 0; i < hm.p_count; i++ {
		mp[hm.p_names[i]] = ps[i]
	}
	return mp
}

func (hm *handleMsg) Parsing() bool {
	regexStr := ""
	patterns := getPatternOrder(hm.pattern)
	if patterns == nil {
		return false
	}
	p_names := make([]string, 0)
	p_count := 0
	for i := 0; i < len(patterns); i++ {
		if pt := analysisPattern(patterns[i]); pt != nil {

			p_count++
			p_names = append(p_names, pt.GetPName())
			if pt.GetPRegex() == "" {
				if i != len(patterns)-1 {
					regexStr += "(.*?)"
				} else {
					regexStr += "(.*)"
				}
			} else {
				regexStr += "(" + pt.pRegex + ")"
			}
		} else {
			regexStr += patterns[i]
		}
	}
	hm.pattern = "^" + regexStr + "$"
	hm.p_names = p_names
	hm.p_count = p_count
	return true
}

func (hm *handleMsg) run(ai *Ai) {
	parameters := hm.GetParamet(ai.Msg.GetRawMessage())
	if parameters == nil {
		return
	}
	ai.Msg.setParameter(parameters)
	(*hm).handler(ai)
}

func newHandleMsg(pattern string, handler handle) *handleMsg {
	return &handleMsg{
		pattern: pattern,
		handler: handler,
	}
}

var roximitors []handle
var postprocessors []handle

var roximitor []handle
var postprocessor []handle

var shit []handle
var shits []handle

var handlers []*handleMsg          //All msg handle
var handlerPrivates []*handleMsg   //Private msg handle
var handlerGroups []*handleMsg     //Group msg handle
var handlerTemporarys []*handleMsg //Temporary msg handle

type powerHandle map[*handleMsg]uint8
type handlePower map[*powerHandle][]int64

var handleByPower handlePower

func (hp *handlePower) containUpdate(id int64, power uint8) bool {
	for k1 := range *hp {
		for _, v2 := range *k1 {
			if v2 <= power {
				(*hp)[k1] = append((*hp)[k1], id)
			} else {
				(*hp)[k1] = remove((*hp)[k1], id)
			}
		}
	}
	return true
}
func (hp *handlePower) containDelete(id ...int64) {
	for k := range *hp {
		(*hp)[k] = remove((*hp)[k], id...)
	}
}

type section map[string][]int64
type handleSection map[*handleMsg]section

var handlerBySection handleSection //section msg handle
func (s *section) getIds() []int64 {
	rs := make([]int64, 0)
	for _, v := range *s {
		rs = append(rs, v...)
	}
	return rs
}
func (hs *handleSection) containUpdate(tag string, ids ...int64) bool {
	for a, v := range *hs {
		for k, j := range v {
			if k == tag {
				(*hs)[a][k] = append(j, ids...)
			}
		}
	}
	return true
}
func (hs *handleSection) containDelete(tag string, ids ...int64) bool {
	for a, v := range *hs {
		for k, j := range v {
			if k == tag {
				(*hs)[a][k] = remove(j, ids...)
			}
		}
	}
	return true
}
func remove(is []int64, ids ...int64) []int64 {
	for n, i := range is {
		for _, j := range ids {
			if i == j {
				is = append(is[:n], is[n+1:]...)
			}
		}
	}
	return is
}

func handleMsgInit() {
	fmt.Println("===================消息格式=======================")
	for _, h := range handlers {
		if h.Parsing() {
			fmt.Println(h.pattern)
		}
	}
	for _, h := range handlerPrivates {
		if h.Parsing() {
			fmt.Println(h.pattern)
		}
	}
	for _, h := range handlerGroups {
		if h.Parsing() {
			fmt.Println(h.pattern)
		}
	}
	for _, h := range handlerTemporarys {
		if h.Parsing() {
			fmt.Println(h.pattern)
		}
	}
	for k := range handleByPower {
		for h := range *k {
			if h.Parsing() {
				fmt.Println(h.pattern)
			}
		}
	}
	for h := range handlerBySection {
		if h.Parsing() {
			fmt.Println(h.pattern)
		}
	}
}

// func analysis(patterns []string) map[string]any {
// 	var res = make(map[string]any)
// 	for _, p := range patterns {

// 	}

// 	return res
// }

type patternFormat struct {
	pName   string
	pRegex  string
	pType   string
	pCqCode string
}

func (p *patternFormat) GetPName() string {
	return (*p).pName
}
func (p *patternFormat) GetPRegex() string {
	return (*p).pRegex
}
func (p *patternFormat) GetPType() string {
	return (*p).pType
}
func (p *patternFormat) GetPCqCode() string {
	return (*p).pCqCode
}
func newPatternFormat(n string, r string, t string, c string) *patternFormat {
	return &patternFormat{
		pName:   n, //parameters name
		pRegex:  r, //regex
		pType:   t, //parameters type
		pCqCode: c, //cqcode type
	}
}

//解析参数
func analysisPattern(pattern string) *patternFormat {
	if !patternMatch(pattern) {
		return nil
	}
	pattern = pattern[1 : len(pattern)-1]
	patterns := strings.Split(pattern, ",")
	rn := regexp.MustCompile(`[0-9a-zA-Z]+`)
	if len(patterns) == 0 {
		return nil
	}
	rstr := ""
	res := newPatternFormat("", "", "", "")
	for i := 0; i < len(patterns); i++ {
		patternsl := strings.Split(patterns[i], ":")
		if len(patternsl) == 1 {
			if len((*res).pName) == 0 && rn.Match([]byte(patternsl[0])) {
				(*res).pName = patternsl[0]
			} else {
				rstr += patternsl[0]
			}
		} else if len(patternsl) == 2 {
			switch patternsl[0] {
			case "n":
				(*res).pName = patternsl[1]
			case "t":
				(*res).pType = patternsl[1]
			case "c":
				(*res).pCqCode = patternsl[1]
			default:
				rstr += patternsl[0] + ":" + patternsl[1]
				rstr += ","
			}
		} else {
			rstr += patternsl[i]
			for i := 1; i < len(patternsl); i++ {
				rstr += ":" + patternsl[i]
			}
			rstr += ","
		}
	}
	if len(rstr) > 1 {
		if rstr[:2] == "r:" {
			if rstr[len(rstr)-1:] == "," {
				res.pRegex = rstr[2 : len(rstr)-1]
			} else {
				res.pRegex = rstr[2:]
			}
		} else {
			return nil
		}
	}
	return res
}
func patternMatch(pattern string) bool {
	if len(pattern) < 3 {
		return false
	}
	if pattern[:1] != "{" || pattern[len(pattern)-1:] != "}" {
		return false
	}
	rn := regexp.MustCompile(`^\{(n:)?[0-9a-zA-Z]+(,((t:[0-9a-z]+)|(c:[a-z]+))){0,2}(,r:.*)?\}$`)
	return rn.Match([]byte(pattern))
}

//分解pattern
func getPatternOrder(pattern string) []string {
	res := make([]string, 0)
	patterns := strings.Split(pattern, "")
	var equilibrium = 0
	var tempStr = ""
	for i := 0; i < len(patterns); i++ {
		if patterns[i] == "{" {
			equilibrium++
			if len(tempStr) != 0 {
				res = append(res, tempStr)
			}
			tempStr = patterns[i]
			for j := i + 1; j < len(patterns); j++ {
				if patterns[j] == "}" {
					equilibrium--
					tempStr += patterns[j]
					if equilibrium == 0 {
						res = append(res, tempStr)
						tempStr = ""
						i = j
						break
					}
				} else if patterns[j] == "{" {
					equilibrium++
					tempStr += patterns[j]
				} else {
					tempStr += patterns[j]
				}
			}
		} else {
			tempStr += patterns[i]
		}
	}
	if len(tempStr) != 0 {
		res = append(res, tempStr)
	}
	return res
}

/**
 * pattern contain {} is parameter, {p:str} is msg.Parameter[str] , {p:abc,g:[\d+\s]} is satisfy regex msg.Parameter[abc]
 * {t:} is type int、float、string... ,
**/
// func getPatternCode(pattern string) []*patternFormat {
// 	patterns := strings.Split(pattern, "")
// 	var res = make([]*patternFormat, 0)
// 	var equilibrium = 0                  //Special symbol balance parameters
// 	for i := 0; i < len(patterns); i++ { //Traversal string
// 		if patterns[i] == "\\" && i != len(patterns)-1 && patterns[i+1] == "{" {
// 			patterns = append(patterns[:i-1], patterns[i+1:]...)
// 		} else if patterns[i] == "\\" && i != len(patterns)-1 && patterns[i+1] == "}" {
// 			patterns = append(patterns[:i-1], patterns[i+1:]...)
// 		} else if patterns[i] == "{" { //Processing cq parameters
// 			j := i + 1
// 			tempStr := ""
// 			var ( //Corresponding patternFormat
// 				n = ""
// 				r = ""
// 				t = ""
// 				c = ""
// 			)
// 			equilibrium++
// 			for ; j < len(patterns); j++ { //When encountering parameter format
// 				if patterns[j] == "r" && j < len(patterns)-1 && patterns[j+1] == ":" {
// 					for k := j + 2; k < len(patterns); k++ { //Regex special analysis
// 						if patterns[k] == "{" && patterns[k-1] != "\\" {
// 							equilibrium++
// 						} else if patterns[k] == "}" && patterns[k-1] != "\\" {
// 							equilibrium--
// 							if equilibrium == 0 { //Encountered the last }
// 								r = tempStr
// 								tempStr = ""
// 								j = k
// 								break
// 							}
// 						} else if patterns[k] == "," && patterns[k-1] != "\\" && k != len(patterns)-2 { //Comma encountered
// 							if patterns[k+2] == ":" && (patterns[k+1] == "n" || patterns[k+1] == "t" || patterns[k+1] == "c") {
// 								r = tempStr
// 								tempStr = ""
// 								j = k + 1
// 								break
// 							}
// 						}
// 						tempStr += patterns[k] //collect regex string
// 					}
// 				}
// 				if equilibrium == 0 {
// 					tempStr = ""
// 					i = j
// 					break
// 				}
// 				if (patterns[j] == "," && j != i+1 && patterns[j-1] != "\\") || patterns[j] == "}" && j != i+1 && patterns[j-1] != "\\" { //Collect sub parameters
// 					if len(tempStr) == 0 {
// 						return nil
// 					}
// 					formatMap := strings.Split(tempStr, ":")
// 					if len(formatMap) == 1 {
// 						if len(n) != 0 {
// 							return nil
// 						}
// 						n = tempStr
// 					} else if len(formatMap) == 2 {
// 						switch formatMap[0] {
// 						case "n":
// 							if len(n) != 0 {
// 								return nil
// 							}
// 							n = formatMap[1]
// 						case "t":
// 							if len(t) != 0 {
// 								return nil
// 							}
// 							t = formatMap[1]
// 						case "c":
// 							if len(c) != 0 {
// 								return nil
// 							}
// 							c = formatMap[1]
// 						default:
// 							return nil
// 						}
// 					} else {
// 						return nil
// 					}
// 					if patterns[j] == "}" && j != i+1 && patterns[j-1] != "\\" { //Encountered the last }
// 						equilibrium--
// 						if equilibrium == 0 {
// 							tempStr = ""
// 							i = j
// 							break
// 						}
// 					} else {
// 						tempStr = ""
// 						continue
// 					}
// 				}
// 				tempStr += patterns[j]

// 			}
// 			if equilibrium != 0 {
// 				return nil
// 			}
// 			res = append(res, newPatternFormat(n, r, t, c))
// 		}
// 	}
// 	return res
// }
