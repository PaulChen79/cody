package domain

type ErrorFormat struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Server Error"`
	Time    string `json:"time" example:"1638779090"`
}

var (
	ErrorBadRequest        = ErrorFormat{Code: 400, Message: "bad request"}
	ErrorUnauthorized      = ErrorFormat{Code: 401, Message: "Unauthorized"}
	ErrorInvalidParameter  = ErrorFormat{Code: 402, Message: "invalid parameter"}
	ErrorAuthTokenExpired  = ErrorFormat{Code: 4011, Message: "auth token expired"}
	ErrorInvalidToken      = ErrorFormat{Code: 4012, Message: "invalid token"}
	ErrorForbidden         = ErrorFormat{Code: 403, Message: "Forbidden"}
	ErrorUnknowServerError = ErrorFormat{Code: 405, Message: "Unknow Internal Error"}
	ErrorServer            = ErrorFormat{Code: 500, Message: "Server error, please retry again later"}
	ErrorDBOperationError  = ErrorFormat{Code: 501, Message: "DB operation error, please retry again later"}
)
