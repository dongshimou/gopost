package service

import (
	"model"
	"utility"
)

func isNullOrEmpty(v interface{}) bool {
	return utility.IsNullorEmpty(v)
}
func parse2uint(s string) (uint, error) {
	return utility.Parse2Uint(s)
}
func parseID(s string)(uint ,error){
	return parse2uint(s)
}

func Login(req *model.REQLogin) (*model.RESLogin, error) {

	return nil, nil
}
