package presentation

import (
	"gingorm1/business"
	"github.com/gin-gonic/gin"
)

func WriteRestOutput[T any](c *gin.Context, out T, cr *business.CommonResponse) {
	if cr.ErrorCode == 0 {
		cr.ErrorCode = 200
	}
	if cr.SetAuthToken != `` {
		c.SetCookie(`authToken`, cr.SetAuthToken, 3600, ``, ``, false, false)
	}
	c.JSON(cr.ErrorCode, out)
}

func ReadRestInput[T any](c *gin.Context, in *T, cr *business.CommonRequest) {
	err := c.BindJSON(&in)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// TODO: read cookie put to cr
}
