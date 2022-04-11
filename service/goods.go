package service

import (
	"reminder/models"
	"strings"
	"time"
)

func AddGoodsService(barcode string, userId int64, teamId int64, groupId int64, name string, price float64, number int, tags string, remarks string, category int, productTime time.Time, shelfLife int, expireTime time.Time, advanceDay int) string {
	planId := models.GetPlanIdByTeamId(teamId)
	plan := models.GetPlanBaseInfo(planId.Id)
	goodsNumber := models.GetTeamGoodsNumber(teamId)
	if goodsNumber.Total < plan.PerTeamGoodsLimit {
		goodsData := models.AddGoods(barcode, userId, teamId, groupId, name, price, number, remarks, category, productTime, shelfLife, expireTime, advanceDay)
		tagsArr := strings.Split(tags, ",")
		if len(tagsArr) > 0 {
			for _, item := range tagsArr {
				if item != "" {
					tagExistId := models.IsExistTag(teamId, groupId, item)
					if tagExistId == 0 {
						tagBaseData := models.AddTagBase(teamId, groupId, item)
						models.AddTagList(tagBaseData.DefaultModel.ID, goodsData.DefaultModel.ID)
					} else {
						models.AddTagList(tagExistId, goodsData.DefaultModel.ID)
					}
				}
			}
		}
		return ""
	} else {
		return "该空间物品数量达到上限"
	}
}
