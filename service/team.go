package service

import (
	"crypto/md5"
	"encoding/hex"
	"reminder/cache"
	"reminder/models"
	"strconv"
	"strings"
	"time"
)

func GetTeamGroupListService(userId int64) []models.RUserTeamGroup {
	teamList := models.GetTeamGroupList(userId)
	for i := 0; i < len(teamList); i++ {
		teamId, _ := strconv.ParseInt(teamList[i].Id, 10, 64)
		groupList := models.GetGoodsGroupList(teamId)
		teamList[i].GroupList = groupList
	}
	return teamList
}

func GetTeamGroupMemberListService(teamId int64, page int, pageSize int) models.PaginationData {
	data := models.GetTeamGroupMemberList(teamId, page, pageSize)
	total := models.GetTeamGroupMemberNumber(teamId)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

type RTeamGroupInfo struct {
	Id               string           `json:"id" gorm:"column:id"`
	Name             string           `json:"name" gorm:"column:name"`
	UserId           string           `json:"user_id" gorm:"column:user_id"`
	CreatedAt        models.LocalTime `json:"created_at" gorm:"column:created_at"`
	MemberTotal      int              `json:"member_total"`
	DeviceGroupTotal int              `json:"device_group_total"`
	DeviceTotal      int              `json:"device_total"`
	Role             int              `json:"role" gorm:"column:role"`
}

func GetTeamGroupInfoService(teamId int64, userId int64) RTeamGroupInfo {
	userTeamInfo := models.GetUserTeamInfo(teamId, userId)
	info := models.GetTeamGroupInfo(teamId)
	memberTotal := models.GetTeamGroupMemberNumber(teamId)
	// deviceGroupTotal := models.GetTeamGroupDeviceGroupNumber(teamId)
	// deviceGroupTotal := models.GetDeviceGroupCount(userId, teamId)
	// deviceTotal := models.GetAllDeviceNumber(teamId)
	result := RTeamGroupInfo{
		Id:               info.Id,
		Name:             info.Name,
		UserId:           info.UserId,
		CreatedAt:        info.CreatedAt,
		MemberTotal:      memberTotal.Total,
		DeviceGroupTotal: 0,
		DeviceTotal:      0,
		Role:             userTeamInfo.Role,
	}
	return result
}

// ????????????
func DeleteTeamGroupService(userId int64, teamId int64) string {
	// GetUserTeamGroupNumber
	count := models.GetOwnerTotalTeamNumber(userId)
	if count.Total < 2 {
		// ??????????????????????????????
		return "??????????????????????????????"
	} else {
		models.DeleteTeamGroup(userId, teamId)
		return ""
	}
}

// ????????????
func ExitTeamService(userId int64, teamId int64) int {
	role := models.CheckTeamMemberRole(userId, teamId)
	if role.Role == 2 {
		// ???????????????????????????
		return 0
	} else {
		models.DeleteTeamMember(teamId, userId)
		return 1
	}
}

// ????????????
func KickOutTeamMemberService(userId int64, teamId int64, targetUserId int64) int {
	role := models.CheckTeamMemberRole(userId, teamId)
	if role.Role == 1 {
		// ????????????????????????
		return 0
	} else {
		models.DeleteTeamMember(teamId, targetUserId)
		return 1
	}
}

// ???????????????
func RemoveAdminService(userId int64, teamId int64, targetUserId int64) {
	role := models.CheckTeamMemberRole(userId, teamId)
	if role.Role == 2 {
		// ??????????????????
		models.UpdateUserTeamRole(teamId, targetUserId, 1)
	} else {
	}
}

// ???????????????
func AddAdminService(userId int64, teamId int64, targetUserId int64) {
	role := models.CheckTeamMemberRole(userId, teamId)
	if role.Role == 2 {
		// ??????????????????
		models.UpdateUserTeamRole(teamId, targetUserId, 3)
	} else {

	}
}

// ????????????????????????
func GenerateTransferTeamLinkService(teamId int64, userId int64) string {
	token := md5.New()
	token.Write([]byte(strconv.FormatInt(teamId, 10) + strconv.FormatInt(userId, 10) + time.Now().String()))
	finalToken := hex.EncodeToString(token.Sum(nil))
	// cache.Set("user:token:"+finalToken, strconv.Itoa(int(result.DefaultModel.ID)), 60*60*24*7)
	data := strconv.FormatInt(teamId, 10) + " " + strconv.FormatInt(userId, 10)
	cache.Set("team:group:"+finalToken, data, 30*60)
	return finalToken
}

// ????????????????????????
func GenerateInviteTeamMemberLinkService(teamId int64, userId int64, role int) string {
	token := md5.New()
	token.Write([]byte(strconv.FormatInt(teamId, 10) + strconv.FormatInt(userId, 10) + strconv.Itoa(role) + time.Now().String()))
	finalToken := hex.EncodeToString(token.Sum(nil))
	// cache.Set("user:token:"+finalToken, strconv.Itoa(int(result.DefaultModel.ID)), 60*60*24*7)
	data := strconv.FormatInt(teamId, 10) + " " + strconv.FormatInt(userId, 10) + " " + strconv.Itoa(role)
	cache.Set("team:member:"+finalToken, data, 30*60)
	return finalToken
}

func GetTransferTeamInfoService(token string) models.RSimpleTeam {
	var result models.RSimpleTeam
	data := cache.Get("team:group:" + token)
	if data != "" {
		arr := strings.Fields(data)
		teamId, _ := strconv.ParseInt(arr[0], 10, 64)
		ownerId, _ := strconv.ParseInt(arr[1], 10, 64)
		result = models.GetSimpleTeamInfo(teamId, ownerId)
		return result
	} else {
		return result
	}
}

func TransferTeamGroupService(token string, userId int64) string {
	data := cache.Get("team:group:" + token)
	if data != "" {
		arr := strings.Fields(data)
		teamId, _ := strconv.ParseInt(arr[0], 10, 64)
		ownerId, _ := strconv.ParseInt(arr[1], 10, 64)

		planId := models.GetPlanIdByUserId(userId)
		plan := models.GetPlanBaseInfo(planId.Id)

		teamNumber := models.GetOwnerTotalTeamNumber(userId)
		if plan.TeamNumber <= teamNumber.Total {
			return "???????????????????????????????????????"
		}
		// ?????? team_group ?????????
		models.UpdateTeamOwner(teamId, ownerId, userId)
		// ???????????????????????????????????????
		num := models.CheckIsUserInTeam(userId, teamId)
		if num.Total > 0 {
			// ?????????????????????
			// ???????????????????????? ????????? 2
			models.UpdateUserTeamRole(teamId, userId, 2)
		} else {
			models.AddUserTeam(userId, teamId, 2)
		}
		// ????????????
		number := models.GetUserTeamGroupNumber(ownerId)
		models.DeleteTeamMember(teamId, ownerId)
		if number.Total > 0 {
		} else {
			// ???????????????????????? ?????????
			InitNewTeamGroupService(ownerId)
		}
		cache.Del("team:group:" + token)
	} else {
		return "????????????"
	}
	return ""
}

type RSimpleTeamWithRole struct {
	models.RSimpleTeam
	Role int `json:"role"`
}

func GetInviteTeamMemberInfoService(token string) RSimpleTeamWithRole {
	var result RSimpleTeamWithRole
	data := cache.Get("team:member:" + token)
	if data != "" {
		arr := strings.Fields(data)
		teamId, _ := strconv.ParseInt(arr[0], 10, 64)
		ownerId, _ := strconv.ParseInt(arr[1], 10, 64)
		role, _ := strconv.Atoi(arr[2])
		result.RSimpleTeam = models.GetSimpleTeamInfo(teamId, ownerId)
		result.Role = role
		return result
	} else {
		return result
	}
}

func CreateInviteTeamMemberService(token string, userId int64) int {
	data := cache.Get("team:member:" + token)
	if data != "" {
		arr := strings.Fields(data)
		teamId, _ := strconv.ParseInt(arr[0], 10, 64)
		role, _ := strconv.Atoi(arr[2])
		teamMemberTotal := models.GetTeamMemberCount(teamId)
		planBaseItem := models.GetPlanBaseInfoByTeamId(teamId)
		if teamMemberTotal.Total < planBaseItem.TeamMemberLimit {
			num := models.CheckIsUserInTeam(userId, teamId)
			if num.Total > 0 {
				// ?????????????????????
				return 0
			} else {
				models.AddUserTeam(userId, teamId, role)
				return 1
			}
		}
	}
	return 0
}

func UpdateTeamGroupService(teamId int64, name string) {
	// role := models.CheckTeamMemberRole(userId, teamId)
	// if role.Role == 2 || role.Role == 3 {
	// 	// ??????????????????
	// 	models.UpdateTeamGroup(userId, teamId, name)
	// } else {

	// }
	models.UpdateTeamGroup(teamId, name)
}
