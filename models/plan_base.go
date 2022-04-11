package models

type PlanBase struct {
	Id                int     `json:"id" gorm:"column:id"`
	Name              string  `json:"name" gorm:"column:name"`
	PriceYear         float64 `json:"price_year" gorm:"column:price_year"`
	TeamNumber        int     `json:"team_number" gorm:"team_number"`
	PerTeamGroupLimit int     `json:"per_team_group_limit" gorm:"column:per_team_group_limit"`
	PerTeamGoodsLimit int     `json:"per_team_goods_limit" gorm:"per_team_goods_limit"`
	TeamMemberLimit   int     `json:"team_member_limit" gorm:"column:team_member_limit"`
	Status            int     `json:"status" gorm:"column:status"`
	UrlKeyMonth       string  `json:"url_key_month" gorm:"column:url_key_month"`
	UrlKeyYear        string  `json:"url_key_year" gorm:"column:url_key_year"`
}

func GetPlanBaseList() []PlanBase {
	var result []PlanBase
	sqlStr := `select * from plan_base`
	DB.Raw(sqlStr).Scan(&result)
	return result
}

func GetPlanBaseInfo(planId int) PlanBase {
	var result PlanBase
	sqlStr := `select * from plan_base where id=@planId`
	DB.Raw(sqlStr, map[string]interface{}{"planId": planId}).Scan(&result)
	return result
}

type RUserPlanId struct {
	Id int `json:"plan_id" gorm:"column:plan_id"`
}

func GetPlanIdByTeamId(teamId int64) RUserPlanId {
	var result RUserPlanId
	sqlStr := `select plan_id from team_group as tg left join user on user.id=tg.user_id where tg.id=@teamId`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

func GetPlanIdByUserId(userId int64) RUserPlanId {
	var result RUserPlanId
	sqlStr := `select plan_id from user where id=@userId`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

func GetPlanBaseInfoByDeviceId(deviceId int64, deviceType int) PlanBase {
	var result PlanBase
	sqlStr := `
SELECT
	*
FROM
	plan_base AS pb
inner JOIN (
	SELECT
		plan_id
	FROM
		user AS u
	INNER JOIN (
		SELECT
			tg.user_id
		FROM
			(
				SELECT
					dg.team_id
				FROM
					web_list AS wl
				INNER JOIN device_group AS dg ON dg.id = wl.group_id
				WHERE
					dg.device_type = @deviceType
				AND wl.id = @deviceId
			) AS a
		INNER JOIN team_group AS tg ON tg.id = a.team_id
	) AS b ON b.user_id = u.id
) AS c ON pb.id = c.plan_id
`
	DB.Raw(sqlStr, map[string]interface{}{
		"deviceId":   deviceId,
		"deviceType": deviceType,
	}).Scan(&result)
	return result
}
