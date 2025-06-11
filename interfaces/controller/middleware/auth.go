package middleware

import (
	"errors"
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/domain/auth/model/valobj"
	"github.com/Agriculture-Develop/agriculturebd/domain/common/respCode"
	utils "github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/jwt"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/controller"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"strings"
)

// 验证用户是否登录的中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := controller.NewAPiContext[struct{}](c)

		// 读取验证token
		tokenString, ok := getToken(c)
		if !ok {
			ctx.NoDataJSON(respCode.InvalidTokenForm)
			c.Abort()
			return
		}

		// 校验token信息
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenMalformed) {
				ctx.NoDataJSON(respCode.InvalidTokenForm)
				c.Abort()
				return
			}

			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.NoDataJSON(respCode.InvalidTokenExpired)
				c.Abort()
				return
			}

			ctx.NoDataJSON(respCode.InvalidToken)
			c.Abort()
			return
		}

		// 存储用户信息
		c.Set("userId", claims.ID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func WithAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := controller.NewAPiContext[struct{}](c)

		if ctx.GetUserIdByRole() != valobj.RoleUser.Int() {
			ctx.NoDataJSON(respCode.Forbidden)
			c.Abort()
			return
		}

	}
}

func WithSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := controller.NewAPiContext[struct{}](c)

		fmt.Println(ctx.GetUserIdByRole())
		if ctx.GetUserIdByRole() != valobj.RoleSuperAdmin.Int() {
			ctx.NoDataJSON(respCode.Forbidden)
			c.Abort()
			return
		}

	}
}

func getToken(c *gin.Context) (string, bool) {
	tokenString := c.GetHeader("Authorization")

	if !strings.HasPrefix(tokenString, "Bearer ") {
		return "", false
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	return tokenString, true
}
