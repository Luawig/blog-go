package utils

const (
	Success    = 200
	UnknownErr = 500

	// User module error
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008
	ErrorEmailUsed      = 1009
	ErrorUsernameEmpty  = 1010
	ErrorEmailEmpty     = 1011
	ErrorPasswordEmpty  = 1012

	// Article module error
	ErrorArticleNotExist = 2001

	// Category module error
	ErrorCategoryNameUsed  = 3001
	ErrorCategoryNotExist  = 3002
	ErrorCategoryNameEmpty = 3003

	// Comment module error
	ErrorCommentNotExist = 4001

	// Common error
	ErrorInvalidParam = 5001
)

var codeMsg = map[int]string{
	Success:    "OK",
	UnknownErr: "Unknown error",

	// User module error
	ErrorUsernameUsed:   "Username has been used",
	ErrorPasswordWrong:  "Password is wrong",
	ErrorUserNotExist:   "User does not exist",
	ErrorTokenExist:     "Token does not exist",
	ErrorTokenRuntime:   "Token has expired",
	ErrorTokenWrong:     "Token is wrong",
	ErrorTokenTypeWrong: "Token format is wrong",
	ErrorUserNoRight:    "User has no right",
	ErrorEmailUsed:      "Email has been used",
	ErrorUsernameEmpty:  "Username is empty",
	ErrorEmailEmpty:     "Email is empty",
	ErrorPasswordEmpty:  "Password is empty",

	// Article module error
	ErrorArticleNotExist: "Article does not exist",

	// Category module error
	ErrorCategoryNameUsed:  "Category name has been used",
	ErrorCategoryNotExist:  "Category does not exist",
	ErrorCategoryNameEmpty: "Category name is empty",

	// Comment module error
	ErrorCommentNotExist: "Comment does not exist",

	// Common error
	ErrorInvalidParam: "Invalid parameter",
}

func GetMsg(code int) string {
	msg, ok := codeMsg[code]
	if ok {
		return msg
	}
	return codeMsg[UnknownErr]
}
