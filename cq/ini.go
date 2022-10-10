package cq

import (
	"fmt"
	"net/http"
)

func New() *CqH {
	ini()
	return &CqH{}
}

//开始监听
func (ch *CqH) Run(ip string) error {
	fmt.Println("cq initialization...")
	handleMsgInit()
	http.HandleFunc("/", accept)
	fmt.Println("============================cq service started...=============================")
	return http.ListenAndServe(ip, nil)
}

func ini() {
	handlers = make([]*handleMsg, 0)
	handlerPrivates = make([]*handleMsg, 0)
	handlerGroups = make([]*handleMsg, 0)
	handlerTemporarys = make([]*handleMsg, 0)
	handleByPower = make(handlePower)
	handlerBySection = make(handleSection)

	adminSupers = make([]int64, 0)
	admins = make(map[int64]uint8)
	whitelists = make([]int64, 0)
	blacklists = make([]int64, 0)

	roximitor = make([]handle, 0)
	postprocessor = make([]handle, 0)
	roximitors = make([]handle, 0)
	postprocessors = make([]handle, 0)
	shit = make([]handle, 0)
	shits = make([]handle, 0)

	timedEventList = make([]*timedEvent, 0)
}
