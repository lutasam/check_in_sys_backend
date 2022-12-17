package utils

import (
	"encoding/json"
	"github.com/lutasam/check_in_sys/biz/common"
	"strconv"
	"time"
)

func Uint64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func StringToUint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, common.USERINPUTERROR
	}
	return i, nil
}

func StringToFloat32(s string) (float32, error) {
	temp, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, common.USERINPUTERROR
	}
	return float32(temp), nil
}

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, common.USERINPUTERROR
	}
	return i, nil
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimeToDateString(t time.Time) string {
	return t.Format("2006-01-02")
}

func StructToMap(s interface{}) (map[string]interface{}, error) {
	result, err := json.Marshal(s)
	if err != nil {
		return nil, common.UNKNOWNERROR
	}
	var m map[string]interface{}
	err = json.Unmarshal(result, &m)
	if err != nil {
		return nil, common.UNKNOWNERROR
	}
	return m, nil
}
