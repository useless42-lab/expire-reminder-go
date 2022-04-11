package service

import (
	"reminder/models"
	"strconv"
)

func GetOverviewTeamService(userId int64) []models.OverviewTeamStruct {
	teamList := models.GetOverviewTeam(userId)
	for i := 0; i < len(teamList); i++ {
		teamId, _ := strconv.ParseInt(teamList[i].Id, 10, 64)
		teamList[i].GoodsOverviewNumber.All = models.GetAllGoodsNumberByTeamId(teamId).Total
		teamList[i].GoodsOverviewNumber.Active = models.GetAllActiveGoodsNumberByTeamId(teamId).Total
		teamList[i].GoodsOverviewNumber.Expired = models.GetAllExpiredGoodsNumberByTeamId(teamId).Total
		teamList[i].GoodsOverviewNumber.Thrown = models.GetAllThrownGoodsNumberByTeamId(teamId).Total
	}
	return teamList
}

func SearchOverviewTeamService(userId int64, searchStr string) []models.OverviewTeamSearchStruct {
	teamList := models.GetOverviewTeamSearch(userId)
	for i := 0; i < len(teamList); i++ {
		teamId, _ := strconv.ParseInt(teamList[i].Id, 10, 64)
		teamList[i].Search = models.SearchOverviewTeamNumber(teamId, searchStr).Total
	}
	return teamList
}
