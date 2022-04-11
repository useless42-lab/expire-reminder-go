package controller

import (
	"reminder/models"
	"reminder/response"
	"reminder/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GoodsForm struct {
	Name        string    `json:"name"`
	Number      int       `json:"number"`
	Category    int       `json:"category"`
	ProductTime time.Time `json:"product_time"`
	ShelfLife   int       `json:"shelf_life"`
	ExpireTime  time.Time `json:"expire_time"`
	AdvanceDay  int       `json:"advance_day"`
}

func (form GoodsForm) ValidateGoodsForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("名称不能为空")),
		validation.Field(&form.Number, validation.Required.Error("数量不能为空")),
		validation.Field(&form.Category, validation.Required.Error("分类不能为空")),
		validation.Field(&form.ProductTime, validation.Required.Error("生产日期不能为空")),
		validation.Field(&form.ShelfLife, validation.Required.Error("保质期限不能为空")),
		validation.Field(&form.ExpireTime, validation.Required.Error("过期时间不能为空")),
		validation.Field(&form.AdvanceDay, validation.Required.Error("提前时间不能为空")),
	)
}

func AddGoods(c *gin.Context) {
	var timeLayoutStr = "2006-01-02 15:04:05"
	barcode := c.PostForm("barcode")
	userId := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	number, _ := strconv.Atoi(c.PostForm("number"))
	name := c.PostForm("name")
	tags := c.PostForm("tags")
	remarks := c.PostForm("remarks")
	category, _ := strconv.Atoi(c.PostForm("category"))
	productTime, _ := time.ParseInLocation(timeLayoutStr, c.PostForm("product_time"), time.Local)
	shelfLife, _ := strconv.Atoi(c.PostForm("shelf_life"))
	expireTime, _ := time.ParseInLocation(timeLayoutStr, c.PostForm("expire_time"), time.Local)
	advanceDay, _ := strconv.Atoi(c.PostForm("advance_day"))
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	goodsForm := GoodsForm{
		Name:        name,
		Number:      number,
		Category:    category,
		ProductTime: productTime,
		ShelfLife:   shelfLife,
		ExpireTime:  expireTime,
		AdvanceDay:  advanceDay,
	}
	err := goodsForm.ValidateGoodsForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	result := service.AddGoodsService(barcode, userId, teamId, groupId, name, price, number, tags, remarks, category, productTime, shelfLife, expireTime, advanceDay)
	if result == "" {
		response.Success(c, 200, "")
	} else {
		response.Error(c, 5101, result)
	}
}

func GetGoodsList(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	tagId, _ := strconv.ParseInt(c.Query("tag_id"), 10, 64)
	category, _ := strconv.Atoi(c.Query("category"))
	searchStr := c.Query("search_str")
	result := models.GetGoodsList(searchStr, teamId, groupId, category, tagId)
	response.Success(c, 200, result)
}

func GetExpiredGoodsList(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	tagId, _ := strconv.ParseInt(c.Query("tag_id"), 10, 64)
	category, _ := strconv.Atoi(c.Query("category"))
	searchStr := c.Query("search_str")
	result := models.GetExpiredGoodsList(searchStr, teamId, groupId, category, tagId)
	response.Success(c, 200, result)
}

func GetGoodsDetail(c *gin.Context) {
	goodsId, _ := strconv.ParseInt(c.Query("goods_id"), 10, 64)
	result := models.GetGoodsDetail(goodsId)
	response.Success(c, 200, result)
}

func DeleteGoods(c *gin.Context) {
	goodsId, _ := strconv.ParseInt(c.PostForm("goods_id"), 10, 64)
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	models.DeleteGoods(goodsId, teamId, groupId)
	response.Success(c, 200, "")
}

func GetGoodsPriceTimeline(c *gin.Context) {
	barcode := c.Query("barcode")
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	goodsId, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if barcode == "" {
		// response.Error(c, 5006, "仅限包含条形码的物品可查询价格")
		result := models.GetGoodsPriceTimelineByGoodsId(teamId, goodsId)
		response.Success(c, 200, result)
	} else {
		result := models.GetGoodsPriceTimeline(teamId, barcode)
		response.Success(c, 200, result)
	}
}

func ThrownGoods(c *gin.Context) {
	goodsId, _ := strconv.ParseInt(c.PostForm("goods_id"), 10, 64)
	models.ThrownGoods(goodsId)
	response.Success(c, 200, "")
}
