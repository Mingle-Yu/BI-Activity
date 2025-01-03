package routes

import (
	"bi-activity/configs"
	"bi-activity/controller/collegeController"
	"bi-activity/dao"
	"bi-activity/dao/collegeDAO"
	"bi-activity/service/collegeService"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 配置实例
var config = configs.InitConfig("./configs")

// 数据库连接实例
var data, _ = dao.NewDateDao(config.Database, logrus.New())

func College(r *gin.Engine) {
	personalCenter(r)
	memberManagement(r)
	activityManagement(r)
}

func personalCenter(r *gin.Engine) {
	// DAO层实例
	pcDAO := collegeDAO.NewPcDAO(data)
	// Service层实例
	pcService := collegeService.NewPcService(pcDAO)

	pc := collegeController.NewPersonalCenter(pcService)
	pcGroup := r.Group("/college/personalCenter")
	pcGroup.GET("/collegeInfo", pc.GetCollegeInfo)
	pcGroup.POST("/collegeInfo", pc.UpdateCollegeInfo)
	pcGroup.GET("/adminInfo", pc.GetAdminInfo)
	pcGroup.POST("/adminInfo", pc.UpdateAdminInfo)
}

func memberManagement(r *gin.Engine) {
	// DAO层实例
	mmDAO := collegeDAO.NewMmDAO(data)
	// Service层实例
	mmService := collegeService.NewMmService(mmDAO)

	mm := collegeController.NewMemberManagement(mmService)
	mmGroup := r.Group("/college/memberManagement")
	mmGroup.GET("/audit", mm.GetAuditRecord)
	mmGroup.POST("/audit", mm.UpdateAuditRecord)
	mmGroup.GET("/query", mm.QueryMember)
	mmGroup.DELETE("/delete", mm.DeleteMember)
}

func activityManagement(r *gin.Engine) {
	// DAO层实例
	amDAO := collegeDAO.NewActivityManagementDAO(data)
	// Service层实例
	amService := collegeService.NewActivityManagementService(amDAO)

	amController := collegeController.NewActivityManagementController(amService)
	amGroup := r.Group("/college/activityManagement")
	amGroup.GET("/activity", amController.GetAuditRecord)
	amGroup.POST("/activity", amController.UpdateAuditRecord)
	amGroup.GET("/activityAdmission", amController.GetAdmissionRecord)
	amGroup.POST("/activityAdmission", amController.UpdateAdmissionRecord)
}
