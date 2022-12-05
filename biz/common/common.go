package common

import "time"

const ISSUER = "LUTASAM"                                // jwt issuer
const PASSWORDSALT = "astaxie12798akljzmknm.ahkjkljl;k" // use only for password encryption
const OTHERSECRETSALT = "9871267812345mn812345xyz"      // user for other encryption
const EXPIRETIME = 86400000                             // jwt expiration time. 1 day's second
const ACTIVECODEEXPTIME = 300 * time.Second             // active code expiration time. 5 min
const ACTIVECODESUFFIX = "_active_code"
const DEFAULTAVATARURL = "http://baidu.com/test.png"
const MAXIMGSPACE = 1024 * 1024 * 1 // img upload should be less than 1 MB
const ALLDEPARTMENTS = 0            // symbolize all the departments

const (
	STATUSOKCODE    = 200
	CLIENTERRORCODE = 400
	SERVERERRORCODE = 500
)

const (
	STATUSOKMSG    = "OK"
	CLIENTERRORMSG = "400 CLIENT ERROR"
	SERVERERRORMSG = "500 SERVER ERROR"
)

// health code status
type HealthCodeStatus int

const (
	ALLHEALTHCODE HealthCodeStatus = iota
	GREEN
	GREY
	YELLOW
	RED
)

func (s HealthCodeStatus) Ints() int {
	return int(s)
}

func ParseHealthCodeStatus(i int) HealthCodeStatus {
	switch i {
	case 0:
		return ALLHEALTHCODE
	case 1:
		return GREEN
	case 2:
		return GREY
	case 3:
		return YELLOW
	case 4:
		return RED
	default:
		return GREEN
	}
}

func (s HealthCodeStatus) String() string {
	switch s {
	case GREEN:
		return "绿码"
	case GREY:
		return "灰码"
	case YELLOW:
		return "黄码"
	case RED:
		return "红码"
	default:
		return "未知"
	}
}

// identity
type Identity int

const (
	USER Identity = iota
	DEPARTMENT_ADMIN
	SUPER_ADMIN
)

func (s Identity) Ints() int {
	return int(s)
}

func ParseIdentity(i int) Identity {
	switch i {
	case 0:
		return USER
	case 1:
		return DEPARTMENT_ADMIN
	case 2:
		return SUPER_ADMIN
	default:
		return USER
	}
}

func (s Identity) String() string {
	switch s {
	case USER:
		return "普通用户"
	case DEPARTMENT_ADMIN:
		return "部门管理员"
	case SUPER_ADMIN:
		return "超级管理员"
	default:
		return "未知"
	}
}

type TemperatureRange int

const (
	Below36 TemperatureRange = iota
	Between36_37
	Between37_38
	Between38_39
	Above39
)

func (s TemperatureRange) Ints() int {
	return int(s)
}

func ParseTemperatureRange(i int) TemperatureRange {
	switch i {
	case 0:
		return Below36
	case 1:
		return Between36_37
	case 2:
		return Between37_38
	case 3:
		return Between38_39
	case 4:
		return Above39
	default:
		return Below36
	}
}

func (s TemperatureRange) String() string {
	switch s {
	case Below36:
		return "小于36°C"
	case Between36_37:
		return "36°C~37°C之间"
	case Between37_38:
		return "37°C~38°C之间"
	case Between38_39:
		return "38°C~39°C之间"
	case Above39:
		return "大于39°C"
	default:
		return "未知"
	}
}
