package controller

import (
	"reminder/models"
	"reminder/response"
	"reminder/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOverviewTeamList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	data := service.GetOverviewTeamService(userIdInt)
	response.Success(c, 200, data)
}

func GetOverviewTeamSearch(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	searchStr := c.PostForm("search_str")
	data := service.SearchOverviewTeamService(userIdInt, searchStr)
	response.Success(c, 200, data)
}

func GetOverviewCalendar(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	year, _ := strconv.Atoi(c.Query("year"))
	month, _ := strconv.Atoi(c.Query("month"))
	result := models.GetOverviewCalendar(userIdInt, year, month)
	response.Success(c, 200, result)
}
