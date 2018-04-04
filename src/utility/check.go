package utility

import (
	"regexp"
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

func VerifyPermission(up int, nps ...int) error {
	for _, np := range nps {
		if up&np <= 0 {
			return NewError(ERROR_AUTH_CODE, ERROR_AUTH_MSG)
		}
	}
	return nil
}

func VerifyUsername(username string) error {
	l := len(username)
	if l < 6 || l > 30 {
		return NewError(ERROR_REQUEST_CODE, ERROR_MSG_LENGTH_USERNAME)
	}
	b, _ := regexp.MatchString("^[0-9a-zA-Z]{6,30}$", username)
	if !b {
		return NewError(ERROR_REQUEST_CODE, ERROR_MSG_ERROR_USERNAME)
	}
	return nil
}
