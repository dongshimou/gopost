package service

import (
	"strings"
	"time"
	"gopost/src/utility"
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
func parseTime(s string) (time.Time, error) {
	return utility.ParseTime(s)
}
func parseUnix(s string) (time.Time, error) {
	i, err := utility.Parse2Int64(s)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}
func parsrDate(s string) (time.Time, error) {
	return utility.ParseDate(s)
}
func formatDatetime(time time.Time) string {
	return utility.FormatDatetime(time)
}
func buildArgs(split string, args ...string) string {
	return strings.Join(args, split)
}
