package cq

var (
	adminSupers []int64         //超级管理员列表
	admins      map[int64]uint8 //管理员列表
	whitelists  []int64         //白名单
	blacklists  []int64         //黑名单
)

// 获取超级管理员列表
func (ch *CqH) GetAdminSupers() []int64 {
	return adminSupers
}

// 获取管理员字典
func (ch *CqH) GetAdmins() map[int64]uint8 {
	return admins
}

// 获取白名单
func (ch *CqH) GetWhitelists() []int64 {
	return whitelists
}

// 获取黑名单
func (ch *CqH) Getblacklists() []int64 {
	return blacklists
}

//添加超级管理员
func (ch *CqH) AddAdminSupers(ids ...int64) {
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if !hasListAdmin(adminSupers, id) {
			adminSupers = append(adminSupers, id)
		}
	}
	handlerBySection.containUpdate("admin6", ids...)
}

//添加管理员
func (ch *CqH) AddAdmins(id int64, power uint8) {
	admins[id] = power
	handleByPower.containUpdate(id, power)
}

//添加白名单
func (ch *CqH) AddWhitelists(ids ...int64) {
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if !hasListAdmin(whitelists, id) {
			whitelists = append(whitelists, id)
		}

	}
	handlerBySection.containUpdate("whitelist", ids...)
}

//添加黑名单
func (ch *CqH) AddBlacklists(ids ...int64) {
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if !hasListAdmin(blacklists, id) {
			blacklists = append(blacklists, id)
		}

	}
	handlerBySection.containUpdate("blacklist", ids...)
}

// 删除超级管理员
func (ch *CqH) DelAdminSupers(ids ...int64) {
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		delList(&adminSupers, id)
	}
	handlerBySection.containDelete("admin6", ids...)
}

//删除管理员
func (ch *CqH) DelAdmins(ids ...int64) {
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		delete(admins, id)
	}
	handleByPower.containDelete(ids...)
}

//删除白名单
func (ch *CqH) DelWhitelists(ids ...int64) {
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		delList(&whitelists, id)
	}
	handlerBySection.containDelete("whitelist", ids...)
}

//删除黑名单
func (ch *CqH) DelBlacklists(ids ...int64) {
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		delList(&blacklists, id)
	}
	handlerBySection.containDelete("blacklist", ids...)
}

func delList(l *[]int64, id int64) {
	for i, v := range *l {
		if v == id {
			*l = append((*l)[:i], (*l)[i+1:]...)
		}
	}
}

func hasListAdmin(list []int64, key int64) bool {
	if len(list) == 0 {
		return false
	}
	for _, l := range list {
		if l == key {
			return true
		}
	}
	return false
}

// func hasMapAdmin(mp *map[int64]uint8, key int64) bool {
// 	if len(*mp) == 0 {
// 		return false
// 	}
// 	for k := range *mp {
// 		if k == key {
// 			return true
// 		}
// 	}
// 	return false
// }
