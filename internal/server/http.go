package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "vega-server/docs"
	"vega-server/internal/handler"
	"vega-server/internal/middleware"
	"vega-server/pkg/config"
	"vega-server/pkg/jwt"
	"vega-server/pkg/log"
	"vega-server/pkg/server/http"
)

func NewHTTPServer(
	config *config.Config,
	logger *log.Logger,
	jwtService *jwt.JWTService,
	userHandler *handler.UserHandler,
	movieHandler *handler.MovieHandler,
	reviewHandler *handler.ReviewHandler,
) *http.Server {
	// Server
	gin.SetMode(gin.DebugMode)
	server := http.NewServer(config)

	// Swagger
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Middleware
	server.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://example.com"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
			MaxAge:           24 * 60 * 60,
		}),
		middleware.RequestIDMiddleware(),
		middleware.LoggerMiddleware(logger),
	)

	// Routes
	v1 := server.Group("/api")
	{
		public := v1.Group("/")
		{
			public.POST("/sign-up", userHandler.SignUp)
			public.POST("/sign-in", userHandler.SignIn)
			public.GET("/movie", movieHandler.GetMovie)
			public.GET("/movie/reviews", movieHandler.GetAdvancedMovie)
			public.GET("/movies", movieHandler.GetMovies)
			public.GET("/movies/cinema/now-playing", movieHandler.GetHotMovies)
			public.GET("/review/user", reviewHandler.GetReviewFromUser)
			public.GET("/review/movie", reviewHandler.GetReviewToMovie)
		}
		user := v1.Group("/")
		user.Use(middleware.UserAuthMiddleware(jwtService, logger))
		{
			user.GET("/user", userHandler.GetUser)
			user.GET("/user/reviews", userHandler.GetAdvancedUser)
			user.GET("/user/review/:movieID", userHandler.GetReview)
			user.POST("/user/review/:movieID", userHandler.SetReview)
		}
	}

	return server
}
