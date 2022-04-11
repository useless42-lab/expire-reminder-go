package controller

import (
	"reminder/models"
	"reminder/response"
	"reminder/service"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GoodsGroupForm struct {
	Name string `json:"name" gorm:"column:name"`
}

func (form GoodsGroupForm) ValidateGoodsGroupForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("名称不能为空")),
	)
}
func AddGoodsGroup(c *gin.Context) {
	// userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	name := c.PostForm("name")
	goodsGroupForm := GoodsGroupForm{
		Name: name,
	}
	err1 := goodsGroupForm.ValidateGoodsGroupForm()
	if err1 != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err1))
		return
	}
	// models.AddGoodsGroup(teamId, name)
	err := service.AddGoodsGroupService(teamId, name)
	if err != "" {
		response.Error(c, 3000, err)
	} else {
		response.Success(c, 200, "")
	}
	// response.Success(c, 200, "")
}

func DeleteGoodsGroup(c *gin.Context) {
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	// models.DeleteGoodsGroup(groupId)
	result := service.DeleteDeviceGroupService(teamId, groupId)
	if result != "" {
		response.Error(c, 4000, result)
	} else {
		response.Success(c, 200, "")
	}
}

func UpdateGoodsGroup(c *gin.Context) {
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	name := c.PostForm("name")
	models.UpdateGoodsGroup(groupId, name)
	response.Success(c, 200, "")
}
