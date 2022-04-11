package jobs

import (
	"reminder/models"
	"reminder/utils"
)

func PushNotification(data models.RTeamNotificationListItem, goodsName string, message string) {
	if data.NotificationTarget == 1 {
		// 邮件
		// data.NotificationItem.NotificationTarget = data.NotificationItem.Email
		utils.SendMail(data.Email, goodsName, message)
	} else if data.NotificationTarget == 2 {
		// bark
		// data.NotificationItem.NotificationTarget = data.NotificationItem.Bark
		utils.PushBark(data.Bark, goodsName, message)
	} else if data.NotificationTarget == 3 {
		// server酱
		// data.NotificationItem.NotificationTarget = data.NotificationItem.ServerChan
		utils.PushServerChan(data.ServerChan, goodsName, message)
	}
}
