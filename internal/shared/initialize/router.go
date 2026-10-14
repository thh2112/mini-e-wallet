package initialize

// import (
// 	"context"
// 	"core-ledger-service/internal/shared/global"
// 	"mini-e-wallet/internal/shared/middleware"
// 	"mini-e-wallet/pkg/response"

// 	"command-line-arguments/Users/apple/Desktop/Practice/mini-e-wallet/internal/shared/wire/wire.go"

// 	"github.com/gin-gonic/gin"
// )

// func InitRouter() *gin.Engine {
// 	if global.Config.Server.Mode == "production" {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	engine := gin.New()

// 	// Global Middlewares
// 	engine.Use(gin.Recovery())
// 	engine.Use(gin.Logger())
// 	engine.Use(middleware.RequestID())

// 	engine.GET("/health", func(c *gin.Context) {
// 		response.OK(c, gin.H{"status": "UP"})
// 	})

// 	// Route Groups
// 	v1 := engine.Group("/api/v1")
// 	publicV1 := v1.Group("")
// 	protectedV1 := v1.Group("")
// 	protectedV1.Use(middleware.AuthRequired())

// 	// Dependency Injection
// 	router, err := wire.InitializeRouter(global.DB)
// 	if err != nil {
// 		global.Logger.Fatal(context.Background(), "failed to initialize root router: "+err.Error())
// 	}

// 	// Register all domain routes
// 	router.Register(publicV1, protectedV1)

// 	return engine
// }
