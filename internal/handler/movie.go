package handler

import (
	"github.com/gin-gonic/gin"
	"vega-server/api"
	"vega-server/internal/service"
)

type MovieHandler struct {
	*Handler
	movieService *service.MovieService
}

func NewMovieHandler(handler *Handler, movieService *service.MovieService) *MovieHandler {
	return &MovieHandler{
		Handler:      handler,
		movieService: movieService,
	}
}

// GetMovie godoc
// @Summary 获取电影信息
// @Description 获取电影信息接口
// @Tags 电影
// @Accept json
// @Produce json
// @Param movieID query uint true "电影ID"
// @Success 200 {object} api.Response "成功获取电影信息"
// @Failure 400 {object} api.Response "请求参数错误"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /movie [get]
func (movieHandler *MovieHandler) GetMovie(c *gin.Context) {
	req := &api.GetMovieRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
	resp, err := movieHandler.movieService.GetMovie(req.ID)
	if err != nil {
		api.HandleError(c, 500, "Internal Server Error", nil)
		return
	}
	api.HandleSuccess(c, resp)
	return
}

// GetAdvancedMovie godoc
// @Summary 获取电影信息和评论列表
// @Description 获取电影信息和评论列表接口
// @Tags 电影
// @Accept json
// @Produce json
// @Param movieID query uint true "电影ID"
// @Success 200 {object} api.Response "成功获取电影信息和评论列表"
// @Failure 400 {object} api.Response "请求参数错误"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /movie/reviews [get]
func (movieHandler *MovieHandler) GetAdvancedMovie(c *gin.Context) {
	req := &api.GetAdvancedMovieRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
	resp, err := movieHandler.movieService.GetAdvancedMovie(req.ID)
	if err != nil {
		api.HandleError(c, 500, "Internal Server Error", nil)
		return
	}
	api.HandleSuccess(c, resp)
	return
}

// GetMovies godoc
// @Summary 获取电影类型榜单
// @Description 获取电影类型榜单接口
// @Tags 电影
// @Accept json
// @Produce json
// @Param genre query string true "Genre"
// @Param category query string true "Category"
// @Param field query string false "Field"
// @Success 200 {object} api.Response "成功获取电影类型榜单"
// @Failure 400 {object} api.Response "请求参数错误"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /movies [get]
func (movieHandler *MovieHandler) GetMovies(c *gin.Context) {
	req := &api.GetMoviesRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
	switch req.Field {
	case "":
		resp, err := movieHandler.movieService.GetIDs(req.Genre, req.Category)
		if err != nil {
			api.HandleError(c, 500, "Internal Server Error", nil)
			return
		}
		api.HandleSuccess(c, resp)
		return
	case "title":
		resp, err := movieHandler.movieService.GetTitles(req.Genre, req.Category)
		if err != nil {
			api.HandleError(c, 500, "Internal Server Error", nil)
			return
		}
		api.HandleSuccess(c, resp)
		return
	default:
		api.HandleError(c, 400, "Bad Request", nil)
		return
	}
}

// GetNowPlaying godoc
// @Summary 获取正在热映电影列表
// @Description 获取正在热映电影列表接口
// @Tags 电影
// @Accept json
// @Produce json
// @Success 200 {object} api.Response "成功获取正在热映电影列表"
// @Failure 500 {object} api.Response "服务器内部错误"
// @Router /movies/cinema/now-playing [get]
func (movieHandler *MovieHandler) GetNowPlaying(c *gin.Context) {
	resp, err := movieHandler.movieService.GetNowPlaying()
	if err != nil {
		api.HandleError(c, 500, "Internal Server Error", nil)
	}
	api.HandleSuccess(c, resp)
}
