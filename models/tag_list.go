package models

type TagListStruct struct {
	DefaultModel
	TagId   int64 `json:"tag_id" gorm:"column:tag_id"`
	GoodsId int64 `json:"goods_id" gorm:"column:goods_id"`
}

func AddTagList(tagId int64, goodsId int64) {
	data := TagListStruct{
		TagId:   tagId,
		GoodsId: goodsId,
	}
	DB.Table("tag_list").Create(&data)
}
