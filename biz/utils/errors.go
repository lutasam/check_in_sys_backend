package utils

import (
	"errors"
	"github.com/lutasam/check_in_sys/biz/common"
)

// Used to judge whether an error is in a error set
func IsIncludedByErrors(err error, errs ...error) bool {
	for _, e := range errs {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

func IsClientError(err error) bool {
	return err.(common.Error).ErrorType == common.CLIENTERRORCODE
}
