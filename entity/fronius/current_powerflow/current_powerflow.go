package froniusCurrentPowerflow

import (
	"github.com/avegao/iot-api/entity/fronius"
)

type CurrentPowerflow struct {
	Body struct {
		Inverters map[int]inverter `json:"Inverters"`
		Site      site             `json:"Site"`
		SmartLoads *struct {
			Ohmpilots *map[int]ohmpilot `json:"Ohmpilots"`
		} `json:"Smartloads"`
		Version string `json:"Version"`
	} `json:"Body"`
	Head fronius.ResponseHeader `json:"Head"`
}

func (powerflow CurrentPowerflow) Persist() error {
	if site, err := powerflow.Body.Site.Persist(); err != nil {
		return err
	} else {
		powerflow.Body.Site = site
	}

	for _, inverter := range powerflow.Body.Inverters {
		inverter.Persist(powerflow.Body.Site.Id)
	}

	if nil != powerflow.Body.SmartLoads && nil != powerflow.Body.SmartLoads.Ohmpilots {
		for _, ohmpilot := range *powerflow.Body.SmartLoads.Ohmpilots {
			ohmpilot.Persist(powerflow.Body.Site.Id)
		}
	}

	return nil
}
