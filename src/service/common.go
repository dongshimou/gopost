package service

import (
	"gopost/src/utility"
	"strings"
	"time"
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

// parsePage 解析页码，默认返回 1
func parsePage(s string) int {
	if s == "" {
		return 1
	}
	val, err := parse2uint(s)
	if err != nil {
		return 1
	}
	if val < 1 {
		return 1
	}
	return int(val)
}

// parsePageSize 解析每页数量，默认返回 10
func parsePageSize(s string) int {
	if s == "" {
		return 10
	}
	val, err := parse2uint(s)
	if err != nil {
		return 10
	}
	if val < 1 {
		return 10
	}
	if val > 100 {
		return 100 // 最大 100
	}
	return int(val)
}

// calcOffset 计算分页偏移量
func calcOffset(page, pageSize int) int {
	return (page - 1) * pageSize
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
