package service

import (
	"fmt"
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
func parseCount(s string) (uint, error) {
	return parse2uint(s)
}
func formatDatetime(time time.Time) string {
	return utility.FormatDatetime(time)
}
func buildArgs(split string, args ...interface{}) string {
	l := len(args)
	v := args[l-1]
	if l == 1 {
		return fmt.Sprintf("%v", v)
	} else {
		return buildArgs(split, args[:len(args)-1]...) + fmt.Sprintf("%s%v", split, v)
	}
}
