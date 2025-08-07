package middleware

import (
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/common/bizcode"
	"github.com/Agriculture-Develop/agriculturebd/interfaces/vo/resp"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ZapRecovery 使用zap库表进行panic错误的封装
func ZapRecovery(stack bool) gin.HandlerFunc {
	return ginzap.CustomRecoveryWithZap(zap.L(), stack, func(c *gin.Context, err interface{}) {
		c.JSON(200, resp.Response{
			StatusCode: bizcode.ServerBusy,
			StatusMsg:  resp.Error(bizcode.ServerBusy, nil).Error(),
		})
	})
}
