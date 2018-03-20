package fronius

const (
	MeterLocationLoad   MeterLocation = "load"
	MeterLocationGrid   MeterLocation = "grid"
	MeterLocationUnkown MeterLocation = "unknown"
)

type MeterLocation string

func (location MeterLocation) String() string {
	return string(location)
}
