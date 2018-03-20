package froniusCurrentPowerflow

import (
	"github.com/avegao/iot-api/entity/fronius"
	"time"
	"github.com/avegao/gocondi"
	"fmt"
	"errors"
)

type site struct {
	Id uint8

	Mode fronius.SiteMode `json:"Mode"`

	// BatteryStandby True when battery is in standby
	BatteryStandby *bool `json:"BatteryStandby"`

	// BackupMode Field is available if configured (false) or active (true)
	// if not available, mandatory config is not set.
	BackupMode *bool `json:"BackupMode"`

	// PowerFromGrid This value is null if no meter is enabled (+ from grid, - to grid)
	PowerFromGrid *float64 `json:"P_Grid"`

	// PowerLoad This value is null if no meter is enabled (+ generator, - consumer)
	PowerLoad *float64 `json:"P_Load"`

	// PowerAkku This value is null if no battery is active (+ charge, - discharge)
	PowerAkku *float64 `json:"P_Akku"`

	// PowerFromPV This value is null if inverter is not running (+ production (default))
	PowerFromPV *float64 `json:"P_PV"`

	// RelativeSelfConsumption Current relative self consumption in %, null if no smart meter is connected
	RelativeSelfConsumption *int `json:"rel_SelfConsumption"`

	// RelativeAutonomy Current relative autonomy in %, null if no smart meter is connected
	RelativeAutonomy *int `json:"rel_Autonomy"`

	MeterLocation *fronius.MeterLocation `json:"Meter_Location"`

	// EnergyDay Energy [Wh] this day, null if no inverter is connected
	EnergyDay *int `json:"E_Day"`

	// EnergyYear Energy [Wh] this year, null if no inverter is connected
	EnergyYear *int `json:"E_Year"`

	// EnergyTotal Energy [Wh] ever since, null if no inverter is connected
	EnergyTotal *int `json:"E_Total"`
}

func (site site) getTableName() string {
	return "\"fronius\".\"current_powerflow_site\""
}

func (site site) Persist() (site, error) {
	if site.Id == 0 {
		id, err := site.insert()

		site.Id = id

		if err != nil {
			return site, err
		}
	} else {
		return site, errors.New("update not supported yet")
	}

	return site, nil
}

func (site *site) insert() (id uint8, err error) {
	const logTag = "CurrentPowerflow.site.insert()"
	startTimeLog := time.Now()
	container := gocondi.GetContainer()

	logger := container.GetLogger()
	logger.
		WithField("site", *site).
		Debugf(fmt.Sprintf("%s -> START", logTag))

	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (
			"battery_standby",
			"backup_mode",
			"power_from_grid",
			"power_load",
			"power_akku",
			"power_from_pv",
			"relative_self_consumption",
			"relative_autonomy",
			"meter_location",
			"energy_day",
			"energy_year",
			"energy_total"
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,
			$10,$11,$12
		) RETURNING id;`,
		site.getTableName(),
	)

	logger.
		WithField("query", insertQuery).
		WithField("parameters", *site).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> Query to execute", logTag))

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.WithError(err).Debugf("%s -> STOP", logTag)
		return
	}

	if err := db.QueryRow(insertQuery,
		site.BatteryStandby,
		site.BackupMode,
		site.PowerFromGrid,
		site.PowerLoad,
		site.PowerAkku,
		site.PowerFromPV,
		site.RelativeSelfConsumption,
		site.RelativeAutonomy,
		site.MeterLocation,
		site.EnergyDay,
		site.EnergyYear,
		site.EnergyTotal,
	).Scan(&id); err != nil {
		logger.WithError(err).Panicf("%s -> STOP", logTag)
	}

	logger.
		WithField("site", *site).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		WithField("id", id).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return
}
