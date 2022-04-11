package models

type NotificationStruct struct {
	DefaultModel
	UserId             int64  `json:"user_id" gorm:"column:user_id"`
	NotificationTarget int    `json:"notification_target" gorm:"notification_target"`
	Email              string `json:"email" gorm:"column:email"`
	Phone              string `json:"phone" gorm:"column:phone"`
	SMS                string `json:"sms" gorm:"column:sms"`
	Telegram           string `json:"telegram" gorm:"column:telegram"`
	Bark               string `json:"bark" gorm:"gorm:bark"`
	ServerChan         string `json:"server_chan" gorm:"column:server_chan"`
}

type RNotificationStruct struct {
	UserId             string `json:"user_id" gorm:"column:user_id"`
	NotificationTarget int    `json:"notification_target" gorm:"notification_target"`
	Email              string `json:"email" gorm:"column:email"`
	Phone              string `json:"phone" gorm:"column:phone"`
	SMS                string `json:"sms" gorm:"column:sms"`
	Telegram           string `json:"telegram" gorm:"column:telegram"`
	Bark               string `json:"bark" gorm:"gorm:bark"`
	ServerChan         string `json:"server_chan" gorm:"column:server_chan"`
}

func AddNotification(userId int64, notificationTarget int, email string, phone string, sms string, telegram string, bark string, serverChan string) {
	notification := NotificationStruct{
		UserId:             userId,
		NotificationTarget: notificationTarget,
		Email:              email,
		Phone:              phone,
		SMS:                sms,
		Telegram:           telegram,
		Bark:               bark,
		ServerChan:         serverChan,
	}
	DB.Table("notification_base").Create(&notification)
}

func GetNotification(userId int64) RNotificationStruct {
	var result RNotificationStruct
	sqlStr := `select * from notification_base where user_id=@userId`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

func UpdateNotification(userId int64, notificationTarget int, email string, phone string, sms string, telegram string, bark string, serverChan string) {
	sqlStr := `update notification_base set notification_target=@notificationTarget,email=@email,phone=@phone,sms=@sms,telegram=@telegram,bark=@bark,server_chan=@serverChan where user_id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{
		"notificationTarget": notificationTarget,
		"email":              email,
		"phone":              phone,
		"sms":                sms,
		"telegram":           telegram,
		"bark":               bark,
		"serverChan":         serverChan,
		"userId":             userId,
	})
}

type RTeamNotificationListItem struct {
	UserId             int64  `json:"user_id" gorm:"column:user_id"`
	NotificationTarget int    `json:"notification_target" gorm:"notification_target"`
	Email              string `json:"email" gorm:"column:email"`
	Bark               string `json:"bark" gorm:"gorm:bark"`
	ServerChan         string `json:"server_chan" gorm:"column:server_chan"`
}

func GetTeamNotificationList(teamId int64) []RTeamNotificationListItem {
	var result []RTeamNotificationListItem
	sqlStr := `select user_id,notification_target,email,bark,server_chan from notification_base where user_id in (select user_id from user_team where team_id=@teamId)`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}
