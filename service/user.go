package service

import (
	"reminder/models"
	"time"
)

func InitUserService(userId int64, username string, email string, expiredAt time.Time) {
	models.AddUser(userId, username, 1, expiredAt)
	models.AddNotification(userId, 1, email, "", "", "", "", "")
	InitNewTeamGroupService(userId)
}
