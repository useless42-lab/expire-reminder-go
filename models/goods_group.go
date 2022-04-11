package models

import (
	"reminder/tools"
	"time"
)

type GoodsGroupStruct struct {
	DefaultModel
	Name   string `json:"name" gorm:"column:name"`
	TeamId int64  `json:"team_id" gorm:"column:team_id"`
}

func AddGoodsGroup(teamId int64, name string) GoodsGroupStruct {
	data := GoodsGroupStruct{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		TeamId:       teamId,
		Name:         name,
	}
	DB.Table("goods_group").Create(&data)
	return data
}

type RGoodsGroup struct {
	Id   string `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func GetGoodsGroupList(teamId int64) []RGoodsGroup {
	var result []RGoodsGroup
	sqlStr := `select id,name from goods_group where team_id=@teamId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

// 删除团队分组
func DeleteGoodsGroup(groupId int64) {
	sqlStr := `update goods_group set deleted_at=@deletedAt where id=@groupId`
	DB.Exec(sqlStr, map[string]interface{}{"groupId": groupId, "deletedAt": time.Now()})
}

func UpdateGoodsGroup(groupId int64, name string) {
	sqlStr := `update goods_group set name=@name where id =@groupId`
	DB.Exec(sqlStr, map[string]interface{}{"groupId": groupId, "name": name})
}

func GetTeamGoodsGroupNumber(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from goods_group where team_id=@teamId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}
