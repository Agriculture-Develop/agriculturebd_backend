package controller

import (
	"github.com/Agriculture-Develop/agriculturebd/interfaces/vo/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Key
const (
	useridKey = "userId"
	pageKey   = "page"
	countKey  = "count"
)

// Default Value
const (
	DefaultPage  = 1
	DefaultCount = 10
)

// ApiContext Controller definition
type ApiContext[T any] struct {
	Request  *T
	Response *resp.Response
	c        *gin.Context
}

// NewAPiContext  Generic factory function for creating an api controller
func NewAPiContext[T any](c *gin.Context) *ApiContext[T] {
	return &ApiContext[T]{
		Request:  new(T),
		Response: new(resp.Response),
		c:        c,
	}
}

// BindJSON bind json to request
func (ctrl *ApiContext[T]) BindJSON() error {
	return ctrl.c.ShouldBindJSON(ctrl.Request)
}

// BindQuery bind query to request
func (ctrl *ApiContext[T]) BindQuery() error {
	return ctrl.c.ShouldBindQuery(ctrl.Request)
}

// GetUserID get user id from context
func (ctrl *ApiContext[T]) GetUserID() int64 {
	uid, exists := ctrl.c.Get(useridKey)
	if !exists {
		return 0
	}
	if id, ok := uid.(int64); ok {
		return id
	}
	return 0
}

// GetPageAndCount get page and count from query
func (ctrl *ApiContext[T]) GetPageAndCount() (page int, count int, err error) {

	page, err = strconv.Atoi(ctrl.c.DefaultQuery(pageKey, strconv.Itoa(DefaultPage)))
	if err != nil {
		return 0, 0, err
	}
	count, err = strconv.Atoi(ctrl.c.DefaultQuery(countKey, strconv.Itoa(DefaultCount)))
	if err != nil {
		return 0, 0, err
	}

	return page, count, err
}

// NoDataJSON parse with Nodata to json and return
func (ctrl *ApiContext[T]) NoDataJSON(code resp.StatusCode) {
	ctrl.Response.SetNoData(code)
	ctrl.c.JSON(http.StatusOK, ctrl.Response)
}

// WithDataJSON parse with data to json and return
func (ctrl *ApiContext[T]) WithDataJSON(code resp.StatusCode, data interface{}) {
	ctrl.Response.SetWithData(code, data)
	ctrl.c.JSON(http.StatusOK, ctrl.Response)
}
