package common

type Error struct {
	error
	ErrorString string
	ErrorCode   int
	ErrorType   int
}

func (e Error) Error() string {
	return e.ErrorString
}

func (e Error) Code() int {
	return e.ErrorCode
}

func (e Error) Type() int {
	return int(e.ErrorType)
}

var (
	UNKNOWNERROR = Error{
		ErrorCode:   -1,
		ErrorType:   SERVERERRORCODE,
		ErrorString: "unknown error. maybe server is error. please wait for sometime",
	}
	USERINPUTERROR = Error{
		ErrorCode:   100001,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "please check your input, there is something wrong",
	}
	HAVENOPERMISSION = Error{
		ErrorCode:   100002,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "you have no access to this operation",
	}
	DATABASEERROR = Error{
		ErrorCode:   100003,
		ErrorType:   SERVERERRORCODE,
		ErrorString: "server's database has some error, please try again later",
	}
	USERDOESNOTEXIST = Error{
		ErrorCode:   100004,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "user does not exist. please check",
	}
	PASSWORDISERROR = Error{
		ErrorCode:   100005,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "password is incorrect. please try again",
	}
	USERNOTLOGIN = Error{
		ErrorCode:   100006,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "you do not login. please login",
	}
	EXCEEDTIMELIMIT = Error{
		ErrorCode:   100007,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "your token has no time. please login again",
	}
	USEREXISTED = Error{
		ErrorCode:   100008,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "account has existed. please rename your account",
	}
	GROUPNOTEXIST = Error{
		ErrorCode:   100009,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "this group dose not exist. maybe your input is error",
	}
	GROUPEXIST = Error{
		ErrorCode:   100010,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "this group has already been created. please use another name",
	}
	DATANOTFOUND = Error{
		ErrorCode:   100011,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "data is not in database. please check your input",
	}
	HAVEBEENFRIEND = Error{
		ErrorCode:   100012,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "you both are already friends, no need to be friend again",
	}
	NOTFRIEND = Error{
		ErrorCode:   100013,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "you both are not friend, no need to delete friend",
	}
	REDISERROR = Error{
		ErrorCode:   100014,
		ErrorType:   SERVERERRORCODE,
		ErrorString: "server redis db error, please try later",
	}
	USERNOTACTIVE = Error{
		ErrorCode:   100015,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "you do not active your email, please active your account to do something next",
	}
	ACTIVECODEERROR = Error{
		ErrorCode:   100016,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "active code is not correct, please check",
	}
	IMGFORMATERROR = Error{
		ErrorCode:   100017,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "we only support jpg/jpeg/png image, please upload correct image",
	}
	IMGTOOLARGEERROR = Error{
		ErrorCode:   100018,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "this image is too big, please upload less smaller image",
	}
	FILEUPLOADERROR = Error{
		ErrorCode:   100019,
		ErrorType:   CLIENTERRORCODE,
		ErrorString: "file upload fail, maybe it's server error, please wait",
	}
)
