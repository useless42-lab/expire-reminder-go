package controller

import (
	"reminder/models"
	"reminder/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateNotification(c *gin.Context) {
	userId := c.GetInt64("userId")
	notificationTarget, _ := strconv.Atoi(c.PostForm("notification_target"))
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	sms := c.PostForm("sms")
	telegram := c.PostForm("telegram")
	bark := c.PostForm("bark")
	serverChan := c.PostForm("server_chan")
	models.UpdateNotification(userId, notificationTarget, email, phone, sms, telegram, bark, serverChan)
	response.Success(c, 200, "")
}

func GetNotification(c *gin.Context) {
	userId := c.GetInt64("userId")
	result := models.GetNotification(userId)
	response.Success(c, 200, result)
}
