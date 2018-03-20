package fronius

const (
	SiteModeProduceOnly   SiteMode = "produce-only"
	SiteModeMeter         SiteMode = "meter"
	SiteModeVagueMeter    SiteMode = "vague-meter"
	SiteModeBidirectional SiteMode = "bidirectional"
	SiteModeAcCoupled     SiteMode = "ac-coupled"
)

type SiteMode string

func (mode SiteMode) String() string {
	return string(mode)
}
