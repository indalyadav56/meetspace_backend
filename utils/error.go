package utils

import "github.com/gin-gonic/gin"


func HandleError(c *gin.Context, err error) {
    resp := ErrorResponse("Error", err.Error())
    c.JSON(resp.StatusCode, resp)
}