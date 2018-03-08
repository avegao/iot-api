package chargerController

import (
	"github.com/gin-gonic/gin"
	"github.com/avegao/iot-api/service/openevse"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/avegao/gocondi"
	"context"
	"net/http"
	"github.com/avegao/iot-api/util"
	"github.com/avegao/iot-api/resource/grpc/openevse"
	"time"
)

// @Router /charger [get]
// @ID charger-index
// @Summary Get all chargers
// @Description Get all chargers
// @Produce json
// @Success 200 {} object "Array of chargers"
func IndexAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)

	ctx := context.Background()
	response, err := client.FindAllChargers(ctx, &empty.Empty{})
	if err != nil {
		logger.WithError(err).Panic()
	}

	var chargers []openevseService.Charger

	for _, charger := range response.Chargers {
		chargers = append(chargers, openevseService.NewChargerFromGrpcResponse(*charger))
	}

	ginContext.JSON(http.StatusOK, chargers)
}

// @Router /charger/{id} [get]
// @ID get-charger
// @Summary Get one charger by ID
// @Description Get one charger by ID
// @Param id path string true "Charger ID"
// @Produce json
// @Success 200 {} object "Charger"
func FindOneAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))
	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	response, err := client.FindChargerById(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(http.StatusOK, openevseService.NewChargerFromGrpcResponse(*response))
}

// @Router /charger/{id}/ammeter_settings [get]
// @ID charger-get-ammeter-settings
// @Summary Get charger ammeter settings
// @Description Get charger ammeter settings
// @Param id path string true "Charger ID"
// @Produce json
// @Success 200 {getAmmeterSettingsResponse} getAmmeterSettingsResponse "Ammeter Settings"
func GetAmmeterSettingsAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetAmmeterSettings(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetAmmeterSettings(*grpcResponse),
	)
}

// @Router /charger/{id}/ammeter_settings [get]
// @ID charger-get-auth-lock-state
// @Summary Get auth lock state
// @Description Get auth lock state
// @Param id path string true "Charger ID"
// @Produce json
// @Success 200 {getAuthLockStateResponse} getAuthLockStateResponse "Auth lock state"
func GetAuthLockStateAction(ginContext *gin.Context) {
	//logger := gocondi.GetContainer().GetLogger()
	//
	//connection, err := openevseService.CreateConnection()
	//if err != nil {
	//	logger.WithError(err).Panic()
	//}
	//
	//client := openevseService.CreateClient(connection)
	//ctx := context.Background()
	//id := util.ParseUint64(ginContext.Param("id"))
	//
	//request := &iot_openevse_service.GetRequest{
	//	Id: id,
	//}
	//
	//grpcResponse, err := client.GetAuthLockState(ctx, request)
	//if err != nil {
	//	logger.WithError(err).Panic()
	//}
	//
	//ginContext.JSON(
	//	http.StatusOK,
	//	newResponseFromGetAuthLockState(*grpcResponse),
	//)

	ginContext.AbortWithStatus(http.StatusNotImplemented)
}

// @Router /charger/{id}/charge_limit [get]
// @ID charger-get-charge-limit
// @Summary Get charge limit
// @Description Get charge limit
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getChargeLimitResponse} getChargeLimitResponse "Limit"
func GetChargeLimitAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetChargeLimit(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetChargeLimit(*grpcResponse),
	)
}

// @Router /charger/{id}/current_capacity_range [get]
// @ID charger-get-current-capacity-range
// @Summary Get current capacity range in amps
// @Description Get current capacity range in amps
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getCurrentCapacityRangeInAmpsResponse} getCurrentCapacityRangeInAmpsResponse "Current capacity range in amps"
func GetCurrentCapacityRangeInAmpsAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetCurrentCapacityRangeInAmps(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetCurrentCapacityRangeInAmps(*grpcResponse),
	)
}

// @Router /charger/{id}/delay_timer [get]
// @ID charger-get-delay-timer
// @Summary Get delay timer
// @Description Get delay timer
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getDelayTimerResponse} getDelayTimerResponse "Delay timer"
func GetDelayTimerAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetDelayTimer(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetDelayTimer(*grpcResponse),
	)
}

// @Router /charger/{id}/energy_usage [get]
// @ID charger-get-energy-usage
// @Summary Get energy usage
// @Description Get energy usage
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getEnergyUsageResponse} getEnergyUsageResponse "Energy usage"
func GetEnergyUsageAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetEnergyUsage(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetEnergyUsage(*grpcResponse),
	)
}

// @Router /charger/{id}/ammeter_settings [get]
// @ID charger-get-ammeter-settings
// @Summary Get charger ammeter settings
// @Description Get charger ammeter settings
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getEvConnectStateResponse} getEvConnectStateResponse "EV connect state"
func GetEvConnectStateAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetEvConnectState(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetEvConnectState(*grpcResponse),
	)
}

// @Router /charger/{id}/fault_counters [get]
// @ID charger-get-fault-counters
// @Summary Get fault counters
// @Description Get fault counters
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getFaultCountersResponse} getFaultCountersResponse "Fault counters"
func GetFaultCountersAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetFaultCounters(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetFaultCounters(*grpcResponse),
	)
}

// @Router /charger/{id}/over_temperature_thresholds [get]
// @ID charger-get-over-temperature-thresholds
// @Summary Get over temperature thresholds
// @Description Get over temperature thresholds
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getOverTemperatureThresholdsResponse} getOverTemperatureThresholdsResponse "Over temperature thresholds"
func GetOverTemperatureThresholdsAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetOverTemperatureThresholds(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromOverTemperatureThresholds(*grpcResponse),
	)
}

// @Router /charger/{id}/rtc_time [get]
// @ID charger-get-rtc-time
// @Summary Get RTC Time
// @Description Get RTC Time
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getRtcTimeResponse} getRtcTimeResponse "RTC Time"
func GetRtcTimeAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetRtcTime(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetRtcTime(*grpcResponse),
	)
}

// @Router /charger/{id}/settings [get]
// @ID charger-get-settings
// @Summary Get settings
// @Description Get settings
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getSettingsResponse} getSettingsResponse "Settings"
func GetSettingsAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetSettings(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetSettings(*grpcResponse),
	)
}

// @Router /charger/{id}/time_limit [get]
// @ID charger-get-time-limit
// @Summary Get time limit
// @Description Get time limit
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getTimeLimitResponse} getTimeLimitResponse "Time limit"
func GetTimeLimitAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetTimeLimit(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetTimeLimit(*grpcResponse),
	)
}

// @Router /charger/{id}/version [get]
// @ID charger-get-version
// @Summary Get version
// @Description Get version
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getVersionResponse} getVersionResponse "Version"
func GetVersionAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetVersion(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetVersion(*grpcResponse),
	)
}

// @Router /charger/{id}/voltmeter_settings [get]
// @ID charger-get-voltmeter-settings
// @Summary Get voltmeter settings
// @Description Get voltmeter settings
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {getVoltmeterSettingsResponse} getVoltmeterSettingsResponse "Voltmeter settings"
func GetVoltmeterSettingsAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	request := &iot_openevse_service.GetRequest{
		Id: id,
	}

	grpcResponse, err := client.GetVoltmeterSettings(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		newResponseFromGetVoltmeterSettings(*grpcResponse),
	)
}

// @Router /charger/{id}/rtc_time [post]
// @ID charger-set-rtc-time
// @Summary Set RTC time
// @Description Set RTC time
// @Param id path int true "Charger ID"
// @Produce json
// @Success 200 {}  ""
func SetRtcTimeAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	connection, err := openevseService.CreateConnection()
	if err != nil {
		logger.WithError(err).Panic()
	}

	client := openevseService.CreateClient(connection)
	ctx := context.Background()
	id := util.ParseUint64(ginContext.Param("id"))

	rtcTime := time.Now().Format(time.RFC3339)

	if ginContext.PostForm("rtc_time") != "" {
		rtcTime = ginContext.PostForm("rtc_time")
	}

	request := &iot_openevse_service.SetRtcTimeRequest{
		Id: id,
		RtcTime: rtcTime,
	}

	grpcResponse, err := client.SetRtcTime(ctx, request)
	if err != nil {
		logger.WithError(err).Panic()
	}

	ginContext.JSON(
		http.StatusOK,
		struct {
			Ok bool
		}{
			Ok: grpcResponse.GetOk(),
		},
	)
}

