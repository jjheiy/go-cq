package cq

import (
	"encoding/json"
)

//发送私聊消息参数
type userMsg struct {
	Id          int64  `json:"user_id"`
	Msg         string `json:"message"`
	Group_id    int64  `json:"group_id,omitempty"`
	Auto_escape bool   `json:"auto_escape"`
}

func newUserMsg(id int64, msg string, auto bool) *userMsg {
	return &userMsg{
		Id:          id,
		Msg:         msg,
		Auto_escape: auto,
	}
}
func newUserMsgByGroup(id int64, msg string, group_id int64, auto bool) *userMsg {
	return &userMsg{
		Id:          id,
		Msg:         msg,
		Group_id:    group_id,
		Auto_escape: auto,
	}
}
func (um *userMsg) jsonToBytes() []byte {
	res, err := json.Marshal(*um)
	if err != nil {
		return nil
	}
	return res
}

//发送群聊消息参数
type gloupMsg struct {
	Id          int64  `json:"group_id"`
	Msg         string `json:"message"`
	Auto_escape bool   `json:"auto_escape"`
}

func newGloupMsg(id int64, msg string, auto bool) *gloupMsg {
	return &gloupMsg{
		Id:          id,
		Msg:         msg,
		Auto_escape: auto,
	}
}
func (gm *gloupMsg) jsonToBytes() []byte {
	res, err := json.Marshal(*gm)
	if err != nil {
		return nil
	}
	return res
}

//撤回消息/获取消息参数
type idMsg struct {
	Id int32 `json:"message_id"`
}

func newIdMsg(id int32) *idMsg {
	return &idMsg{
		Id: id,
	}
}
func (im *idMsg) jsonToBytes() []byte {
	res, err := json.Marshal(*im)
	if err != nil {
		return nil
	}
	return res
}

//撤回消息/获取消息
type imgMsg struct {
	File string `json:"file"`
}

func newImgMsg(file string) *imgMsg {
	return &imgMsg{
		File: file,
	}
}
func (im *imgMsg) jsonToBytes() []byte {
	res, err := json.Marshal(*im)
	if err != nil {
		return nil
	}
	return res
}

type groupKick struct {
	Id        int64 `json:"group_id"`
	UserId    int64 `json:"user_id"`
	RejectAdd bool  `json:"reject_add_request"`
}

func newGroupKick(id, userId int64, rejectAdd bool) *groupKick {
	return &groupKick{
		Id:        id,
		UserId:    userId,
		RejectAdd: rejectAdd,
	}
}
func (gk *groupKick) jsonToBytes() []byte {
	res, err := json.Marshal(*gk)
	if err != nil {
		return nil
	}
	return res
}

type groupBan struct {
	Id       int64 `json:"group_id"`
	UserId   int64 `json:"user_id"`
	Duration int   `json:"duration"`
}

func newGroupBan(id, userId int64, duration int) *groupBan {
	return &groupBan{
		Id:       id,
		UserId:   userId,
		Duration: duration,
	}
}
func (gb *groupBan) jsonToBytes() []byte {
	res, err := json.Marshal(*gb)
	if err != nil {
		return nil
	}
	return res
}

type groupAllBan struct {
	Id     int64 `json:"group_id"`
	Enable bool  `json:"enable"`
}

func newGroupAllBan(id int64, enable bool) *groupAllBan {
	return &groupAllBan{
		Id:     id,
		Enable: enable,
	}
}
func (gab *groupAllBan) jsonToBytes() []byte {
	res, err := json.Marshal(*gab)
	if err != nil {
		return nil
	}
	return res
}

type groupSetAdmin struct {
	Id     int64 `json:"group_id"`
	UserId int64 `json:"user_id"`
	Enable bool  `json:"enable"`
}

func newGroupSetAdmin(id int64, userId int64, enable bool) *groupSetAdmin {
	return &groupSetAdmin{
		Id:     id,
		UserId: userId,
		Enable: enable,
	}
}
func (gsa *groupSetAdmin) jsonToBytes() []byte {
	res, err := json.Marshal(*gsa)
	if err != nil {
		return nil
	}
	return res
}

type groupSetCard struct {
	Id     int64  `json:"group_id"`
	UserId int64  `json:"user_id"`
	Card   string `json:"card"`
}

func newGroupSetCard(id int64, userId int64, card string) *groupSetCard {
	return &groupSetCard{
		Id:     id,
		UserId: userId,
		Card:   card,
	}
}
func (gsc *groupSetCard) jsonToBytes() []byte {
	res, err := json.Marshal(*gsc)
	if err != nil {
		return nil
	}
	return res
}

type groupSetName struct {
	Id   int64  `json:"group_id"`
	Name string `json:"group_name"`
}

func newGroupSetName(id int64, name string) *groupSetName {
	return &groupSetName{
		Id:   id,
		Name: name,
	}
}
func (gsn *groupSetName) jsonToBytes() []byte {
	res, err := json.Marshal(*gsn)
	if err != nil {
		return nil
	}
	return res
}

type groupLeave struct {
	Id      int64 `json:"group_id"`
	Dismiss bool  `json:"is_dismiss"`
}

func newGroupLeave(id int64, dismiss bool) *groupLeave {
	return &groupLeave{
		Id:      id,
		Dismiss: dismiss,
	}
}
func (gl *groupLeave) jsonToBytes() []byte {
	res, err := json.Marshal(*gl)
	if err != nil {
		return nil
	}
	return res
}

type groupTitle struct {
	Id       int64  `json:"group_id"`
	UserId   int64  `json:"user_id"`
	Title    string `json:"special_title"`
	Duration int    `json:"duration"`
}

func newGroupTitle(id int64, userId int64, title string, duration int) *groupTitle {
	return &groupTitle{
		Id:       id,
		UserId:   userId,
		Title:    title,
		Duration: duration,
	}
}
func (gt *groupTitle) jsonToBytes() []byte {
	res, err := json.Marshal(*gt)
	if err != nil {
		return nil
	}
	return res
}

type groupIdp struct {
	Id int64 `json:"group_id"`
}

func newGroupIdp(id int64) *groupIdp {
	return &groupIdp{
		Id: id,
	}
}
func (gi *groupIdp) jsonToBytes() []byte {
	res, err := json.Marshal(*gi)
	if err != nil {
		return nil
	}
	return res
}

type msgIdp struct {
	Id int32 `json:"message_id"`
}

func newMsgIdp(id int32) *msgIdp {
	return &msgIdp{
		Id: id,
	}
}
func (mi *msgIdp) jsonToBytes() []byte {
	res, err := json.Marshal(*mi)
	if err != nil {
		return nil
	}
	return res
}

type userIdp struct {
	Id int64 `json:"user_id"`
}

func newUserIdp(id int64) *userIdp {
	return &userIdp{
		Id: id,
	}
}
func (ui *userIdp) jsonToBytes() []byte {
	res, err := json.Marshal(*ui)
	if err != nil {
		return nil
	}
	return res
}

type setAiInfo struct {
	Nickname     string `json:"nickname"`
	Company      string `json:"company"`
	Email        string `json:"email"`
	College      string `json:"college"`
	PersonalNote string `json:"personal_note"`
}

func newSetAiInfo(nickname, company, email, college, personalNote string) *setAiInfo {
	return &setAiInfo{
		Nickname:     nickname,
		Company:      company,
		Email:        email,
		College:      college,
		PersonalNote: personalNote,
	}
}
func (sai *setAiInfo) jsonToBytes() []byte {
	res, err := json.Marshal(*sai)
	if err != nil {
		return nil
	}
	return res
}

type strangerInfo struct {
	Id      int64 `json:"user_id"`
	NoCache bool  `json:"no_cache"`
}

func newStrangerInfo(id int64, noCache bool) *strangerInfo {
	return &strangerInfo{
		Id:      id,
		NoCache: noCache,
	}
}
func (si *strangerInfo) jsonToBytes() []byte {
	res, err := json.Marshal(*si)
	if err != nil {
		return nil
	}
	return res
}

type groupInfo struct {
	Id      int64 `json:"group_id"`
	NoCache bool  `json:"no_cache"`
}

func newGroupInfo(id int64, noCache bool) *groupInfo {
	return &groupInfo{
		Id:      id,
		NoCache: noCache,
	}
}
func (gi *groupInfo) jsonToBytes() []byte {
	res, err := json.Marshal(*gi)
	if err != nil {
		return nil
	}
	return res
}

type groupUserInfo struct {
	Id      int64 `json:"group_id"`
	UserId  int64 `json:"user_id"`
	NoCache bool  `json:"no_cache"`
}

func newGroupUserInfo(id int64, userId int64, noCache bool) *groupUserInfo {
	return &groupUserInfo{
		Id:      id,
		UserId:  userId,
		NoCache: noCache,
	}
}
func (gui *groupUserInfo) jsonToBytes() []byte {
	res, err := json.Marshal(*gui)
	if err != nil {
		return nil
	}
	return res
}

type groupHonorInfo struct {
	Id   int64  `json:"group_id"`
	Type string `json:"type"`
}

func newGroupHonorInfo(id int64, htype string) *groupHonorInfo {
	return &groupHonorInfo{
		Id:   id,
		Type: htype,
	}
}
func (ghi *groupHonorInfo) jsonToBytes() []byte {
	res, err := json.Marshal(*ghi)
	if err != nil {
		return nil
	}
	return res
}

type setGroupImg struct {
	Id    int64  `json:"group_id"`
	File  string `json:"file"`
	Cache int    `json:"cache"`
}

func newSetGroupImg(id int64, file string, cache int) *setGroupImg {
	return &setGroupImg{
		Id:    id,
		File:  file,
		Cache: cache,
	}
}
func (sgi *setGroupImg) jsonToBytes() []byte {
	res, err := json.Marshal(*sgi)
	if err != nil {
		return nil
	}
	return res
}

type upUserFile struct {
	Id   int64  `json:"user_id"`
	File string `json:"file"`
	Name string `json:"name"`
}

func newUpUserFile(id int64, file string, name string) *upUserFile {
	return &upUserFile{
		Id:   id,
		File: file,
		Name: name,
	}
}
func (uuf *upUserFile) jsonToBytes() []byte {
	res, err := json.Marshal(*uuf)
	if err != nil {
		return nil
	}
	return res
}

type upGroupFile struct {
	Id     int64  `json:"group_id"`
	File   string `json:"file"`
	Name   string `json:"name"`
	Folder string `json:"folder,omitempty"`
}

func newUpGroupFile(id int64, file string, name string) *upGroupFile {
	return &upGroupFile{
		Id:   id,
		File: file,
		Name: name,
	}
}
func (ugf *upGroupFile) jsonToBytes() []byte {
	res, err := json.Marshal(*ugf)
	if err != nil {
		return nil
	}
	return res
}

type getGroupFileList struct {
	Id     int64  `json:"group_id"`
	Folder string `json:"folder_id"`
}

func newGetGroupFileList(id int64, folder string) *getGroupFileList {
	return &getGroupFileList{
		Id:     id,
		Folder: folder,
	}
}
func (ggfl *getGroupFileList) jsonToBytes() []byte {
	res, err := json.Marshal(*ggfl)
	if err != nil {
		return nil
	}
	return res
}

type createFolder struct {
	Id       int64  `json:"group_id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

func newCreateFolder(id int64, name string) *createFolder {
	return &createFolder{
		Id:       id,
		Name:     name,
		ParentId: "/",
	}
}
func (cf *createFolder) jsonToBytes() []byte {
	res, err := json.Marshal(*cf)
	if err != nil {
		return nil
	}
	return res
}

type groupFile struct {
	Id     int64  `json:"group_id"`
	FileId string `json:"file_id"`
	Busid  int32  `json:"busid"`
}

func newGroupFile(id int64, fileId string, busid int32) *groupFile {
	return &groupFile{
		Id:     id,
		FileId: fileId,
		Busid:  busid,
	}
}
func (gf *groupFile) jsonToBytes() []byte {
	res, err := json.Marshal(*gf)
	if err != nil {
		return nil
	}
	return res
}

type sendGroupNotice struct {
	Id      int64  `json:"group_id"`
	Content string `json:"content"`
	Image   string `json:"image,omitempty"`
}

func newSendGroupNotice(id int64, content string) *sendGroupNotice {
	return &sendGroupNotice{
		Id:      id,
		Content: content,
	}
}
func (sgn *sendGroupNotice) jsonToBytes() []byte {
	res, err := json.Marshal(*sgn)
	if err != nil {
		return nil
	}
	return res
}

type getGroupMsgHistory struct {
	Id         int64 `json:"group_id"`
	MessageSeq int64 `json:"message_seq"`
}

func newGetGroupMsgHistory(id int64, messageSeq int64) *getGroupMsgHistory {
	return &getGroupMsgHistory{
		Id:         id,
		MessageSeq: messageSeq,
	}
}
func (ggmh *getGroupMsgHistory) jsonToBytes() []byte {
	res, err := json.Marshal(*ggmh)
	if err != nil {
		return nil
	}
	return res
}

type urlStr struct {
	Url string `json:"url"`
}

func newUrlStr(url string) *urlStr {
	return &urlStr{
		Url: url,
	}
}
func (us *urlStr) jsonToBytes() []byte {
	res, err := json.Marshal(*us)
	if err != nil {
		return nil
	}
	return res
}
