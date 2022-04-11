package main

import (
	"fmt"
	"reminder/jobs"
	"reminder/routes"

	"github.com/bamzi/jobrunner"
)

func main() {

	go routes.InitApiRoute()

	jobrunner.Start()
	jobrunner.Schedule("@every 8h", jobs.EventJob{})
	jobrunner.Schedule("@daily", jobs.CheckPlanJob{})

	var str string
	fmt.Scan(&str)
	// path := `https://jumbarcode.market.alicloudapi.com/bar-code/query`
	// barcode := "6931987042704"
	// client := &http.Client{}
	// urlValues := url.Values{}
	// urlValues.Set("code", barcode)
	// requestData := urlValues.Encode()
	// req, err := http.NewRequest("POST", path, strings.NewReader(requestData))
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	// req.Header.Set("Authorization", "APPCODE "+os.Getenv("ALI_API_TOKEN"))

	// resp, err := client.Do(req)
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	// return "", err
	// 	fmt.Println(err)
	// }
	// fmt.Println(body)
	// var respJson controller.RBarcodeGoodsItem
	// json.Unmarshal([]byte(string(body)), &respJson)
	// fmt.Println(respJson)
}
