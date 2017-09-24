package giny

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//WriteErrUnauthorized write unauthorize error (401 Unauthorized )
func WriteErrUnauthorized(ctx *gin.Context, err error) {
	if writeLog != nil {
		writeLog(LogTypeErr, fmt.Sprintf("%+v", err))
	}
	ctx.String(http.StatusUnauthorized, err.Error())
}

//WriteErrInternalServerError write internal server error (500 Internal Server Error)
func WriteErrInternalServerError(ctx *gin.Context, err error) {
	if writeLog != nil {
		writeLog(LogTypeErr, fmt.Sprintf("%+v", err))
	}
	ctx.String(http.StatusInternalServerError, err.Error())
}

//WriteErrBadRequest write bad reques error  (400 Bad Reques)
func WriteErrBadRequest(ctx *gin.Context, err error) {
	if writeLog != nil {
		writeLog(LogTypeErr, fmt.Sprintf("%+v\n", err))
	}
	ctx.String(http.StatusBadRequest, err.Error())
}

//WriteOK ok
func WriteOK(ctx *gin.Context, a interface{}) {
	ctx.JSON(http.StatusOK, a)
}
