package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/toeggeler-server/eval"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

type EvalController struct {
	GameService *models.GameService
	UserService *models.UserService
	EvalEngine  *eval.EvalEngine
}

func (e EvalController) Eval(c *gin.Context) {
	e.EvalEngine.Eval()
}
