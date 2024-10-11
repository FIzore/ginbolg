package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code=1000...用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002 //登陆时密码错误
	ERROR_USER_NOT_EXIST   = 1003 //用户不存在
	ERROR_TOKEN_EXIST      = 1004 //token不存在
	ERROR_TOKEN_RUNTIME    = 1005 //token已过期
	ERROR_TOKEN_WRONG      = 1006 //token不正确
	ERROR_TOKEN_TYPE_WRONG = 1007 //token格式错误
	// code=2000...文章模块的错误
	ERROR_ART_NOT_EXIST = 2001 //文章不存在
	// code=3000...分类模块的错误
	ERROR_CATEGORY_NOT_EXIST = 3000 //分类不存在
	ERROR_CATEGORY_USED      = 3001 //分类已存在

)

var codeMsg = map[int]string{ //
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "用户名已存在",
	ERROR_PASSWORD_WRONG:     "密码错误",     //登陆时密码错误
	ERROR_USER_NOT_EXIST:     "用户不存在",    //用户不存在
	ERROR_TOKEN_EXIST:        "token不存在", //token不存在
	ERROR_TOKEN_RUNTIME:      "token已过期", //token已过期
	ERROR_TOKEN_WRONG:        "token不正确",
	ERROR_TOKEN_TYPE_WRONG:   "token格式错误",
	ERROR_CATEGORY_USED:      "分类已存在",
	ERROR_ART_NOT_EXIST:      "文章不存在",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
