package utils

const (
	SUCCESS = 200
	ERROR   = 500

	// User module error
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008

	// Article module error
	ErrorArticleNotExist = 2001

	// Category module error
	ErrorCategoryNameUsed = 3001
	ErrorCategoryNotExist = 3002
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	// User module error
	ErrorUsernameUsed:   "Username has been used",
	ErrorPasswordWrong:  "Password is wrong",
	ErrorUserNotExist:   "User does not exist",
	ErrorTokenExist:     "Token does not exist",
	ErrorTokenRuntime:   "Token has expired",
	ErrorTokenWrong:     "Token is wrong",
	ErrorTokenTypeWrong: "Token format is wrong",
	ErrorUserNoRight:    "User has no right",

	// Article module error
	ErrorArticleNotExist: "Article does not exist",

	// Category module error
	ErrorCategoryNameUsed: "Category name has been used",
	ErrorCategoryNotExist: "Category does not exist",
}

func GetMsg(code int) string {
	msg, ok := codeMsg[code]
	if ok {
		return msg
	}
	return codeMsg[ERROR]
}
