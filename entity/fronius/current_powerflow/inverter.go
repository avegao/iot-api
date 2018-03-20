package froniusCurrentPowerflow

import (
	"github.com/avegao/iot-api/entity/fronius"
	"time"
	"fmt"
	"github.com/avegao/gocondi"
)

type inverter struct {
	BatteryMode *fronius.BatteryMode `json:"battery_mode"`

	// DeviceType Device type of inverter
	DeviceType int `json:"DT"`

	// EnergyDay Energy in Wh this day, null if no inverter is connected
	EnergyDay *int `json:"E_Day"`

	// EnergyDay Energy in Wh ever since, null if no inverter is connected
	EnergyTotal *int `json:"E_Total"`

	// EnergyDay Energy in Wh this year, null if no inverter is connected
	EnergyYear *int `json:"E_Year"`

	// CurrentPower current power in Watt, null if not running
	CurrentPower *int `json:"P"`

	// Soc Current state of charge in % ( 0 - 100% )
	Soc *int `json:"SOC"`
}

func (inverter inverter) getTableName() string {
	return "\"fronius\".\"current_powerflow_inverter\""
}

func (inverter inverter) Persist(siteId uint8) (error) {
	return inverter.insert(siteId)
}

func (inverter inverter) insert(siteId uint8) (err error) {
	const logTag = "CurrentPowerflow.inverter.insert()"
	startTimeLog := time.Now()
	container := gocondi.GetContainer()

	logger := container.GetLogger()
	logger.
		WithField("inverter", inverter).
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
		inverter.getTableName(),
	)

	logger.
		WithField("query", insertQuery).
		WithField("parameters", inverter).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> Query to execute", logTag))

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.WithError(err).Debugf("%s -> STOP", logTag)
		return
	}

	if _, err = db.Exec(insertQuery,
		siteId,
		inverter.BatteryMode,
		inverter.DeviceType,
		inverter.EnergyDay,
		inverter.EnergyYear,
		inverter.EnergyTotal,
		inverter.CurrentPower,
		inverter.Soc,
	); err != nil {
		logger.WithError(err).Panicf("%s -> STOP", logTag)
	}

	logger.
		WithField("inverter", inverter).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return
}
