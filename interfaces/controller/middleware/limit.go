package middleware

import (
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/units"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

var (
	bucket   *ratelimit.Bucket
	AuthConf = config.Get().Auth
)

// RateLimitInit 令牌桶限流策略
func RateLimitInit() {
	Interval, _ := units.Duration(AuthConf.RateLimitInterval)
	caps := config.Get().Auth.RateLimitCap
	bucket = ratelimit.NewBucket(Interval, int64(caps))
}

func RateLimitMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//res := new(apiCode.Response)
		//if bucket.TakeAvailable(1) < 1 {
		//	c.JSON(http.StatusOK, res.NoDataResponse(apiCode.CodeVisitLimitExceeded))
		//	c.Abort()
		//	return
		//}
		//c.Next()
	}
}
