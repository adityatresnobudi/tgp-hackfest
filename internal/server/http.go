package server

import (
	"strings"

	"github.com/dinata1312/TechGP-Project/internal/middleware/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func applicationJsonResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.Path, "/swagger/") {
			c.Writer.Header().Set("Content-Type", "application/json")
		}

		c.Next()
	}
}

func (s *server) runGinServer() error {
	s.r.Use(applicationJsonResponseMiddleware())
	s.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	s.r.Use(cors.CORS(s.cfg))
	return s.r.Run(s.cfg.Http.Port)
}
