package models

import (
	"reminder/tools"
	"time"
)

type GoodsItem struct {
	DefaultModel
	Barcode     string    `json:"barcode" gorm:"column:barcode"`
	UserId      int64     `json:"user_Id" gorm:"column:user_id"`
	TeamId      int64     `json:"team_id" gorm:"column:team_id"`
	GroupId     int64     `json:"group_id" gorm:"column:group_id"`
	Name        string    `json:"name" gorm:"column:name"`
	Price       float64   `json:"price" gorm:"column:price"`
	Number      int       `json:"number" gorm:"column:number"`
	Remarks     string    `json:"remarks" gorm:"column:remarks"`
	Category    int       `json:"category" gorm:"column:category"`
	ProductTime time.Time `json:"product_time" gorm:"column:product_time"`
	ShelfLife   int       `json:"shelf_life" gorm:"column:shelf_life"`
	ExpireTime  time.Time `json:"expire_time" gorm:"column:expire_time"`
	AdvanceDay  int       `json:"advance_day" gorm:"column:advance_day"`
}

func AddGoods(barcode string, userId int64, teamId int64, groupId int64, name string, price float64, number int, remarks string, category int, productTime time.Time, shelfLife int, expireTime time.Time, advanceDay int) GoodsItem {
	data := GoodsItem{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		Barcode:      barcode,
		UserId:       userId,
		TeamId:       teamId,
		GroupId:      groupId,
		Name:         name,
		Price:        price,
		Number:       number,
		Remarks:      remarks,
		Category:     category,
		ProductTime:  productTime,
		ShelfLife:    shelfLife,
		ExpireTime:   expireTime,
		AdvanceDay:   advanceDay,
	}
	DB.Table("goods_list").Create(&data)
	return data
}

type RGoodsItem struct {
	Id          string    `json:"id" gorm:"column:id"`
	UserId      string    `json:"user_Id" gorm:"column:user_id"`
	TeamId      string    `json:"team_id" gorm:"column:team_id"`
	GroupId     string    `json:"group_id" gorm:"column:group_id"`
	Name        string    `json:"name" gorm:"column:name"`
	Number      int       `json:"number" gorm:"column:number"`
	Remarks     string    `json:"remarks" gorm:"column:remarks"`
	Category    int       `json:"category" gorm:"column:category"`
	ProductTime LocalTime `json:"product_time" gorm:"column:product_time"`
	ShelfLife   int       `json:"shelf_life" gorm:"column:shelf_life"`
	ExpireTime  LocalTime `json:"expire_time" gorm:"column:expire_time"`
	AdvanceDay  int       `json:"advance_day" gorm:"column:advance_day"`
}

func GetGoodsList(searchStr string, teamId int64, groupId int64, category int, tagId int64) []RGoodsItem {
	var result []RGoodsItem
	var sqlStr string
	if category == 0 {
		if tagId == 0 {
			sqlStr = `select * from goods_list where team_id=@teamId and group_id=@groupId and deleted_at is null and is_expired=0 and is_thrown=0 and name like '%` + searchStr + `%' order by expire_time`
		} else {
			sqlStr = `
			SELECT
	gl.id,gl.user_id,gl.team_id,gl.group_id,gl.name,gl.number,gl.remarks,gl.category,gl.product_time,gl.shelf_life,gl.expire_time,gl.advance_day 
FROM
	goods_list AS gl
	LEFT JOIN tag_list AS tl ON gl.id = tl.goods_id 
WHERE
	gl.team_id = @teamId 
	AND gl.group_id = @groupId 
	AND gl.deleted_at IS NULL 
	AND gl.is_expired = 0 
	AND tl.tag_id=@tagId
	and is_thrown=0
	and name like '%` + searchStr + `%'
ORDER BY
	expire_time
			`
		}

	} else {
		if tagId == 0 {
			sqlStr = `select * from goods_list where team_id=@teamId and group_id=@groupId and category=@category and deleted_at is null and is_expired=0 and is_thrown=0 and name like '%` + searchStr + `%' order by expire_time`
		} else {
			sqlStr = `
			SELECT
	gl.id,gl.user_id,gl.team_id,gl.group_id,gl.name,gl.number,gl.remarks,gl.category,gl.product_time,gl.shelf_life,gl.expire_time,gl.advance_day 
FROM
	goods_list AS gl
	LEFT JOIN tag_list AS tl ON gl.id = tl.goods_id 
WHERE
	gl.team_id = @teamId 
	AND gl.group_id = @groupId 
	and category=@category
	AND gl.deleted_at IS NULL 
	AND gl.is_expired = 0 
	AND tl.tag_id=@tagId
	and is_thrown=0
	and name like '%` + searchStr + `%'
ORDER BY
	expire_time
			`
		}
	}
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId":   teamId,
		"groupId":  groupId,
		"category": category,
		"tagId":    tagId,
	}).Scan(&result)
	return result
}

func GetExpiredGoodsList(searchStr string, teamId int64, groupId int64, category int, tagId int64) []RGoodsItem {
	var result []RGoodsItem
	var sqlStr string
	// if category == 0 {
	// 	sqlStr = `select * from goods_list where team_id=@teamId and group_id=@groupId and deleted_at is null and is_expired=1 order by expire_time desc`
	// } else {
	// 	sqlStr = `select * from goods_list where team_id=@teamId and group_id=@groupId and category=@category and deleted_at is null and is_expired=1 order by expire_time desc`
	// }
	if category == 0 {
		if tagId == 0 {
			sqlStr = `select * from goods_list where team_id=@teamId and group_id=@groupId and deleted_at is null and is_expired=1 and is_thrown=0 and name like '%` + searchStr + `%' order by expire_time desc`
		} else {
			sqlStr = `
			SELECT
	gl.id,gl.user_id,gl.team_id,gl.group_id,gl.name,gl.number,gl.remarks,gl.category,gl.product_time,gl.shelf_life,gl.expire_time,gl.advance_day 
FROM
	goods_list AS gl
	LEFT JOIN tag_list AS tl ON gl.id = tl.goods_id 
WHERE
	gl.team_id = @teamId 
	AND gl.group_id = @groupId 
	AND gl.deleted_at IS NULL 
	AND gl.is_expired = 1 
	AND tl.tag_id=@tagId
	and is_thrown=0
	and name like '%` + searchStr + `%'
ORDER BY
	expire_time desc
			`
		}

	} else {
		if tagId == 0 {
			sqlStr = `select * from goods_list where team_id=@teamId and group_id=@groupId and category=@category and deleted_at is null and is_expired=1 and is_thrown=0 and name like '%` + searchStr + `%' order by expire_time desc`
		} else {
			sqlStr = `
			SELECT
	gl.id,gl.user_id,gl.team_id,gl.group_id,gl.name,gl.number,gl.remarks,gl.category,gl.product_time,gl.shelf_life,gl.expire_time,gl.advance_day 
FROM
	goods_list AS gl
	LEFT JOIN tag_list AS tl ON gl.id = tl.goods_id 
WHERE
	gl.team_id = @teamId 
	AND gl.group_id = @groupId 
	and category=@category
	AND gl.deleted_at IS NULL 
	AND gl.is_expired = 1 
	AND tl.tag_id=@tagId
	and is_thrown=0
	and name like '%` + searchStr + `%'
ORDER BY
	expire_time desc
			`
		}
	}
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId":   teamId,
		"groupId":  groupId,
		"category": category,
	}).Scan(&result)
	return result
}

type RGoodListItem struct {
	Id         int64     `json:"id" gorm:"column:id"`
	TeamId     int64     `json:"team_id" gorm:"column:team_id"`
	GroupId    string    `json:"group_id" gorm:"column:group_id"`
	Name       string    `json:"name" gorm:"column:name"`
	ExpireTime time.Time `json:"expire_time" gorm:"column:expire_time"`
	AdvanceDay int       `json:"advance_day" gorm:"column:advance_day"`
}

func GetAllGoodsList() []RGoodListItem {
	var result []RGoodListItem
	sqlStr := `select id,team_id,group_id,name,expire_time,advance_day from goods_list where deleted_at is null and is_expired=0`
	DB.Raw(sqlStr).Scan(&result)
	return result
}

func GetTeamGoodsNumber(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from goods_list where team_id=@teamId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

func SetGoodsExpire(id int64) {
	sqlStr := `update goods_list set is_expired=1 where id=@id`
	DB.Exec(sqlStr, map[string]interface{}{
		"id": id,
	})
}

type GoodsDetail struct {
	Id          string    `json:"id" gorm:"column:id"`
	Barcode     string    `json:"barcode" gorm:"column:barcode"`
	Name        string    `json:"name" gorm:"column:name"`
	ProductTime LocalTime `json:"product_time" gorm:"column:product_time"`
	ExpireTime  LocalTime `json:"expire_time" gorm:"column:expire_time"`
	Img         string    `json:"img" gorm:"column:img"`
	Remarks     string    `json:"remarks" gorm:"column:remarks"`
}

func GetGoodsDetail(id int64) GoodsDetail {
	sqlStr := `select goods_list.id,goods_list.barcode,goods_list.name,goods_list.product_time,goods_list.expire_time,goods_list.remarks,goods_base.img from goods_list 
	left join goods_base on goods_base.barcode=goods_list.barcode where goods_list.id=@id`
	var result GoodsDetail
	DB.Raw(sqlStr, map[string]interface{}{
		"id": id,
	}).Scan(&result)
	return result
}

func GetGroupGoodsCount(groupId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from goods_list where group_id=@groupId and deleted_at is null and is_expired=0`
	DB.Raw(sqlStr, map[string]interface{}{
		"groupId": groupId,
	}).Scan(&result)
	return result
}

func DeleteGoods(goodsId int64, teamId int64, groupId int64) {
	sqlStr := `update goods_list set deleted_at=@deletedAt where id=@goodsId and team_id=@teamId and group_id=@groupId`
	DB.Exec(sqlStr, map[string]interface{}{
		"deletedAt": time.Now(),
		"goodsId":   goodsId,
		"teamId":    teamId,
		"groupId":   groupId,
	})
}

type GoodsPrice struct {
	Price     float64   `json:"price" gorm:"column:price"`
	CreatedAt LocalTime `json:"created_at" gorm:"created_at"`
}

func GetGoodsPriceTimeline(teamId int64, barcode string) []GoodsPrice {
	var result []GoodsPrice
	sqlStr := `select price,created_at from goods_list where barcode=@barcode and team_id=@teamId order by id desc`
	DB.Raw(sqlStr, map[string]interface{}{
		"barcode": barcode,
		"teamId":  teamId,
	}).Scan(&result)
	return result
}

func GetGoodsPriceTimelineByGoodsId(teamId int64, goodsId int64) []GoodsPrice {
	var result []GoodsPrice
	sqlStr := `select price,created_at from goods_list where id=@goodsId and team_id=@teamId order by id desc`
	DB.Raw(sqlStr, map[string]interface{}{
		"goodsId": goodsId,
		"teamId":  teamId,
	}).Scan(&result)
	return result
}

func ThrownGoods(goodsId int64) {
	sqlStr := `update goods_list set is_thrown=1 where id=@goodsId`
	DB.Exec(sqlStr, map[string]interface{}{
		"goodsId": goodsId,
	})
}
