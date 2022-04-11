package models

import (
	"reminder/tools"
)

type TagBaseStruct struct {
	DefaultModel
	TeamId  int64  `json:"team_id" gorm:"column:team_id"`
	GroupId int64  `json:"group_id" gorm:"column:group_id"`
	Name    string `json:"name gorm:"column:name`
}

func AddTagBase(teamId int64, groupId int64, name string) TagBaseStruct {
	data := TagBaseStruct{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		TeamId:       teamId,
		GroupId:      groupId,
		Name:         name,
	}
	DB.Table("tag_base").Create(&data)
	return data
}

func IsExistTag(teamId int64, groupId int64, name string) int64 {
	var result TagBaseStruct
	sqlStr := `select id from tag_base where team_id=@teamId and group_id=@groupId and name=@name`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId":  teamId,
		"groupId": groupId,
		"name":    name,
	}).Scan(&result)
	return result.ID
}

type RTagBase struct {
	Id   string `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func GetTagList(teamId int64, groupId int64) []RTagBase {
	var result []RTagBase
	sqlStr := `select id,name from tag_base where team_id=@teamId and group_id=@groupId`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId":  teamId,
		"groupId": groupId,
	}).Scan(&result)
	return result
}
