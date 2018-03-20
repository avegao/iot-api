package froniusCurrentPowerflow

import (
	"github.com/avegao/iot-api/entity/fronius"
	"time"
	"github.com/avegao/gocondi"
	"fmt"
)

type ohmpilot struct {
	// PowerAcTotal Current power consumption in Watt
	PowerAcTotal int `json:"P_AC_Total"`

	State fronius.OhmpilotState `json:"State"`

	// Temperature Temperature of storage / tank in degree Celsius
	Temperature float64 `json:"Temperature"`
}

func (ohmpilot ohmpilot) getTableName() string {
	return "\"fronius\".\"current_powerflow_ohmpilot\""
}

func (ohmpilot ohmpilot) Persist(siteId uint8) (error) {
	return ohmpilot.insert(siteId)
}

func (ohmpilot ohmpilot) insert(siteId uint8) (err error) {
	const logTag = "CurrentPowerflow.inverter.insert()"
	startTimeLog := time.Now()
	container := gocondi.GetContainer()

	logger := container.GetLogger()
	logger.
		WithField("ohmpilot", ohmpilot).
		Debugf(fmt.Sprintf("%s -> START", logTag))

	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (
			"id_site",
			"battery_mode",
			"device_type",
			"energy_day",
			"energy_year",
			"energy_total",
			"current_power",
			"soc"
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8
		);`,
		ohmpilot.getTableName(),
	)

	logger.
		WithField("query", insertQuery).
		WithField("parameters", ohmpilot).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> Query to execute", logTag))

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.WithError(err).Debugf("%s -> STOP", logTag)
		return
	}

	if _, err = db.Exec(insertQuery,
		siteId,
		ohmpilot.PowerAcTotal,
		ohmpilot.State,
		ohmpilot.Temperature,
	); err != nil {
		logger.WithError(err).Panicf("%s -> STOP", logTag)
	}

	logger.
		WithField("ohmpilot", ohmpilot).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return
}
