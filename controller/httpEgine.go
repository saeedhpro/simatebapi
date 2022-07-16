package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/controller/adminController"
	"github.com/saeedhpro/apisimateb/controller/appointmentController"
	"github.com/saeedhpro/apisimateb/controller/authController"
	"github.com/saeedhpro/apisimateb/controller/caseTypeController"
	"github.com/saeedhpro/apisimateb/controller/cityController"
	"github.com/saeedhpro/apisimateb/controller/countyController"
	"github.com/saeedhpro/apisimateb/controller/fileController"
	"github.com/saeedhpro/apisimateb/controller/medicalHistoryController"
	"github.com/saeedhpro/apisimateb/controller/messageController"
	"github.com/saeedhpro/apisimateb/controller/organizationController"
	"github.com/saeedhpro/apisimateb/controller/provinceController"
	"github.com/saeedhpro/apisimateb/controller/scheduleController"
	"github.com/saeedhpro/apisimateb/controller/userController"
	"github.com/saeedhpro/apisimateb/middleware"
)

func Run(Port string) {
	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.Static("/images", "./images")
	engine.Static("/file", "./res/file")
	engine.Static("/img", "./res/img")
	engine.Static("/public", "./res/public")
	engine.Use(middleware.CORSMiddleware)
	v1 := engine.Group("api/v1")

	authCont := authController.NewAuthController()
	organizationCont := organizationController.NewOrganizationController()
	userCont := userController.NewUserController()
	appointmentCont := appointmentController.NewAppointmentController()
	fileCont := fileController.NewFileController()
	provinceCont := provinceController.NewProvinceController()
	countyCont := countyController.NewCountyController()
	cityCont := cityController.NewCityController()
	caseCont := caseTypeController.NewCaseTypeController()
	adminCont := adminController.NewAdminController()
	messageCont := messageController.NewMessageController()
	medicalHistoryCont := medicalHistoryController.NewMedicalHistoryController()
	scheduleCont := scheduleController.NewScheduleController()

	v1.POST("/auth/login", authCont.Login)
	v1.GET("/own", middleware.GinJwtAuth(userCont.Own, true, false))

	organizations := v1.Group("/organizations")
	users := v1.Group("/users")
	appointments := v1.Group("/appointments")
	files := v1.Group("/files")
	provinces := v1.Group("/provinces")
	counties := v1.Group("/counties")
	cities := v1.Group("/cities")
	cases := v1.Group("/cases")
	admin := v1.Group("/admin")
	messages := v1.Group("/messages")

	{
		organizations.GET("/type", middleware.GinJwtAuth(organizationCont.GetOrganizationByType, true, false))
		organizations.GET("/users", middleware.GinJwtAuth(userCont.GetOrganizationUsersList, true, false))
		organizations.GET("/:id/users/all", middleware.GinJwtAuth(userCont.GetOrganizationUserList, true, false))
		organizations.GET("/appointments", middleware.GinJwtAuth(appointmentCont.GetOrganizationAppointmentList, true, false))
		organizations.POST("/appointments", middleware.GinJwtAuth(appointmentCont.CreateAppointment, true, false))
		organizations.GET("/messages", middleware.GinJwtAuth(messageCont.GetOrganizationMessages, true, false))
		organizations.GET("/cases", middleware.GinJwtAuth(caseCont.GetOrganizationCaseTypeList, true, false))
		organizations.POST("/cases", middleware.GinJwtAuth(caseCont.CreateCaseType, true, false))
		organizations.GET("/schedules", middleware.GinJwtAuth(scheduleCont.GetOrganizationScheduleList, true, false))
		organizations.GET("/:id", middleware.GinJwtAuth(organizationCont.Get, true, false))
		organizations.PUT("/:id/about", middleware.GinJwtAuth(organizationCont.UpdateOrganizationAbout, true, false))
		organizations.GET("/:id/users", middleware.GinJwtAuth(userCont.GetOrganizationPatientList, true, false))
		organizations.GET("/holidays", middleware.GinJwtAuth(organizationCont.GetHolidays, true, false))
		organizations.POST("/holidays", middleware.GinJwtAuth(organizationCont.CreateHoliday, true, false))
		organizations.PUT("/holidays/:id", middleware.GinJwtAuth(organizationCont.UpdateHoliday, true, false))
		organizations.DELETE("/holidays/:id", middleware.GinJwtAuth(organizationCont.DeleteHoliday, true, false))
	}
	{
		v1.POST("/users", middleware.GinJwtAuth(userCont.CreateUser, true, false))
		users.POST("/delete", middleware.GinJwtAuth(userCont.DeleteUsers, true, false))
		users.GET("/:id", middleware.GinJwtAuth(userCont.GetUser, true, false))
		users.PUT("/:id", middleware.GinJwtAuth(userCont.UpdateUser, true, false))
		users.DELETE("/:id", middleware.GinJwtAuth(userCont.DeleteUser, true, false))
		users.GET("/:id/appointments", middleware.GinJwtAuth(appointmentCont.GetUserAppointmentList, true, false))
		users.GET("/:id/appointments/resulted", middleware.GinJwtAuth(appointmentCont.GetUserResultedAppointmentList, true, false))
		users.GET("/:id/histories", middleware.GinJwtAuth(medicalHistoryCont.GetUserMedicalHistory, true, false))
		users.POST("/:id/histories", middleware.GinJwtAuth(medicalHistoryCont.CreateUserMedicalHistory, true, false))
	}
	{
		appointments.GET("/que", middleware.GinJwtAuth(appointmentCont.GetQueList, true, false))
		appointments.GET("/search", middleware.GinJwtAuth(appointmentCont.FilterOrganizationAppointment, true, false))
		appointments.GET("/:id", middleware.GinJwtAuth(appointmentCont.GetAppointment, true, false))
		appointments.PUT("/:id", middleware.GinJwtAuth(appointmentCont.UpdateAppointment, true, false))
		appointments.GET("/:id/results", middleware.GinJwtAuth(appointmentCont.GetAppointmentResults, true, false))
		appointments.POST("/:id/accept", middleware.GinJwtAuth(appointmentCont.AcceptAppointment, true, false))
		appointments.POST("/:id/cancel", middleware.GinJwtAuth(appointmentCont.CancelAppointment, true, false))
		appointments.POST("/:id/reserve", middleware.GinJwtAuth(appointmentCont.ReserveAppointment, true, false))
	}
	{
		v1.POST("/files", middleware.GinJwtAuth(fileCont.CreateFile, true, false))
		files.GET("/:id", middleware.GinJwtAuth(fileCont.GetUserFileList, true, false))
		files.DELETE("/:id", middleware.GinJwtAuth(fileCont.DeleteFile, true, false))
	}
	{
		provinces.GET("", middleware.GinJwtAuth(provinceCont.GetProvinceList, true, false))
		provinces.GET("/:id", middleware.GinJwtAuth(provinceCont.GetProvince, true, false))
		provinces.GET("/:id/counties", middleware.GinJwtAuth(countyCont.GetCountyList, true, false))
	}
	{
		counties.GET("/:id", middleware.GinJwtAuth(countyCont.GetCounty, true, false))
		counties.GET("/:id/cities", middleware.GinJwtAuth(cityCont.GetCityList, true, false))
	}
	{
		cities.GET("/:id", middleware.GinJwtAuth(cityCont.GetCity, true, false))
	}
	{
		cases.GET("/:id", middleware.GinJwtAuth(caseCont.Get, true, false))
		cases.PUT("/:id", middleware.GinJwtAuth(caseCont.UpdateCaseType, true, false))
		cases.DELETE("/:id", middleware.GinJwtAuth(caseCont.DeleteCaseType, true, false))
	}

	{
		admin.GET("/users/online", middleware.GinJwtAuth(adminCont.LastOnlineUsers, true, false))
		admin.GET("/users", middleware.GinJwtAuth(adminCont.GetUsers, true, false))
		admin.POST("/users", middleware.GinJwtAuth(adminCont.CreateUser, true, false))
		admin.GET("/patients/online", middleware.GinJwtAuth(adminCont.LastOnlinePatients, true, false))
		admin.GET("/organizations", middleware.GinJwtAuth(adminCont.GetOrganizations, true, false))
		admin.POST("/organizations", middleware.GinJwtAuth(adminCont.CreateOrganization, true, false))
		admin.GET("/organizations/:id/prof", middleware.GinJwtAuth(adminCont.GetOrganizationsByProfession, true, false))
		admin.PUT("/organizations/:id", middleware.GinJwtAuth(adminCont.UpdateOrganization, true, false))
		admin.GET("/groups", middleware.GinJwtAuth(adminCont.GetUserGroups, true, false))
		admin.GET("/messages", middleware.GinJwtAuth(adminCont.GetMessages, true, false))
		admin.POST("/messages", middleware.GinJwtAuth(messageCont.SendSms, true, false))
		admin.GET("/professions", middleware.GinJwtAuth(adminCont.GetProfessions, true, false))
		admin.POST("/messages/delete", middleware.GinJwtAuth(adminCont.DeleteMessages, true, false))
		admin.GET("/holidays", middleware.GinJwtAuth(adminCont.GetHolidays, true, false))
		admin.POST("/holidays", middleware.GinJwtAuth(adminCont.CreateHoliday, true, false))
		admin.PUT("/holidays/:id", middleware.GinJwtAuth(adminCont.UpdateHoliday, true, false))
		admin.DELETE("/holidays/:id", middleware.GinJwtAuth(adminCont.DeleteHoliday, true, false))
	}

	{
		v1.POST("/messages", middleware.GinJwtAuth(messageCont.SendSms, true, false))
		messages.POST("/delete", middleware.GinJwtAuth(messageCont.DeleteMessages, true, false))
	}

	fmt.Println(engine.Run(fmt.Sprintf(":%s", Port)))
}
