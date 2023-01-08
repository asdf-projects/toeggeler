package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/toeggeler-server/controllers"
	"github.com/steinm91/toeggeler/toeggeler-server/eval"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
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

	if env.EnableJwt {
		r.Use(JwtAuthMiddleware(env.SecretKey))
	}

	securityRoutes(env, r)
	userRoutes(env, r)
	gameRoutes(env, r)
	evalEngineRoutes(env, r)

	r.Run()
}

func securityRoutes(env *Env, r *gin.Engine) {
	userService := models.UserService{DB: env.DB}
	securityCtrl := controllers.SecurityController{UserService: &userService, SecretKey: env.SecretKey}

	r.POST("/api/authenticate", securityCtrl.Authenticate)
}

func userRoutes(env *Env, r *gin.Engine) {
	userService := models.UserService{DB: env.DB}
	userCtrl := controllers.UserController{UserService: &userService}

	r.GET("/api/users", userCtrl.GetUsers)
	r.GET("/api/users/:id", userCtrl.GetUser)
	r.POST("/api/users", userCtrl.CreateUser)
	r.PUT("/api/users/:id", userCtrl.UpdateUser)
	r.DELETE("/api/users/:id", userCtrl.DeleteUser)
}

func gameRoutes(env *Env, r *gin.Engine) {
	gameService := models.GameService{DB: env.DB}
	gameCtrl := controllers.GameController{GameService: &gameService}

	//r.GET("/api/games", gameCtrl.GetGamesPlayed)
	r.POST("/api/games", gameCtrl.SubmitGame)

	if env.DevMode == true {
		r.POST("/api/games/clear", gameCtrl.ClearGames)
		r.POST("/api/player-stats/clear", gameCtrl.ClearPlayerStats)
	}
}

func evalEngineRoutes(env *Env, r *gin.Engine) {
	gameService := models.GameService{DB: env.DB}
	userService := models.UserService{DB: env.DB}
	evalController := controllers.EvalController{
		GameService: &gameService,
		UserService: &userService,
		EvalEngine:  &env.EvalEngine,
	}

	r.POST("/api/eval", evalController.Eval)
}
