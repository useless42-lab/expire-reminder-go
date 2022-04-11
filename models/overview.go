package models

type RGoodsOverviewNumber struct {
	All     int `json:"all"`
	Active  int `json:"active"`
	Expired int `json:"expired"`
	Thrown  int `json:"thrown"`
}

type OverviewTeamStruct struct {
	Id                  string               `json:"id" grom:"column:id"`
	Name                string               `json:"name" gorm:"column:"name"`
	Username            string               `json:"username" gorm:"column:username"`
	Role                int                  `json:"role" gorm:"column:role"`
	CreatedAt           LocalTime            `json:"created_at" gorm:"column:created_at" type:"datetime"`
	GoodsOverviewNumber RGoodsOverviewNumber `json:"goods_overview_number" gorm:"column:-"`
}

func GetOverviewTeam(userId int64) []OverviewTeamStruct {
	var result []OverviewTeamStruct
	sqlStr := `
	SELECT tg.id,
        tg.name,
		tg.created_at,user.username,
		ut.role FROM user_team as ut
LEFT JOIN team_group as tg ON tg.id=ut.team_id
LEFT JOIN user ON user.id=tg.user_id
WHERE
	tg.deleted_at IS NULL
AND ut.deleted_at IS NULL
AND ut.user_id = @userId
	`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userId}).Scan(&result)
	return result
}

type OverviewTeamSearchStruct struct {
	Id        string    `json:"id" grom:"column:id"`
	Name      string    `json:"name" gorm:"column:"name"`
	Username  string    `json:"username" gorm:"column:username"`
	Role      int       `json:"role" gorm:"column:role"`
	CreatedAt LocalTime `json:"created_at" gorm:"column:created_at" type:"datetime"`
	Search    int       `json:"search" gorm:"column:-"`
}

func GetOverviewTeamSearch(userId int64) []OverviewTeamSearchStruct {
	var result []OverviewTeamSearchStruct
	sqlStr := `
	SELECT tg.id,
        tg.name,
		tg.created_at,user.username,
		ut.role FROM user_team as ut
LEFT JOIN team_group as tg ON tg.id=ut.team_id
LEFT JOIN user ON user.id=tg.user_id
WHERE
	tg.deleted_at IS NULL
AND ut.deleted_at IS NULL
AND ut.user_id = @userId
	`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userId}).Scan(&result)
	return result
}

func SearchOverviewTeamNumber(teamId int64, searchStr string) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from goods_list where team_id=@teamId and deleted_at is null and is_thrown=0 and name like '%` + searchStr + `%'`

	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

func GetAllGoodsNumberByTeamId(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from goods_list where team_id=@teamId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

func GetAllActiveGoodsNumberByTeamId(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from goods_list where team_id=@teamId and deleted_at is null and is_expired=0 and is_thrown=0`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

func GetAllExpiredGoodsNumberByTeamId(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from goods_list where team_id=@teamId and deleted_at is null and is_expired=1 and is_thrown=0`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

func GetAllThrownGoodsNumberByTeamId(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from goods_list where team_id=@teamId and deleted_at is null and is_thrown=1`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

type RCalendarStruct struct {
	TeamId     string       `json:"team_id" gorm:"column:team_id"`
	GroupId    string       `json:"group_id" gorm:"column:group_id"`
	TeamName   string       `json:"team_name" gorm:"column:team_name"`
	GroupName  string       `json:"group_name" gorm:"column:group_name"`
	GoodsName  string       `json:"goods_name" gorm:"column:goods_name"`
	ExpireTime LocalTimeIOS `json:"expire_time" gorm:"expire_time"`
}

func GetOverviewCalendar(userId int64, year int, month int) []RCalendarStruct {
	var result []RCalendarStruct
	sqlStr := `
	SELECT
	gl.team_id,
	gl.group_id,
	gl.NAME as goods_name,
	tg.NAME AS team_name,
	gg.NAME AS group_name,
	gl.expire_time
FROM
	goods_list AS gl
	LEFT JOIN team_group AS tg ON tg.id = gl.team_id
	LEFT JOIN goods_group AS gg ON gg.id = gl.group_id 
WHERE
	YEAR ( gl.expire_time )= @year 
	AND MONTH ( gl.expire_time )= @month
	AND gl.team_id IN ( SELECT team_id FROM user_team WHERE user_id = @userId AND deleted_at IS NULL ) 
	AND gl.deleted_at IS NULL 
	and tg.deleted_at is null
	and gg.deleted_at is null
	AND gl.is_thrown =0
	order by gl.expire_time
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
		"year":   year,
		"month":  month,
	}).Scan(&result)
	return result
}
