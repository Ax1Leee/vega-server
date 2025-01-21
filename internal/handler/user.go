package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"vega-server/api"
	"vega-server/internal/service"

	"errors"
)

type UserHandler struct {
	*Handler
	userService   *service.UserService
	reviewService *service.ReviewService
}

func NewUserHandler(handler *Handler, userService *service.UserService, reviewService *service.ReviewService) *UserHandler {
	return &UserHandler{
		Handler:       handler,
		userService:   userService,
		reviewService: reviewService,
	}
}

// SignUp godoc
// @Summary 用户注册
// @Description 用户注册接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body api.SignUpRequest true "用户信息"
// @Success 200 {object} api.Response "注册成功"
// @Failure 400 {object} api.Response "请求参数错误"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /sign-up [post]
func (userHandler *UserHandler) SignUp(c *gin.Context) {
	req := &api.SignUpRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
	resp, err := userHandler.userService.SignUp(req)
	if err != nil {
		if errors.Is(err, errors.New("email already in use")) {
			api.HandleError(c, 400, "Email already in use", nil)
			return
		} else {
			api.HandleError(c, 500, "Internal Server Error", nil)
			return
		}
	}
	api.HandleSuccess(c, resp)
	return
}

// SignIn godoc
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param request body api.SignInRequest true "登录信息"
// @Success 200 {object} api.Response "登录成功"
// @Failure 400 {object} api.Response "请求参数错误"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /sign-in [post]
func (userHandler *UserHandler) SignIn(c *gin.Context) {
	req := &api.SignInRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
	resp, err := userHandler.userService.SignIn(req)
	if err != nil {
		if errors.Is(err, errors.New("invalid email")) || errors.Is(err, errors.New("invalid password")) {
			api.HandleError(c, 400, "Invalid email or password", nil)
			return
		} else {
			api.HandleError(c, 500, "Internal Server Error", nil)
			return
		}
	}
	api.HandleSuccess(c, resp)
	return
}

// GetUser godoc
// @Summary 获取用户信息
// @Description 获取用户信息接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} api.Response "成功获取用户信息"
// @Failure 401 {object} api.Response "未授权"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /user [get]
func (userHandler *UserHandler) GetUser(c *gin.Context) {
	id := c.MustGet("id").(uint)
	resp, err := userHandler.userService.GetUser(id)
	if err != nil {
		if errors.Is(err, errors.New("invalid id")) {
			api.HandleError(c, 401, "Unauthorized", nil)
			return
		} else {
			api.HandleError(c, 500, "Internal Server Error", nil)
			return
		}
	}
	api.HandleSuccess(c, resp)
	return
}

// GetAdvancedUser godoc
// @Summary 获取用户信息和评论列表
// @Description 获取用户信息和评论列表接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} api.Response "成功获取用户信息和评论列表"
// @Failure 401 {object} api.Response "未授权"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /user/reviews [get]
func (userHandler *UserHandler) GetAdvancedUser(c *gin.Context) {
	id := c.MustGet("id").(uint)
	resp, err := userHandler.userService.GetAdvancedUser(id)
	if err != nil {
		if errors.Is(err, errors.New("invalid id")) {
			api.HandleError(c, 401, "Unauthorized", nil)
			return
		} else {
			api.HandleError(c, 500, "Internal Server Error", nil)
			return
		}
	}
	api.HandleSuccess(c, resp)
	return
}

// GetReview godoc
// @Summary 获取用户对电影的评论信息
// @Description 获取用户对电影的评论信息接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param movieID query uint true "电影ID"
// @Success 200 {object} api.Response "成功获取用户对电影的评论信息"
// @Failure 400 {object} api.Response "请求参数错误"
// @Failure 401 {object} api.Response "未授权"
// @Failure 404 {object} api.Response "未找到资源"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /user/review [get]
func (userHandler *UserHandler) GetReview(c *gin.Context) {
	userID := c.MustGet("id").(uint)
	req := &api.GetReviewRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
	resp, err := userHandler.reviewService.GetReview(userID, req.MovieID)
	if err != nil {
		if errors.Is(err, errors.New("review not found")) {
			api.HandleError(c, 404, "Not Found", nil)
			return
		} else {
			api.HandleError(c, 500, "Internal Server Error", nil)
			return
		}
	}
	api.HandleSuccess(c, resp)
	return
}

// SetReview godoc
// @Summary 创建或更新用户对电影的评论信息
// @Description 创建或更新用户对电影的评论信息接口
// @Tags 用户
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param movieID path uint true "电影ID"
// @Param request body api.SetReviewRequest true "评论信息"
// @Success 200 {object} api.Response "成功获取用户对电影的评论信息"
// @Failure 400 {object} api.Response "请求参数错误"
// @Failure 401 {object} api.Response "未授权"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /user/review [post]
func (userHandler *UserHandler) SetReview(c *gin.Context) {
	userID := c.MustGet("id").(uint)
	movieID, err := strconv.ParseUint(c.Param("movieID"), 10, 64)
	if err != nil {
		api.HandleError(c, 500, "Internal Server Error", nil)
		return
	}
	req := &api.SetReviewRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
	resp, err := userHandler.reviewService.SetReview(userID, uint(movieID), req)
	if err != nil {
		api.HandleError(c, 500, "Internal Server Error", nil)
		return
	}
	api.HandleSuccess(c, resp)
	return
}
