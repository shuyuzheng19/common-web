package middle

import (
	"common-web-framework/common"
	"common-web-framework/models"
	"common-web-framework/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RoleName string

const (
	UserRole  RoleName = "USER"
	AdminRole RoleName = "ADMIN"
	SuperRole RoleName = "SUPER_ADMIN"
)

const tokenType = "Bearer "

const tokenHeader = "Authorization"

func JwtMiddle(roleName RoleName, f func(id int) *models.User) gin.HandlerFunc {
	return func(context *gin.Context) {
		var header = context.GetHeader(tokenHeader)

		if header == "" || !strings.HasPrefix(header, tokenType) {
			context.JSON(http.StatusOK, common.AutoFail(common.NoLogin))
			context.Abort()
			return
		}

		var token = strings.Replace(header, tokenType, "", 1)

		fmt.Println(token)

		var uid = utils.ParseTokenToUserId(token)

		if uid == -1 {
			context.JSON(http.StatusOK, common.AutoFail(common.ParseTokenFail))
			context.Abort()
			return
		}

		var user = f(uid)

		if user == nil {
			context.JSON(http.StatusOK, common.Fail(500, "该用户找不到了...."))
			context.Abort()
			return
		}

		var role = user.Role.Name

		var isAuth = false

		if roleName == UserRole || role == string(SuperRole) {
			isAuth = true
		} else if roleName == AdminRole && role == string(AdminRole) {
			isAuth = true
		}

		if isAuth {
			context.Set("user", *user)
			context.Next()
		} else {
			context.JSON(http.StatusOK, common.AutoFail(common.RoleAuthenticationFail))
			context.Abort()
			return
		}
	}
}
