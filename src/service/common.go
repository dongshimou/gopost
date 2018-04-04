package service

import (
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
