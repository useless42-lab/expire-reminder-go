package models

import "time"

type CouponCodeStruct struct {
	DefaultModel
	CouponCode     string `json:"coupon_code" gorm:"column:coupon_code"`
	TargetPlanId   int    `json:"target_plan_id" gorm:"column:target_plan_id"`
	TargetPlanTime int    `json:"target_plan_time" gorm:"column:target_plan_time"`
	Status         int    `json:"status" gorm:"column:status"`
}

func GetCouponDetail(couponCode string) CouponCodeStruct {
	sqlStr := `select * from coupon_code where coupon_code=@couponCode`
	var result CouponCodeStruct
	DB.Raw(sqlStr, map[string]interface{}{
		"couponCode": couponCode,
	}).Scan(&result)
	return result
}

func UseCouponCode(couponCode string) {
	sqlStr := `update coupon_code set status=0 , deleted_at =@deletedAt where coupon_code=@couponCode`
	DB.Exec(sqlStr, map[string]interface{}{
		"deletedAt":  time.Now(),
		"couponCode": couponCode,
	})
}
