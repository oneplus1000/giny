package giny

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//LogTypeErr  error
const LogTypeErr = 1

//LogTypeWarn warn
const LogTypeWarn = 2

//LogTypeInfo info
const LogTypeInfo = 3

//LogTypeDebug debuf
const LogTypeDebug = 4

var writeLog func(logtype int, msg string)

//SetWriteLog set logger
func SetWriteLog(wl func(logtype int, msg string)) {
	writeLog = wl
}

//WriteErrUnauthorized write unauthorize error
func WriteErrUnauthorized(ctx *gin.Context, err error) {
	if writeLog != nil {
		writeLog(LogTypeErr, fmt.Sprintf("%s\n", err.Error()))
	}
	ctx.String(http.StatusUnauthorized, err.Error())
}

//WriteErrInternalServerError write internal server error
func WriteErrInternalServerError(ctx *gin.Context, err error) {
	if writeLog != nil {
		writeLog(LogTypeErr, fmt.Sprintf("%s\n", err.Error()))
	}
	ctx.String(http.StatusInternalServerError, err.Error())
}

//WriteOK ok
func WriteOK(ctx *gin.Context, a interface{}) {
	ctx.JSON(http.StatusOK, a)
}