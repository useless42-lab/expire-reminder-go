package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"reminder/models"
	"reminder/response"
	"reminder/service"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrderList(c *gin.Context) {
	userId := c.GetInt64("userId")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetOrderListService(userId, page, pageSize)
	response.Success(c, 200, result)
}

type PlanJsonStruct struct {
	Data []PlanDataJsonstruct
}

type PlanDataJsonstruct struct {
	UrlKey string `json:"urlkey"`
	PlanId int    `json:"plan_id"`
	Time   int    `json:"time"`
}

func AddOrder(c *gin.Context) {
	userId := c.GetInt64("userId")
	orderId := c.PostForm("order_id")
	isCouponCode := strings.Contains(orderId, "yhm_")
	if isCouponCode {
		couponCodeDetail := models.GetCouponDetail(orderId)
		if couponCodeDetail.CouponCode == "" {
			response.Error(c, 2010, "没有该优惠码")
			return
		} else if couponCodeDetail.CouponCode != "" && couponCodeDetail.Status == 0 {
			response.Error(c, 2010, "该优惠码已被使用")
			return
		} else if couponCodeDetail.CouponCode != "" && couponCodeDetail.Status == 1 {
			models.UseCouponCode(couponCodeDetail.CouponCode)
			models.AddOrder(userId, time.Now().Unix(), 0, "success", "coupon", couponCodeDetail.CouponCode, "", couponCodeDetail.TargetPlanId, couponCodeDetail.TargetPlanTime, 1)
			response.Success(c, 200, "请静待发放")
		}
	} else {
		mbdDetail := OnGetOrderDetailApi(orderId)
		filePtr, err := os.Open("plan.json")
		if err != nil {
			return
		}
		defer filePtr.Close()

		var respJson PlanJsonStruct
		decoder := json.NewDecoder(filePtr)
		err = decoder.Decode(&respJson)
		if err != nil {
			fmt.Println("Decoder failed", err.Error())
		} else {
			var targetPlanId, targetPlanTime int
			for _, item := range respJson.Data {
				if item.UrlKey == mbdDetail.Result.UrlKey {
					targetPlanId = item.PlanId
					targetPlanTime = item.Time
				}
			}
			if mbdDetail.Code == 200 {
				if mbdDetail.Result.State == "success" {
					if targetPlanId == 0 {
						response.Error(c, 2011, "订单号有误")
					} else {
						isExist := models.IsExistOrder(userId, mbdDetail.Result.OrderId)
						if isExist {
							response.Error(c, 2001, "该订单已存在")
						} else {
							userPlan := models.GetUserPlan(userId)
							if userPlan.PlanId == 1 {
								models.AddOrder(userId, mbdDetail.Result.OrderTime, mbdDetail.Result.OrderAmount, mbdDetail.Result.State, mbdDetail.Result.Payway, mbdDetail.Result.OrderId, mbdDetail.Result.UrlKey, targetPlanId, targetPlanTime, 0)
								models.UpdateUserPlan(userId, targetPlanId, time.Now().AddDate(0, targetPlanTime, 0))
							} else {
								if time.Now().After(userPlan.ExpiredAt) {
									models.AddOrder(userId, mbdDetail.Result.OrderTime, mbdDetail.Result.OrderAmount, mbdDetail.Result.State, mbdDetail.Result.Payway, mbdDetail.Result.OrderId, mbdDetail.Result.UrlKey, targetPlanId, targetPlanTime, 0)
									models.UpdateUserPlan(userId, targetPlanId, time.Now().AddDate(0, targetPlanTime, 0))
								} else {
									models.AddOrder(userId, mbdDetail.Result.OrderTime, mbdDetail.Result.OrderAmount, mbdDetail.Result.State, mbdDetail.Result.Payway, mbdDetail.Result.OrderId, mbdDetail.Result.UrlKey, targetPlanId, targetPlanTime, 1)
								}
							}
							response.Success(c, 200, "")
						}
					}
				}
			} else {
				response.Error(c, 2010, "订单号有误")
			}
		}
	}
}
