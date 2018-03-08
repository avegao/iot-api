package main

import (
	"github.com/avegao/iot-api/controller/charger"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/avegao/gocondi"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	rootRouter := router.Group(fmt.Sprintf("/%s", apiVersion))

	initDocRouter(router)
	initChargeRouter(rootRouter)

	return router
}

func initDocRouter(router *gin.Engine) {
	if gocondi.GetContainer().GetBoolParameter("debug") {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func initChargeRouter(router *gin.RouterGroup) {
	chargerRouter := router.Group("/charger")
	chargerRouter.GET("/", chargerController.IndexAction)
	chargerRouter.GET("/:id", chargerController.FindOneAction)
	chargerRouter.GET("/:id/ev_connect_state", chargerController.GetEvConnectStateAction)
	chargerRouter.GET("/:id/ammeter_settings", chargerController.GetAmmeterSettingsAction)
	chargerRouter.GET("/:id/auth_lock_state", chargerController.GetAuthLockStateAction)
	chargerRouter.GET("/:id/charge_limit", chargerController.GetChargeLimitAction)
	chargerRouter.GET("/:id/current_capacity_range", chargerController.GetCurrentCapacityRangeInAmpsAction)
	chargerRouter.GET("/:id/delay_timer", chargerController.GetDelayTimerAction)
	chargerRouter.GET("/:id/energy_usage", chargerController.GetEnergyUsageAction)
	chargerRouter.GET("/:id/fault_counters", chargerController.GetFaultCountersAction)
	chargerRouter.GET("/:id/over_temperature_thresholds", chargerController.GetOverTemperatureThresholdsAction)
	chargerRouter.GET("/:id/rtc_time", chargerController.GetRtcTimeAction)
	chargerRouter.GET("/:id/settings", chargerController.GetSettingsAction)
	chargerRouter.GET("/:id/time_limit", chargerController.GetTimeLimitAction)
	chargerRouter.GET("/:id/version", chargerController.GetVersionAction)
	chargerRouter.GET("/:id/voltmeter_settings", chargerController.GetVoltmeterSettingsAction)
	chargerRouter.POST("/:id/rtc_time", chargerController.SetRtcTimeAction)
}
