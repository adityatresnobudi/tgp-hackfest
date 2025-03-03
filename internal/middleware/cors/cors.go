package cors

import (
	"fmt"

	"github.com/dinata1312/TechGP-Project/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(config config.Config) gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowMethods:     []string{"POST", "PUT", "GET", "PATCH", "DELETE"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	}

	if config.App.OriginDomain == "localhost" {
		origin := "http://localhost:8000"
		corsConfig.AllowOrigins = []string{origin}
	} else {
		origin := fmt.Sprintf("https://%s", config.App.OriginDomain)
		corsConfig.AllowOrigins = []string{origin}
	}

	return cors.New(corsConfig)
}
