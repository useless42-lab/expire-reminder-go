package models

type NotificationLog struct {
	DefaultModel
	UserId  int64 `json:"user_id" gorm:"column:user_id"`
	GoodsId int64 `json:"goods_id" gorm:"column:goods_id"`
}

func AddNotificationLog(userId int64, goodsId int64) {
	data := NotificationLog{
		UserId:  userId,
		GoodsId: goodsId,
	}
	DB.Table("notification_log").Create(&data)
}

func GetLatestNotificationLog(goodsId int64) NotificationLog {
	var result NotificationLog
	sqlStr := `select * from notification_log where goods_id=@goodsId order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"goodsId": goodsId,
	}).Scan(&result)
	return result
}
