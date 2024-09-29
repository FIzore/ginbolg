package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code=1000...用户模块的错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002 //登陆时密码错误
	ERROR_USER_NOT_EXIST = 1003 //用户不存在
	ERROR_TOKEN_EXIST    = 1004 //token不存在
	ERROR_TOKEN_RUNTIME  = 1005 //token已过期
	ERROR_TOKEN_WRONG    = 1006 //token不正确
	// code=2000...文章模块的错误

	// code=3000...分类模块的错误

)

var codemsg = map[int]string{ //

}

func GetErrMsg(code int) string {
	return codemsg[code]
}
