package response


// IException equivalent
type Exception struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	HttpStatusCode int    `json:"httpStatusCode,omitempty"`
}

// ISuccess<T> equivalent
type Success[T any] struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	HttpStatusCode int    `json:"httpStatusCode"`
	Data           T     `json:"data,omitempty"`
}

// IServiceOutput<T> equivalent
type ServiceOutput[T any] struct {
	Success   *Success[T]  `json:"success,omitempty"`
	Exception *Exception   `json:"exception,omitempty"`
}

// IFunctionOutput<T> equivalent
type FunctionOutput[T any] struct {
	Data      T          `json:"data,omitempty"`
	Exception *Exception  `json:"exception,omitempty"`
}

type ApiResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}