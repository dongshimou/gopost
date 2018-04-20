package utility

import (
	"strconv"
	"time"
)

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
func ParseDate(datetime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", datetime, time.Local)
}
func ParseTime(datetime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
}
