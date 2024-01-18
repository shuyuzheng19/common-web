package middle

import (
	"common-web-framework/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ErrorMiddle(ctx *gin.Context) {
	defer func() {
		var err any = recover()

		if err != nil {
			switch t := err.(type) {
			case common.F:
				ctx.JSON(200, t)
				ctx.Abort()
				break
			default:
				fmt.Println(t)
				ctx.JSON(200, gin.H{"msg": "服务器发生错误"})
				ctx.Abort()
				break
			}
		}
	}()

	ctx.Next()
}
