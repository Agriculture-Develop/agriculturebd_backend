package middleware

import (
	"errors"
	"net/http"
	"strings"

	"dream_program/pkg/auth/role"
	"dream_program/pkg/auth/token"
	"dream_program/types/apiCode"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Parse 宽松认证
func Parse() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := getToken(c)

		// 解析并校验Token
		claims, _ := token.ParseToken(tokenString)
		if claims.UserId != 0 {
			c.Set("userId", claims.UserId)
		}
		c.Next()
	}
}

// 验证用户是否登录的中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := new(apiCode.Response)

		// 读取验证token
		tokenString, ok := getToken(c)
		if !ok {
			c.JSON(http.StatusOK, res.NoDataResponse(apiCode.CodeInvalidToken))
			c.Abort()
			return
		}

		// 校验token信息
		claims, err := token.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenMalformed) {
				c.JSON(http.StatusOK, res.NoDataResponse(apiCode.CodeInvalidTokenForm))
				c.Abort()
				return
			}

			if errors.Is(err, jwt.ErrTokenExpired) && claims.TokenType == 0 {
				// 提示需要刷新token
				c.JSON(http.StatusOK, res.NoDataResponse(apiCode.CodeInvalidTokenExpired))
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, res.NoDataResponse(apiCode.CodeInvalidToken))
			c.Abort()
			return
		}

		// 认证用户角色权限
		StatusCode := role.CheckAdmin(c, claims.UserId)
		if StatusCode != apiCode.CodeSuccess {
			c.JSON(http.StatusUnauthorized, res.NoDataResponse(StatusCode))
			c.Abort()
			return
		}

		// 存储用户信息
		c.Set("userId", claims.UserId)
		c.Next()
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
