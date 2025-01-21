package handler

import (
	"github.com/gin-gonic/gin"
	"vega-server/api"
	"vega-server/internal/service"
)

type ReviewHandler struct {
	*Handler
	reviewService *service.ReviewService
}

func NewReviewHandler(handler *Handler, reviewService *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		Handler:       handler,
		reviewService: reviewService,
	}
}

// GetReviewFromUser godoc
// @Summary 获取评论信息（用户页面）
// @Description 获取评论信息（用户页面）接口
// @Tags 评论
// @Accept json
// @Produce json
// @Param reviewID query uint true "评论ID"
// @Success 200 {object} api.Response "成功获取评论信息（用户页面）"
// @Failure 400 {object} api.Response "请求参数错误"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /review/user [get]
func (reviewHandler *ReviewHandler) GetReviewFromUser(c *gin.Context) {
	req := &api.GetReviewFromUserRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
	resp, err := reviewHandler.reviewService.GetReviewFromUser(req.ID)
	if err != nil {
		api.HandleError(c, 500, "Internal Server Error", nil)
	}
	api.HandleSuccess(c, resp)
}

// GetReviewToMovie godoc
// @Summary 获取评论信息（电影页面）
// @Description 获取评论信息（电影页面）接口
// @Tags 评论
// @Accept json
// @Produce json
// @Param reviewID query uint true "评论ID"
// @Success 200 {object} api.Response "成功获取评论信息（电影页面）"
// @Failure 400 {object} api.Response "请求参数错误"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /review/movie [get]
func (reviewHandler *ReviewHandler) GetReviewToMovie(c *gin.Context) {
	req := &api.GetReviewToMovieRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
	resp, err := reviewHandler.reviewService.GetReviewToMovie(req.ID)
	if err != nil {
		api.HandleError(c, 500, "Internal Server Error", nil)
	}
	api.HandleSuccess(c, resp)
}
