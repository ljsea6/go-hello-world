package pkg

import (
	"github.com/ljsea6/go-hello-world/app/constants"
	"github.com/ljsea6/go-hello-world/app/domain/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constants.ResponseStatus, data T) dto.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
