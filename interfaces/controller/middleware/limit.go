package middleware

//var (
//	bucket   *ratelimit.Bucket
//	AuthConf = config.Get().Auth
//)
//
//// RateLimitInit 令牌桶限流策略
//func RateLimitInit() {
//	caps := config.Get().Auth.RateLimitCap
//	bucket = ratelimit.NewBucket(units.Duration(AuthConf.RateLimitInterval), int64(caps))
//}
//
//func RateLimitMiddleware() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		res := new(apiCode.Response)
//		if bucket.TakeAvailable(1) < 1 {
//			c.JSON(http.StatusOK, res.NoDataResponse(apiCode.CodeVisitLimitExceeded))
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
