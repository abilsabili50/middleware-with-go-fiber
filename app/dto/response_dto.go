package dto

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    T      `json:"data"`
}

func (r *Response[T]) Mapper(status, message string, code int, data T) {
	r.Status = status
	r.Code = code
	r.Message = message
	r.Data = data
}
