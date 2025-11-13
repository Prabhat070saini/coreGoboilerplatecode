package constants
// type Exception struct {
// 	Code           int    `json:"code"`
// 	Message        string `json:"message"`
// 	HttpStatusCode int    `json:"httpStatusCode"`
// }

// // IServiceOutput<T> equivalent
// type ServiceOutput[T any] struct {
// 	Message        string     `json:"message,omitempty"`
// 	OutputData     T          `json:"outputData,omitempty"`
// 	Exception      *Exception `json:"exception,omitempty"`
// 	HttpStatusCode int        `json:"httpStatusCode"`
// 	RespStatusCode int        `json:"respStatusCode"`
// }

// // Final API response structure

type ApiResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
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
