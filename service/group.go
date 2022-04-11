package service

import "reminder/models"

func InitNewTeamGroupService(userId int64) {
	team := models.AddTeamGroup(userId, "默认空间")
	models.AddUserTeam(userId, int64(team.ID), 2)
	models.AddGoodsGroup(int64(team.ID), "默认分组")
}

func CreateTeamGroupService(userId int64, teamName string) string {
	planId := models.GetPlanIdByUserId(userId)
	plan := models.GetPlanBaseInfo(planId.Id)
	teamNumber := models.GetUserTeamGroupNumber(userId)
	if teamNumber.Total < plan.TeamNumber {
		team := models.AddTeamGroup(userId, teamName)
		models.AddUserTeam(userId, int64(team.ID), 2)
		models.AddGoodsGroup(int64(team.ID), "默认分组")
		return ""
	} else {
		return "团队空间数量超出上限"
	}
}

func AddGoodsGroupService(teamId int64, name string) string {
	planId := models.GetPlanIdByTeamId(teamId)
	plan := models.GetPlanBaseInfo(planId.Id)
	goodsGroupNumber := models.GetTeamGoodsGroupNumber(teamId)
	if goodsGroupNumber.Total < plan.PerTeamGroupLimit {
		models.AddGoodsGroup(teamId, name)
		return ""
	} else {
		return "空间分组数量超出上限"
	}
}

func DeleteDeviceGroupService(teamId int64, groupId int64) string {
	total := models.GetGroupGoodsCount(groupId)
	if total.Total > 0 {
		return "该分组下还有物品，无法删除"
	} else {
		total = models.GetTeamGoodsGroupNumber(teamId)
		if total.Total == 1 {
			return "无法删除团队分类下最后一个分组"
		} else {
			models.DeleteGoodsGroup(groupId)
			return ""
		}
	}
}
