package utility

import (
	"strconv"
	"strings"
	"time"
)

func IsNullorEmpty(para interface{}) bool {
	switch v := para.(type) {
	case uint:
		if v == 0 {
			return true
		}
	case *uint:
		if v == nil || *v == 0 {
			return true
		}
	case int64:
		if v == 0 {
			return true
		}
	case *int64:
		if v == nil || *v == 0 {
			return true
		}
	case string:
		if strings.TrimSpace(v) == "" {
			return true
		}
	case *string:
		if v == nil || strings.TrimSpace(*v) == "" {
			return true
		}
	case *int:
		if v == nil || *v == 0 {
			return true
		}
	case int:
		if v == 0 {
			return true
		}
	case time.Time:
		if v.IsZero() {
			return true
		}
	case *time.Time:
		if v == nil || v.IsZero() {
			return true
		}
	default:
		return false
	}
	return false
}
func Parse2Uint(value interface{}) (uint, error) {
	switch v := value.(type) {
	case string:
		if IsNullorEmpty(v) {
			return 0, nil
		}
		return err_s2ui(v)
	case uint:
		return v, nil
	}
	return 0, NewError(ERROR_PARSE_CODE, ERROR_PARSE_MSG)
}
func Parse2Int(value interface{}) (int, error) {
	switch v := value.(type) {
	case string:
		if IsNullorEmpty(v) {
			return 0, nil
		}
		return err_s2i(v)
	case int:
		return v, nil
	}
	return 0, NewError(ERROR_PARSE_CODE, ERROR_PARSE_MSG)
}
func Parse2Int64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case string:
		if IsNullorEmpty(v) {
			return 0, nil
		}
		return err_s2i64(v)
	case int64:
		return v, nil
	}
	return 0, NewError(ERROR_PARSE_CODE, ERROR_PARSE_MSG)
}
func err_s2i(s string) (int, error) {
	id, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, NewError(ERROR_PARSE_CODE, ERROR_PARSE_MSG)
	} else {
		return int(id), err
	}
}
func safe_s2i(s string) int {
	id, _ := err_s2i(s)
	return int(id)
}

func safe_s2ui(s string) uint {
	id, _ := err_s2ui(s)
	return uint(id)
}
func err_s2ui(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, NewError(ERROR_PARSE_CODE, ERROR_PARSE_MSG)
	} else {
		return uint(id), err
	}
}
func err_s2i64(s string) (int64, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, NewError(ERROR_PARSE_CODE, ERROR_PARSE_MSG)
	} else {
		return int64(id), err
	}
}
func safe_s2i64(s string) int64 {
	id, _ := err_s2i64(s)
	return id
}
func parseTime(datetime string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", datetime)
	return t
}
func formatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
