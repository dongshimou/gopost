package utility

func VerifyPermission(up int, nps ...int) error {
	for _, np := range nps {
		if up&np <= 0 {
			return NewError(ERROR_AUTH_CODE, ERROR_AUTH_MSG)
		}
	}
	return nil
}

func CreatePermission(nps ...int) int {
	up := 0
	for _, np := range nps {
		up += np
	}
	return up
}
