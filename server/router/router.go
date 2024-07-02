package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/internal/user"
	"server/internal/ws"
	"server/middleware"

	"time"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	authRouter := r.Group("/ws")
	authRouter.Use(middleware.JWTAuth)
	authRouter.POST("/createRoom", wsHandler.CreateRoom)
	authRouter.GET("/joinRoom/:roomId", wsHandler.JoinRoom)
	authRouter.GET("/getRooms", wsHandler.GetRooms)
	authRouter.GET("/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
