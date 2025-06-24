package auth

import (
	"github.com/Agriculture-Develop/agriculturebd/api/routes/Interface"
	"github.com/gin-gonic/gin"
)

func AuthModels(auth *gin.RouterGroup, authCtrl Interface.IAuthCtrl) {
	auth.POST("/login/pwd", authCtrl.LoginByPassword) // 密码登录
	auth.POST("/login/code", authCtrl.LoginByCode)    // 验证码登录
	auth.POST("/register", authCtrl.Register)         // 注册
	auth.POST("/code", authCtrl.SendPhoneCode)        // 发送验证码

	auth.PUT("password", authCtrl.UpdateUserPassword) // 更新用户密码
}
