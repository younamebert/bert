package vender

import (
	"time"

	"github.com/yudeguang/ratelimit"
)

// 限流器
func Userlimit() string {
	//    key := str, _ := vender.RedisGet(usertoken)
	//初始化若干条访问控制规则
	//每10秒只允许访问5次
	key := "aaa"
	//每30分钟只允许访问50次
	//每天只允许访问500次
	r := ratelimit.NewRule()
	r.AddRule(time.Second*10, 5)
	r.AddRule(time.Minute*30, 50)
	r.AddRule(time.Hour*24, 500)

	if r.AllowVisit(key) {

		aa := r.RemainingVisits(key)

		if aa[0] == 0 {
			return "600"
		}
		if aa[1] == 0 {
			return "601"
		}
		if aa[2] == 0 {
			return "602"
		}

	} else {
		return "500"
	}
	time.Sleep(time.Second * 1)
	return "1"
}
