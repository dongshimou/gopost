package service

import (
	"model"
	"time"
	"utility"
)

func isNullOrEmpty(v interface{}) bool {
	return utility.IsNullorEmpty(v)
}
func parse2uint(s string) (uint, error) {
	return utility.Parse2Uint(s)
}
func parseID(s string) (uint, error) {
	return parse2uint(s)
}
func formatDatetime(time time.Time) string {
	return utility.FormatDatetime(time)
}

func SignIn(req *model.REQSignin) (*model.RESSignIn, error) {

	return nil, nil
}
func SignUp(req *model.REQSignUp) (*model.RESSignUp, error) {

	return nil, nil
}
