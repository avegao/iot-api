package froniusController

import (
	"github.com/gin-gonic/gin"
	"github.com/avegao/gocondi"
	"io/ioutil"
	"encoding/json"
)

// @Router /fronius/current_data/meter [post]
// @ID charger-index
// @Summary Get all chargers
// @Description Get all chargers
// @Produce json
// @Success 200 {} object "Array of chargers"
func PostCurrentDataMeterAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		logger.WithError(err).Panic()
	}

	var currentData CurrentDataMeter
	if err := json.Unmarshal(body, &currentData); err != nil {
		logger.WithError(err).Panic()
	}

	for _, bodyElement := range currentData.Body {
		bodyElement.Persist()
	}
}

// @Router /fronius/current_data/inverter [post]
// @ID charger-index
// @Summary Get all chargers
// @Description Get all chargers
// @Produce json
// @Success 200 {} object "Array of chargers"
func PostCurrentDataInverterAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		return
	}

	logger.WithField("body", string(body)).Debug()

	//var currentData CurrentDataMeter
	//json.Unmarshal(body, &currentData)
	//
	//logger.WithField("body", currentData).Debug()
}

// @Router /fronius/current_data/meter [post]
// @ID charger-index
// @Summary Get all chargers
// @Description Get all chargers
// @Produce json
// @Success 200 {} object "Array of chargers"
func PostCurrentDataPowerflowAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		return
	}

	logger.WithField("body", string(body)).Debug()

	//var currentData CurrentDataMeter
	//json.Unmarshal(body, &currentData)
	//
	//logger.WithField("body", currentData).Debug()
}

// @Router /fronius/current_data/sensor_card [post]
// @ID charger-index
// @Summary Get all chargers
// @Description Get all chargers
// @Produce json
// @Success 200 {} object "Array of chargers"
func PostCurrentDataSensorCardAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		return
	}

	logger.WithField("body", string(body)).Debug()

	//var currentData CurrentDataMeter
	//json.Unmarshal(body, &currentData)
	//
	//logger.WithField("body", currentData).Debug()
}

// @Router /fronius/current_data/string_control [post]
// @ID charger-index
// @Summary Get all chargers
// @Description Get all chargers
// @Produce json
// @Success 200 {} object "Array of chargers"
func PostCurrentDataStringControlAction(ginContext *gin.Context) {
	logger := gocondi.GetContainer().GetLogger()

	defer ginContext.Request.Body.Close()

	body, err := ioutil.ReadAll(ginContext.Request.Body)

	if err != nil {
		return
	}

	logger.WithField("body", string(body)).Debug()

	//var currentData CurrentDataMeter
	//json.Unmarshal(body, &currentData)
	//
	//logger.WithField("body", currentData).Debug()
}
