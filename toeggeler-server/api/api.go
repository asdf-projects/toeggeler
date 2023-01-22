package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/toeggeler-server/controllers"
	_ "github.com/steinm91/toeggeler/toeggeler-server/docs"
	"github.com/steinm91/toeggeler/toeggeler-server/eval"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Env struct {
	DevMode    bool
	DB         *sql.DB
	Port       int64
	EvalEngine eval.EvalEngine
	SecretKey  string
	EnableJwt  bool
}

func StartApiServer(env *Env) {
	if env.DevMode == true {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	unauthorized := r.Group("/api")
	swaggerRoutes(env, unauthorized)
	securityRoutes(env, unauthorized)

	authorized := r.Group("/api")
	if env.EnableJwt {
		authorized.Use(JwtAuthMiddleware(env.SecretKey))
	}

	userRoutes(env, unauthorized, authorized)
	gameRoutes(env, unauthorized, authorized)
	statsRoutes(env, unauthorized)
	evalEngineRoutes(env, unauthorized)

	r.Run()
}

func swaggerRoutes(env *Env, unauthorized *gin.RouterGroup) {
	unauthorized.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func securityRoutes(env *Env, unauthorized *gin.RouterGroup) {
	userService := models.UserService{DB: env.DB}
	securityCtrl := controllers.SecurityController{UserService: &userService, SecretKey: env.SecretKey}

	unauthorized.POST("/authenticate", securityCtrl.Authenticate)
}

func userRoutes(env *Env, unauthorized *gin.RouterGroup, authorized *gin.RouterGroup) {
	userService := models.UserService{DB: env.DB}
	userCtrl := controllers.UserController{UserService: &userService}

	unauthorized.GET("/users", userCtrl.GetUsers)
	unauthorized.GET("/users/:id", userCtrl.GetUser)
	unauthorized.POST("/users", userCtrl.CreateUser)

	authorized.PUT("/users/:id", userCtrl.UpdateUser)
	authorized.DELETE("/users/:id", userCtrl.DeleteUser)
}

func gameRoutes(env *Env, unauthorized *gin.RouterGroup, authorized *gin.RouterGroup) {
	gameService := models.GameService{DB: env.DB}
	gameCtrl := controllers.GameController{GameService: &gameService}

	unauthorized.GET("/games", gameCtrl.GetGamesPlayed)
	authorized.POST("/games", gameCtrl.SubmitGame)

	if env.DevMode == true {
		unauthorized.POST("/games/clear", gameCtrl.ClearGames)
		unauthorized.POST("/player-stats/clear", gameCtrl.ClearPlayerStats)
	}
}

func statsRoutes(env *Env, unauthorized *gin.RouterGroup) {
	statsCtrl := controllers.StatsController{}
	unauthorized.GET("/stats", statsCtrl.GetStats)
	unauthorized.GET("/stats/:id", statsCtrl.GetStatsForPlayer)
}

func evalEngineRoutes(env *Env, unauthorized *gin.RouterGroup) {
	gameService := models.GameService{DB: env.DB}
	userService := models.UserService{DB: env.DB}
	evalController := controllers.EvalController{
		GameService: &gameService,
		UserService: &userService,
		EvalEngine:  &env.EvalEngine,
	}

	unauthorized.POST("/eval", evalController.Eval)
}
