package api

import (
	"jokyo3/probingai/api/helper"
	"jokyo3/probingai/common"
	"net/http"
)

type SysConfig struct {
	// 是否系统配置 cookie
	IsSysCK bool `json:"isSysCK"`
	// 是否已授权
	IsAuth bool `json:"isAuth"`
}

func SysConf(w http.ResponseWriter, r *http.Request) {
	isAuth := helper.CheckAuth(r)
	conf := SysConfig{
		IsSysCK: len(common.USER_TOKEN_LIST) > 0,
		IsAuth:  isAuth,
	}
	helper.SuccessResult(w, conf)
}
