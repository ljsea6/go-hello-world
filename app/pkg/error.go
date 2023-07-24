package pkg

import (
	"errors"
	"fmt"

	"github.com/ljsea6/go-hello-world/app/constants"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constants.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}
