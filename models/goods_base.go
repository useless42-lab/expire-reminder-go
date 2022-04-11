package models

type GoodsBaseItem struct {
	Barcode     string `json:"barcode" gorm:"column:barcode"`
	SptmImg     string `json:"sptmImg" gorm:"column:sptm_img"`
	Img         string `json:"img" gorm:"column:img"`
	GoodsType   string `json:"goodsType" gorm:"column:goods_type"`
	Trademark   string `json:"trademark" gorm:"column:trademark"`
	GoodsName   string `json:"goodsName" gorm:"column:goods_name"`
	Spec        string `json:"spec" gorm:"column:spec"`
	Note        string `json:"note" gorm:"column:note"`
	Price       string `json:"price" gorm:"column:price"`
	Ycg         string `json:"ycg" gorm:"column:ycg"`
	ManuName    string `json:"manuName" gorm:"column:manu_name"`
	ManuAddress string `json:"manuAddress" gorm:"column:manu_address"`
	Qs          string `json:"qs"`
	Nw          string `json:"nw"`
	Description string `json:"description"`
	Gw          string `json:"gw"`
	Width       string `json:"width"`
	Hight       string `json:"hight"`
	Depth       string `json:"depth"`
	Gpc         string `json:"gpc"`
	GpcType     string `json:"gpcType" gorm:"column:gpc_type"`
	Keyword     string `json:"keyword"`
	// ImgList     []interface{} `json:"imgList"`
}

func AddGoodsBase(
	barcode string, sptmImg string, img string, goodsType string, trademark string, goodsName string, spec string, note string,
	price string, ycg string, manuName string, manuAddress, qs string, nw string, description string, gw string, width string, hight string,
	depth string, gpc string, gpcType string, keyword string,
) {
	data := GoodsBaseItem{
		Barcode:     barcode,
		SptmImg:     sptmImg,
		Img:         img,
		GoodsType:   goodsType,
		Trademark:   trademark,
		GoodsName:   goodsName,
		Spec:        spec,
		Note:        note,
		Price:       price,
		Ycg:         ycg,
		ManuName:    manuName,
		ManuAddress: manuAddress,
		Qs:          qs,
		Nw:          nw,
		Description: description,
		Gw:          gw,
		Width:       width,
		Hight:       hight,
		Depth:       depth,
		Gpc:         gpc,
		GpcType:     gpcType,
		Keyword:     keyword,
	}
	DB.Table("goods_base").Create(&data)
}

func GetBarcodeGoods(barcode string) string {
	var result GoodsBaseItem
	sqlStr := `select goods_name from  goods_base where barcode=@barcode`
	DB.Raw(sqlStr, map[string]interface{}{
		"barcode": barcode,
	}).Scan(&result)
	return result.GoodsName
}
