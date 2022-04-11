package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type RBarcodeGoodsItem struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	TaskNo string `json:"taskNo"`
	Data   struct {
		Code        string        `json:"code"`
		SptmImg     string        `json:"sptmImg"`
		Img         string        `json:"img"`
		GoodsType   string        `json:"goodsType"`
		Trademark   string        `json:"trademark"`
		GoodsName   string        `json:"goodsName"`
		Spec        string        `json:"spec"`
		Note        string        `json:"note"`
		Price       string        `json:"price"`
		Ycg         string        `json:"ycg"`
		ManuName    string        `json:"manuName"`
		ManuAddress string        `json:"manuAddress"`
		Qs          string        `json:"qs"`
		Nw          string        `json:"nw"`
		Description string        `json:"description"`
		Gw          string        `json:"gw"`
		Width       string        `json:"width"`
		Hight       string        `json:"hight"`
		Depth       string        `json:"depth"`
		Gpc         string        `json:"gpc"`
		GpcType     string        `json:"gpcType"`
		Keyword     string        `json:"keyword"`
		ImgList     []interface{} `json:"imgList"`
	} `json:"data"`
}

func OnGetGoodsApi(barcode string) RBarcodeGoodsItem {
	path := `https://jumbarcode.market.alicloudapi.com/bar-code/query`
	// resp, err := http.PostForm(path, url.Values{"code": {barcode}})
	// resp.Header.Set("Authorization", "APPCODE "+os.Getenv("ALI_API_TOKEN"))
	// if err != nil {
	// 	// handle error
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	// handle error
	// }
	// var respJson RBarcodeGoodsItem
	// json.Unmarshal([]byte(string(body)), &respJson)
	// return respJson

	client := &http.Client{}
	urlValues := url.Values{}
	urlValues.Set("code", barcode)
	requestData := urlValues.Encode()
	req, err := http.NewRequest("POST", path, strings.NewReader(requestData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Authorization", "APPCODE "+os.Getenv("ALI_API_TOKEN"))

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// return "", err
		fmt.Println(err)
	}
	fmt.Println(body)
	var respJson RBarcodeGoodsItem
	json.Unmarshal([]byte(string(body)), &respJson)
	return respJson
}
