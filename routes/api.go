package routes

import (
	"os"
	"reminder/controller"
	"reminder/middleware"

	"github.com/gin-gonic/gin"
)

func InitApiRoute() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	v1 := router.Group("/v1")
	{
		v1.GET("/team/invite", controller.GetInviteTeamMemberInfo)
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login)
			auth.POST("/register", controller.Register)
			auth.GET("/captcha", controller.GenerateCaptchaHandler)
			auth.POST("/reset/link", controller.GenerateResetPasswordLink)
			auth.POST("/reset", controller.ResetPassword)
		}
		user := v1.Group("/user")
		user.Use(middleware.CheckAuthMiddleware())
		{
			overview := user.Group("/overview")
			{
				overview.GET("/list", controller.GetOverviewTeamList)
				overview.POST("/search", controller.GetOverviewTeamSearch)
				overview.GET("/calendar", controller.GetOverviewCalendar)
			}
			team := user.Group("/team")
			{
				team.GET("/info", controller.GetTeamGroupInfo)
				team.GET("/list", controller.GetTeamGroupList)
				team.GET("/detail", controller.GetTeamGroupDetail)
				team.POST("/invite", controller.CreateInviteTeamMember)
				team.POST("/transfer", controller.TransferTeamGroup)
				team.POST("/create", controller.AddTeamGroup)
				team.POST("/update", controller.UpdateTeamGroup)
				team.DELETE("/delete", controller.DeleteTeamGroup)
				member := team.Group("/member")
				{
					member.GET("/list", controller.GetTeamGroupMemberList)
					member.POST("/exit", controller.ExitTeam)
					member.POST("/kick", controller.KickOutTeamMember)

					// 生成转让团队链接
					member.POST("/generate/transfer", controller.GenerateTransferTeamLink)
					// 生成邀请成员链接
					member.POST("/generate/invite", controller.GenerateInviteTeamMemberLink)
					//
					member.POST("/admin/remove", controller.RemoveAdmin)
					member.POST("/admin/add", controller.AddAdmin)
				}
			}
			goods := user.Group("goods")
			{
				goods.POST("/add", controller.AddGoods)
				goods.POST("/delete", controller.DeleteGoods)
				goods.GET("/list", controller.GetGoodsList)
				goods.GET("/list/expired", controller.GetExpiredGoodsList)
				goods.GET("/name", controller.GetGoodsName)
				goods.GET("/detail", controller.GetGoodsDetail)
				goods.GET("/price/timeline", controller.GetGoodsPriceTimeline)
				goods.POST("/thrown", controller.ThrownGoods)
			}
			group := user.Group("group")
			{
				group.GET("/list", controller.GetGoodsGroupList)
				group.POST("/add", controller.AddGoodsGroup)
				group.POST("/delete", controller.DeleteGoodsGroup)
				group.POST("/update", controller.UpdateGoodsGroup)
			}
			notification := user.Group("notification")
			{
				notification.GET("/detail", controller.GetNotification)
				notification.POST("/update", controller.UpdateNotification)
			}
			plan := user.Group("/plan")
			{
				plan.GET("/config", controller.GetPlanBaseInfo)
				plan.GET("/list", controller.GetPlanBaseList)
				plan.GET("/user", controller.GetUserPlanDetail)
			}
			order := user.Group("/order")
			{
				order.POST("/add", controller.AddOrder)
				order.GET("/list", controller.GetOrderList)
			}
			tag := user.Group("/tag")
			{
				tag.GET("/list", controller.GetTagList)
			}
		}
	}
	router.Run(":" + os.Getenv("ROUTE_PORT"))
}
