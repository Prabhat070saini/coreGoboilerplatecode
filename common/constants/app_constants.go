package constants



type contextKey string
const RequestIDKey       contextKey = "x-request-id"
const (

	RequestStartTimeKey contextKey = "request_start_time"
)
const (
	ForgotPasswordRedisKey    = "auth:forgotPassword:%s"
	LoginAccessTokenRedisKey  = "auth:user:accessToken:%s"
	LoginRefreshTokenRedisKey = "auth:user:refreshToken:%s"
	VerifyAccessTokenRedisKey = "auth:user:ids:%s"
)


// custom error codes
const (
	InvalidAccessToken        int = 601 // take refreshToken
	InvalidTokenLogout        int = 602 //logout
	ResetTokenInvalidOrExpire int = 603 //resetToken
)
