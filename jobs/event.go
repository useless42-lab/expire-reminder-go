package jobs

import (
	"reminder/models"
	"time"
)

type EventJob struct{}

func (eventJob EventJob) Run() {
	result := models.GetAllGoodsList()
	for _, item := range result {
		if time.Now().After(item.ExpireTime) {
			// 已经过期
			models.SetGoodsExpire(item.Id)
			notificationList := models.GetTeamNotificationList(item.TeamId)
			for _, item1 := range notificationList {
				go PushNotification(item1, item.Name, "已经过期！")
				models.AddNotificationLog(item1.UserId, item.Id)
			}
		} else {
			notificationLog := models.GetLatestNotificationLog(item.Id)
			// 每隔七天提醒
			if time.Now().After(notificationLog.CreatedAt.Add(+time.Hour * time.Duration(24*7))) {
				targetTime := item.ExpireTime.Format("2006-01-02 15:04:05")
				targetTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", targetTime, time.Local)
				targetTimeLocation = targetTimeLocation.Add(-time.Hour * time.Duration(item.AdvanceDay*24))
				if time.Now().After(targetTimeLocation) {
					// 过期提醒
					notificationList := models.GetTeamNotificationList(item.TeamId)
					for _, item1 := range notificationList {
						go PushNotification(item1, item.Name, "即将过期！")
						models.AddNotificationLog(item1.UserId, item.Id)
					}
				}
			}
		}
	}
}
