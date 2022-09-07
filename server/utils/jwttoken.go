package utils

import (
	"txc-njl/server/jwtdemo"
	"txc-njl/server/model"
)

func Login(p *model.User) (token string, err error) {

	//生成jwt的token
	return jwtdemo.GenToken(int64(p.ID), p.Name)
}
