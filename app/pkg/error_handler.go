package pkg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ljsea6/go-hello-world/app/constants"
)

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		switch key {
		case
			constants.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			c.Abort()
		case
			constants.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
			c.Abort()
		}
	}
}
