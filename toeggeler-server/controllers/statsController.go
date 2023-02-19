package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/toeggeler-server/eval"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

type StatsController struct {
	GameService *models.GameService
	UserService *models.UserService
	EvalEngine  *eval.EvalEngine
}

// GetStats godoc
// @Summary      Get statistics for all available users
// @Description  Get statistics for all available users
// @Tags		 Stats
// @Accept       json
// @Produce      json
// @Success      200  {object}  []Stats
// @Router       /stats [get]
func (ctrl StatsController) GetStats(c *gin.Context) {
	stats := ctrl.EvalEngine.GetStats()
	c.JSON(http.StatusOK, stats)
}

// GetStatsForUser godoc
// @Summary      Get statistics for a user
// @Description  Get statistics for a user
// @Tags		 Stats
// @Accept       json
// @Produce      json
// @Param		 id path int true "User ID"
// @Success      200  {object} Stats
// @Failure      404
// @Router       /stats/{id} [get]
func (ctrl StatsController) GetStatsForPlayer(c *gin.Context) {
	stats := ctrl.EvalEngine.GetStats()

	var statsForPlayer eval.Stats
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	for i := range stats {
		if stats[i].PlayerId == id {
			statsForPlayer = stats[i]
		}
	}

	c.JSON(http.StatusOK, statsForPlayer)
}
