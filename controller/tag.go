package controller

import (
	"reminder/models"
	"reminder/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTagList(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	result := models.GetTagList(teamId, groupId)
	response.Success(c, 200, result)
}
