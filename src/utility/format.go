package utility

import "time"

func FormatDatetime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
