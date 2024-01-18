package common

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type F struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const okStatusCode = 200

const errorStatusCode = 500

type ErrorCode int

const (
	BadRequestCode ErrorCode = 400
	FailCode       ErrorCode = iota + 10001
	RegisteredCode
	SendEmailFailCode
	EmailCodeValidate
	CreateTokenFail
	ParseTokenFail
	LoginFail
	NoLogin
	RoleAuthenticationFail
	AddFileFail
	NoFile
)

var messageMap = map[ErrorCode]string{
	FailCode:               "处理失败",
	BadRequestCode:         "参数验证失败",
	RegisteredCode:         "注册失败",
	SendEmailFailCode:      "发送邮件失败",
	EmailCodeValidate:      "邮件验证码不正确",
	CreateTokenFail:        "创建Token失败",
	ParseTokenFail:         "解析Token失败",
	LoginFail:              "账号或密码错误",
	NoLogin:                "你还未登录",
	RoleAuthenticationFail: "你没有权限访问",
	AddFileFail:            "添加文件失败",
	NoFile:                 "没有文件,请选择文件",
}

func OK() R {
	return Success(nil)
}

func Success(data interface{}) R {
	return R{Code: okStatusCode, Message: "success", Data: data}
}

func Error() F {
	return F{Code: errorStatusCode, Message: "error"}
}

func Fail(code ErrorCode, message string) F {
	return F{Code: int(code), Message: message}
}

func AutoFail(code ErrorCode) F {
	return F{Code: int(code), Message: messageMap[code]}
}

func FieldError(message string) F {
	return F{Code: int(BadRequestCode), Message: message}
}
