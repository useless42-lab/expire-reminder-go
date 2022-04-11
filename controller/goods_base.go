package controller

import (
	"reminder/models"
	"reminder/response"

	"github.com/gin-gonic/gin"
)

func GetGoodsName(c *gin.Context) {
	barcode := c.Query("barcode")
	existGoodsInfo := models.GetBarcodeGoods(barcode)
	if existGoodsInfo == "" {
		goodsInfo := OnGetGoodsApi(barcode)
		if goodsInfo.Data.GoodsName != "" && goodsInfo.Code == 200 {
			models.AddGoodsBase(goodsInfo.Data.Code, goodsInfo.Data.SptmImg, goodsInfo.Data.Img, goodsInfo.Data.GoodsType, goodsInfo.Data.Trademark, goodsInfo.Data.GoodsName,
				goodsInfo.Data.Spec, goodsInfo.Data.Note, goodsInfo.Data.Price, goodsInfo.Data.Ycg, goodsInfo.Data.ManuName, goodsInfo.Data.ManuAddress, goodsInfo.Data.Qs, goodsInfo.Data.Nw, goodsInfo.Data.Description, goodsInfo.Data.Gw, goodsInfo.Data.Width, goodsInfo.Data.Hight, goodsInfo.Data.Depth, goodsInfo.Data.Gpc, goodsInfo.Data.GpcType, goodsInfo.Data.Keyword)
		}
		response.Success(c, 200, goodsInfo.Data.GoodsName)
	} else {
		response.Success(c, 200, existGoodsInfo)
	}
}
